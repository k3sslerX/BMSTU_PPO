#!/usr/bin/env python3
from __future__ import annotations

from dataclasses import dataclass, field
import datetime as dt
import os
import sys
from typing import Any
from urllib.parse import urlparse
from uuid import UUID

import psycopg
from pymongo import MongoClient


POSTGRES = "postgres"
MONGO = "mongo"
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
        if self.dbms != POSTGRES:
            raise ValueError(f"{self.dbms} does not use SQL identifier quoting")
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
        mongo = Database(MONGO, connect_mongo())
        databases = {POSTGRES: postgres, MONGO: mongo}

        try:
            ensure_state_table(postgres)
            ensure_state_table(mongo)

            source_name, target_name, reason = choose_direction(active, forced_source, postgres, mongo)
            source = databases[source_name]
            target = databases[target_name]

            print(
                f"db-sync: source={source_name} target={target_name} active={active} "
                f"reason={reason} delete_stale={str(delete_stale).lower()}",
                flush=True,
            )

            summary = sync_databases(source, target, delete_stale)
            write_last_active(postgres, active)
            write_last_active(mongo, active)

            commit(postgres)
            commit(mongo)
        except Exception:
            rollback(postgres)
            rollback(mongo)
            raise
        finally:
            close_database(postgres)
            close_database(mongo)

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
    if normalized in ("mongodb", MONGO):
        return MONGO
    raise ValueError(f"unsupported DBMS {value!r}; use postgres or mongo")


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


def connect_mongo() -> Any:
    url = os.getenv("APP_MONGO_DATABASE_URL", "").strip()
    if not url:
        raise ValueError("APP_MONGO_DATABASE_URL is required")
    client = MongoClient(url)
    client.admin.command("ping")
    return client[mongo_database_name(url)]


def mongo_database_name(raw: str) -> str:
    parsed = urlparse(raw)
    database = parsed.path.strip("/")
    if database:
        return database
    database = os.getenv("APP_MONGO_DATABASE_NAME", "").strip()
    if database:
        return database
    raise ValueError("mongo database name is required in APP_MONGO_DATABASE_URL path or APP_MONGO_DATABASE_NAME")


def choose_direction(active: str, forced_source: str, postgres: Database, mongo: Database) -> tuple[str, str, str]:
    if forced_source == "disabled":
        write_last_active(postgres, active)
        write_last_active(mongo, active)
        commit(postgres)
        commit(mongo)
        print("db-sync: disabled by DB_SYNC_SOURCE", flush=True)
        sys.exit(0)

    if forced_source != "auto":
        return forced_source, other_dbms(forced_source), "forced"

    state = newest_state(read_state(postgres), read_state(mongo))
    if state.valid:
        if state.dbms == active:
            return active, other_dbms(active), "last-active-current"
        return state.dbms, active, "last-active-switch"

    postgres_count = total_row_count(postgres)
    mongo_count = total_row_count(mongo)
    if postgres_count > mongo_count:
        return POSTGRES, MONGO, "initial-row-count"
    if mongo_count > postgres_count:
        return MONGO, POSTGRES, "initial-row-count"

    return other_dbms(active), active, "initial-active-switch"


def other_dbms(dbms: str) -> str:
    return MONGO if dbms == POSTGRES else POSTGRES


def newest_state(left: State, right: State) -> State:
    if left.valid and right.valid:
        return left if left.updated_at >= right.updated_at else right
    if left.valid:
        return left
    return right


def read_state(database: Database) -> State:
    if database.dbms == MONGO:
        row = database.conn[STATE_TABLE].find_one({"_id": LAST_ACTIVE_KEY}) or database.conn[STATE_TABLE].find_one(
            {"name": LAST_ACTIVE_KEY}
        )
        if not row:
            return State()
        dbms = normalize_dbms(str(row.get("value", "")))
        updated_at = normalize_value(row.get("updated_at")) or dt.datetime.min
        return State(dbms=dbms, updated_at=updated_at, valid=True)

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
    if database.dbms == MONGO:
        database.conn[STATE_TABLE].update_one(
            {"_id": LAST_ACTIVE_KEY},
            {"$set": {"name": LAST_ACTIVE_KEY, "value": active, "updated_at": now}},
            upsert=True,
        )
        return

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


def ensure_state_table(database: Database) -> None:
    if database.dbms == MONGO:
        database.conn[STATE_TABLE].create_index("name", unique=True)
        for table in TABLES:
            if table.key_columns == ("id",):
                database.conn[table.name].create_index("id", unique=True)
        return

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


def total_row_count(database: Database) -> int:
    total = 0
    for table in TABLES:
        if database.dbms == MONGO:
            total += int(database.conn[table.name].count_documents({}))
            continue
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
    if database.dbms == MONGO:
        projection = {column: True for column in table.columns}
        projection["_id"] = False
        rows = []
        for document in database.conn[table.name].find({}, projection):
            rows.append({column: normalize_value(document.get(column)) for column in table.columns})
        return rows

    query = (
        f"SELECT {quoted_columns(database, table.columns)} "
        f"FROM {database.quote(table.name)}"
    )
    return rows_to_dicts(table, fetch_all(database, query))


def fetch_existing(database: Database, table: Table, row: dict[str, Any]) -> dict[str, Any] | None:
    if database.dbms == MONGO:
        projection = {column: True for column in table.columns}
        projection["_id"] = False
        document = database.conn[table.name].find_one({"_id": row_key(table, row)}, projection)
        if document is None:
            document = database.conn[table.name].find_one(
                {column: row[column] for column in table.key_columns},
                projection,
            )
        if document is None:
            return None
        return {column: normalize_value(document.get(column)) for column in table.columns}

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
    if database.dbms == MONGO:
        database.conn[table.name].insert_one(mongo_document(table, row))
        return

    placeholders = ", ".join(["%s"] * len(table.columns))
    query = (
        f"INSERT INTO {database.quote(table.name)} "
        f"({quoted_columns(database, table.columns)}) VALUES ({placeholders})"
    )
    execute(database, query, tuple(row[column] for column in table.columns))


def update_row(database: Database, table: Table, row: dict[str, Any]) -> None:
    if database.dbms == MONGO:
        database.conn[table.name].replace_one(
            {"_id": row_key(table, row)},
            mongo_document(table, row),
            upsert=True,
        )
        return

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
        if database.dbms == MONGO:
            database.conn[table.name].delete_one({"_id": row_key(table, row)})
            deleted += 1
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
    if isinstance(value, dt.date):
        return value.isoformat()
    return value


def mongo_document(table: Table, row: dict[str, Any]) -> dict[str, Any]:
    document = {"_id": row_key(table, row)}
    for column in table.columns:
        document[column] = normalize_value(row[column])
    return document


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


def commit(database: Database) -> None:
    if database.dbms == POSTGRES:
        database.conn.commit()


def rollback(database: Database) -> None:
    if database.dbms == POSTGRES:
        database.conn.rollback()


def close_database(database: Database) -> None:
    if database.dbms == MONGO:
        database.conn.client.close()
        return
    database.conn.close()


if __name__ == "__main__":
    raise SystemExit(main())
