#!/usr/bin/env python3
from __future__ import annotations

from dataclasses import dataclass, field
import datetime as dt
import os
import re
import sys
from typing import Any
from urllib.parse import unquote_plus, urlparse
from uuid import UUID

import pymysql
import psycopg


POSTGRES = "postgres"
MYSQL = "mysql"
STATE_TABLE = "db_sync_state"
LAST_ACTIVE_KEY = "last_active_dbms"


@dataclass(frozen=True)
class Table:
    name: str
    columns: tuple[str, ...]
    key_columns: tuple[str, ...]
    binary_columns: frozenset[str] = field(default_factory=frozenset)


@dataclass
class Database:
    dbms: str
    conn: Any

    def quote(self, identifier: str) -> str:
        if self.dbms == MYSQL:
            return "`" + identifier.replace("`", "``") + "`"
        return '"' + identifier.replace('"', '""') + '"'


@dataclass
class State:
    dbms: str = ""
    updated_at: dt.datetime = dt.datetime.min
    valid: bool = False


@dataclass
class Summary:
    source: str
    target: str
    inserted: int = 0
    updated: int = 0
    deleted: int = 0


def make_table(
    name: str,
    columns: tuple[str, ...],
    key_columns: tuple[str, ...],
    binary_columns: tuple[str, ...] = (),
) -> Table:
    return Table(name, columns, key_columns, frozenset(binary_columns))


TABLES = (
    make_table("manufacturer", ("name", "country", "year_of_foundation", "id"), ("id",)),
    make_table("raceclass", ("name", "docs", "id"), ("id",)),
    make_table("team", ("name", "logo", "country", "id"), ("id",), ("logo",)),
    make_table("organizer", ("name", "id"), ("id",)),
    make_table("track", ("name", "country", "schema", "lap_length", "turns", "id"), ("id",), ("schema",)),
    make_table("driver", ("name", "photo", "nationality", "birthday", "id"), ("id",), ("photo",)),
    make_table("car", ("model", "year_of_production", "id", "manufacturer", "raceclass"), ("id",)),
    make_table("championship", ("year", "id", "organizer"), ("id",)),
    make_table("car_p", ("number", "id", "car", "team"), ("id",)),
    make_table("race", ("name", "date", "type", "duration", "id", "track", "championship"), ("id",)),
    make_table("team_p", ("car_p", "driver"), ("car_p", "driver")),
    make_table("users", ("id", "name", "email", "passwordhash", "role", "created_at"), ("id",)),
    make_table("sudoku_matrix_completions", ("user_id", "matrix_type", "completed_on"), ("user_id", "matrix_type", "completed_on")),
    make_table("favourite_drivers", ("user_id", "driver"), ("user_id", "driver")),
    make_table("favourite_teams", ("user_id", "team"), ("user_id", "team")),
    make_table("finish", ("pos", "race", "car_p"), ("race", "car_p")),
    make_table("qualifying", ("pos", "car_p", "race"), ("race", "car_p")),
)


def main() -> int:
    try:
        validate_tables()
        active = normalize_dbms(os.getenv("DBMS", POSTGRES))
        forced_source = normalize_source(os.getenv("DB_SYNC_SOURCE", "auto"))
        delete_stale = env_bool("DB_SYNC_DELETE_STALE", False)

        postgres = Database(POSTGRES, connect_postgres())
        mysql = Database(MYSQL, connect_mysql())
        databases = {POSTGRES: postgres, MYSQL: mysql}

        try:
            ensure_state_table(postgres)
            ensure_state_table(mysql)

            source_name, target_name, reason = choose_direction(active, forced_source, postgres, mysql)
            source = databases[source_name]
            target = databases[target_name]

            print(
                f"db-sync: source={source_name} target={target_name} active={active} "
                f"reason={reason} delete_stale={str(delete_stale).lower()}",
                flush=True,
            )

            summary = sync_databases(source, target, delete_stale)
            write_last_active(postgres, active)
            write_last_active(mysql, active)

            postgres.conn.commit()
            mysql.conn.commit()
        except Exception:
            postgres.conn.rollback()
            mysql.conn.rollback()
            raise
        finally:
            postgres.conn.close()
            mysql.conn.close()

        print(
            f"db-sync: done source={summary.source} target={summary.target} "
            f"inserted={summary.inserted} updated={summary.updated} deleted={summary.deleted}",
            flush=True,
        )
        return 0
    except Exception as err:
        print(f"db-sync: failed: {err}", file=sys.stderr, flush=True)
        return 1


def validate_tables() -> None:
    names = set()
    for table in TABLES:
        if table.name in names:
            raise ValueError(f"duplicate table config: {table.name}")
        names.add(table.name)
        columns = set(table.columns)
        if len(columns) != len(table.columns):
            raise ValueError(f"duplicate columns in table config: {table.name}")
        missing_keys = [column for column in table.key_columns if column not in columns]
        if missing_keys:
            raise ValueError(f"unknown key columns for {table.name}: {', '.join(missing_keys)}")


def normalize_dbms(value: str) -> str:
    normalized = value.strip().lower()
    if normalized in ("", "pg", "postgresql", POSTGRES):
        return POSTGRES
    if normalized in ("mariadb", MYSQL):
        return MYSQL
    raise ValueError(f"unsupported DBMS {value!r}; use postgres or mysql")


def normalize_source(value: str) -> str:
    normalized = value.strip().lower()
    if normalized in ("", "auto"):
        return "auto"
    if normalized in ("off", "none", "false", "disabled"):
        return "disabled"
    return normalize_dbms(normalized)


def env_bool(name: str, default: bool) -> bool:
    raw = os.getenv(name)
    if raw is None or raw.strip() == "":
        return default
    return raw.strip().lower() in ("1", "true", "yes", "y", "on")


def connect_postgres() -> Any:
    url = os.getenv("APP_POSTGRES_DATABASE_URL", "").strip()
    if not url:
        raise ValueError("APP_POSTGRES_DATABASE_URL is required")
    return psycopg.connect(url)


def connect_mysql() -> Any:
    url = os.getenv("APP_MYSQL_DATABASE_URL", "").strip()
    if not url:
        raise ValueError("APP_MYSQL_DATABASE_URL is required")
    params = parse_mysql_dsn(url)
    return pymysql.connect(
        host=params["host"],
        port=int(params["port"]),
        user=params["user"],
        password=params["password"],
        database=params["database"],
        charset="utf8mb4",
        autocommit=False,
    )


def parse_mysql_dsn(raw: str) -> dict[str, str]:
    raw = raw.strip()
    without_scheme = raw.removeprefix("mysql://")
    go_match = re.match(
        r"^(?P<user>[^:@/]+)(?::(?P<password>[^@]*))?@tcp\((?P<host>[^:)]+)(?::(?P<port>\d+))?\)/(?P<database>[^?]+)",
        without_scheme,
    )
    if go_match:
        matched = go_match.groupdict()
        return {
            "user": unquote_plus(matched["user"]),
            "password": unquote_plus(matched.get("password") or ""),
            "host": matched["host"],
            "port": matched.get("port") or "3306",
            "database": unquote_plus(matched["database"]),
        }

    parsed = urlparse(raw if raw.startswith("mysql://") else "mysql://" + raw)
    if not parsed.hostname or not parsed.username or not parsed.path.strip("/"):
        raise ValueError("APP_MYSQL_DATABASE_URL must be a Go MySQL DSN or mysql:// URL")
    return {
        "user": unquote_plus(parsed.username),
        "password": unquote_plus(parsed.password or ""),
        "host": parsed.hostname,
        "port": str(parsed.port or 3306),
        "database": unquote_plus(parsed.path.strip("/")),
    }


def choose_direction(active: str, forced_source: str, postgres: Database, mysql: Database) -> tuple[str, str, str]:
    if forced_source == "disabled":
        write_last_active(postgres, active)
        write_last_active(mysql, active)
        postgres.conn.commit()
        mysql.conn.commit()
        print("db-sync: disabled by DB_SYNC_SOURCE", flush=True)
        sys.exit(0)

    if forced_source != "auto":
        return forced_source, other_dbms(forced_source), "forced"

    state = newest_state(read_state(postgres), read_state(mysql))
    if state.valid:
        if state.dbms == active:
            return active, other_dbms(active), "last-active-current"
        return state.dbms, active, "last-active-switch"

    postgres_count = total_row_count(postgres)
    mysql_count = total_row_count(mysql)
    if postgres_count > mysql_count:
        return POSTGRES, MYSQL, "initial-row-count"
    if mysql_count > postgres_count:
        return MYSQL, POSTGRES, "initial-row-count"

    return other_dbms(active), active, "initial-active-switch"


def other_dbms(dbms: str) -> str:
    return MYSQL if dbms == POSTGRES else POSTGRES


def newest_state(left: State, right: State) -> State:
    if left.valid and right.valid:
        return left if left.updated_at >= right.updated_at else right
    if left.valid:
        return left
    return right


def read_state(database: Database) -> State:
    query = (
        f"SELECT {database.quote('value')}, {database.quote('updated_at')} "
        f"FROM {database.quote(STATE_TABLE)} "
        f"WHERE {database.quote('name')} = %s"
    )
    rows = fetch_all(database, query, (LAST_ACTIVE_KEY,))
    if not rows:
        return State()
    dbms = normalize_dbms(str(rows[0][0]))
    return State(dbms=dbms, updated_at=normalize_value(rows[0][1]), valid=True)


def write_last_active(database: Database, active: str) -> None:
    now = dt.datetime.utcnow().replace(microsecond=0)
    if database.dbms == POSTGRES:
        query = (
            f"INSERT INTO {database.quote(STATE_TABLE)} "
            f"({database.quote('name')}, {database.quote('value')}, {database.quote('updated_at')}) "
            "VALUES (%s, %s, %s) "
            f"ON CONFLICT ({database.quote('name')}) DO UPDATE SET "
            f"{database.quote('value')} = EXCLUDED.{database.quote('value')}, "
            f"{database.quote('updated_at')} = EXCLUDED.{database.quote('updated_at')}"
        )
        execute(database, query, (LAST_ACTIVE_KEY, active, now))
        return

    query = (
        f"INSERT INTO {database.quote(STATE_TABLE)} "
        f"({database.quote('name')}, {database.quote('value')}, {database.quote('updated_at')}) "
        "VALUES (%s, %s, %s) "
        f"ON DUPLICATE KEY UPDATE {database.quote('value')} = %s, {database.quote('updated_at')} = %s"
    )
    execute(database, query, (LAST_ACTIVE_KEY, active, now, active, now))


def ensure_state_table(database: Database) -> None:
    if database.dbms == POSTGRES:
        execute(
            database,
            f"""
            CREATE TABLE IF NOT EXISTS {database.quote(STATE_TABLE)} (
                {database.quote('name')} text PRIMARY KEY,
                {database.quote('value')} text NOT NULL,
                {database.quote('updated_at')} timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
            )
            """,
        )
        return

    execute(
        database,
        f"""
        CREATE TABLE IF NOT EXISTS {database.quote(STATE_TABLE)} (
            {database.quote('name')} VARCHAR(128) NOT NULL,
            {database.quote('value')} VARCHAR(128) NOT NULL,
            {database.quote('updated_at')} DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
            PRIMARY KEY ({database.quote('name')})
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
        """,
    )


def total_row_count(database: Database) -> int:
    total = 0
    for table in TABLES:
        rows = fetch_all(database, f"SELECT COUNT(*) FROM {database.quote(table.name)}")
        total += int(rows[0][0])
    return total


def sync_databases(source: Database, target: Database, delete_stale: bool) -> Summary:
    summary = Summary(source=source.dbms, target=target.dbms)
    source_keys: dict[str, set[str]] = {}

    for table in TABLES:
        rows = fetch_table(source, table)
        source_keys[table.name] = {row_key(table, row) for row in rows}

        for row in rows:
            status = upsert_row(target, table, row)
            if status == "inserted":
                summary.inserted += 1
            elif status == "updated":
                summary.updated += 1

    if delete_stale:
        for table in reversed(TABLES):
            summary.deleted += delete_missing_rows(target, table, source_keys[table.name])

    return summary


def fetch_table(database: Database, table: Table) -> list[dict[str, Any]]:
    query = (
        f"SELECT {quoted_columns(database, table.columns)} "
        f"FROM {database.quote(table.name)}"
    )
    return rows_to_dicts(table, fetch_all(database, query))


def fetch_existing(database: Database, table: Table, row: dict[str, Any]) -> dict[str, Any] | None:
    where, args = build_where(database, table.key_columns, row)
    query = (
        f"SELECT {quoted_columns(database, table.columns)} "
        f"FROM {database.quote(table.name)} WHERE {where} LIMIT 1"
    )
    rows = fetch_all(database, query, args)
    if not rows:
        return None
    return rows_to_dicts(table, rows)[0]


def upsert_row(database: Database, table: Table, row: dict[str, Any]) -> str:
    existing = fetch_existing(database, table, row)
    if existing is None:
        insert_row(database, table, row)
        return "inserted"
    if rows_equal(table, existing, row):
        return "unchanged"
    update_row(database, table, row)
    return "updated"


def insert_row(database: Database, table: Table, row: dict[str, Any]) -> None:
    placeholders = ", ".join(["%s"] * len(table.columns))
    query = (
        f"INSERT INTO {database.quote(table.name)} "
        f"({quoted_columns(database, table.columns)}) VALUES ({placeholders})"
    )
    execute(database, query, tuple(row[column] for column in table.columns))


def update_row(database: Database, table: Table, row: dict[str, Any]) -> None:
    update_columns = [column for column in table.columns if column not in table.key_columns]
    if not update_columns:
        return
    assignments = ", ".join(f"{database.quote(column)} = %s" for column in update_columns)
    where, where_args = build_where(database, table.key_columns, row)
    query = f"UPDATE {database.quote(table.name)} SET {assignments} WHERE {where}"
    args = tuple(row[column] for column in update_columns) + where_args
    execute(database, query, args)


def delete_missing_rows(database: Database, table: Table, source_keys: set[str]) -> int:
    deleted = 0
    for row in fetch_table(database, table):
        if row_key(table, row) in source_keys:
            continue
        where, args = build_where(database, table.key_columns, row)
        execute(database, f"DELETE FROM {database.quote(table.name)} WHERE {where}", args)
        deleted += 1
    return deleted


def rows_to_dicts(table: Table, rows: list[tuple[Any, ...]]) -> list[dict[str, Any]]:
    result = []
    for values in rows:
        result.append({column: normalize_value(value) for column, value in zip(table.columns, values)})
    return result


def rows_equal(table: Table, left: dict[str, Any], right: dict[str, Any]) -> bool:
    return all(left[column] == right[column] for column in table.columns)


def row_key(table: Table, row: dict[str, Any]) -> str:
    return "\x1f".join(key_value(row[column]) for column in table.key_columns)


def key_value(value: Any) -> str:
    if value is None:
        return "\x00"
    if isinstance(value, bytes):
        return value.hex()
    if isinstance(value, dt.datetime):
        return value.isoformat(timespec="seconds")
    if isinstance(value, dt.date):
        return value.isoformat()
    return str(value)


def normalize_value(value: Any) -> Any:
    if isinstance(value, UUID):
        return str(value)
    if isinstance(value, bytearray):
        return bytes(value)
    if isinstance(value, memoryview):
        return value.tobytes()
    if isinstance(value, dt.datetime):
        if value.tzinfo is not None:
            value = value.astimezone(dt.timezone.utc).replace(tzinfo=None)
        return value.replace(microsecond=0)
    return value


def build_where(database: Database, columns: tuple[str, ...], row: dict[str, Any]) -> tuple[str, tuple[Any, ...]]:
    parts = []
    args = []
    for column in columns:
        if row[column] is None:
            parts.append(f"{database.quote(column)} IS NULL")
        else:
            parts.append(f"{database.quote(column)} = %s")
            args.append(row[column])
    return " AND ".join(parts), tuple(args)


def quoted_columns(database: Database, columns: tuple[str, ...]) -> str:
    return ", ".join(database.quote(column) for column in columns)


def fetch_all(database: Database, query: str, args: tuple[Any, ...] = ()) -> list[tuple[Any, ...]]:
    with database.conn.cursor() as cursor:
        cursor.execute(query, args)
        return list(cursor.fetchall())


def execute(database: Database, query: str, args: tuple[Any, ...] = ()) -> None:
    with database.conn.cursor() as cursor:
        cursor.execute(query, args)


if __name__ == "__main__":
    raise SystemExit(main())
