CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE FUNCTION public.calculatedriverstats(p_driver_id uuid)
RETURNS TABLE(
    total_races bigint,
    total_wins bigint,
    total_podiums bigint,
    total_points bigint,
    total_poles bigint,
    best_finish integer,
    best_qualifying integer,
    championship_wins bigint,
    best_championship_position bigint
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    WITH driver_lineups AS (
        SELECT DISTINCT
            tp.car_p AS car_p_id,
            rc.name AS raceclass_name
        FROM team_p tp
        JOIN car_p cp ON cp.id = tp.car_p
        JOIN car c ON c.id = cp.car
        JOIN raceclass rc ON rc.id = c.raceclass
        WHERE tp.driver = p_driver_id
    ),
    race_results AS (
        SELECT
            r.id AS race_id,
            r.championship AS championship_id,
            cp.id AS car_p_id,
            f.pos AS finish_pos,
            CEIL(
                CASE f.pos
                    WHEN 1 THEN 25
                    WHEN 2 THEN 18
                    WHEN 3 THEN 15
                    WHEN 4 THEN 12
                    WHEN 5 THEN 10
                    WHEN 6 THEN 8
                    WHEN 7 THEN 6
                    WHEN 8 THEN 4
                    WHEN 9 THEN 2
                    WHEN 10 THEN 1
                    ELSE 0
                END *
                CASE
                    WHEN r.type = 2 AND r.duration > 6 AND r.duration < 24 THEN 1.5
                    WHEN r.type = 2 AND r.duration >= 24 THEN 2
                    ELSE 1
                END
            )::int AS points
        FROM driver_lineups dl
        JOIN car_p cp ON cp.id = dl.car_p_id
        JOIN finish f ON f.car_p = cp.id
        JOIN race r ON r.id = f.race
    ),
    qualifying_results AS (
        SELECT q.pos
        FROM qualifying q
        JOIN driver_lineups dl ON dl.car_p_id = q.car_p
    ),
    championships AS (
        SELECT DISTINCT r.championship AS championship_id
        FROM race r
    ),
    driver_championships AS (
        SELECT DISTINCT championship_id
        FROM race_results
    ),
    championship_positions AS (
        SELECT
            ch.championship_id,
            cpp.car_p_id,
            cpp.raceclass,
            ROW_NUMBER() OVER (
                PARTITION BY ch.championship_id, cpp.raceclass
                ORDER BY cpp.points DESC, cpp.car_p_id
            ) AS position
        FROM championships ch
        CROSS JOIN LATERAL calculatepersonalpoints(ch.championship_id) cpp
    ),
    driver_championship_positions AS (
        SELECT
            cp.championship_id,
            cp.raceclass,
            MIN(cp.position) AS position
        FROM championship_positions cp
        WHERE cp.car_p_id IN (
            SELECT dl.car_p_id
            FROM driver_lineups dl
        )
        AND EXISTS (
            SELECT 1
            FROM driver_lineups dl
            WHERE dl.car_p_id = cp.car_p_id
              AND dl.raceclass_name = cp.raceclass
        )
        AND cp.championship_id IN (
            SELECT championship_id FROM driver_championships
        )
        GROUP BY cp.championship_id, cp.raceclass
    )
    SELECT
        (SELECT COUNT(DISTINCT race_id) FROM race_results),
        (SELECT COUNT(*) FROM race_results WHERE finish_pos = 1),
        (SELECT COUNT(*) FROM race_results WHERE finish_pos <= 3),
        (SELECT COALESCE(SUM(points), 0) FROM race_results),
        (SELECT COUNT(*) FROM qualifying_results WHERE pos = 1),
        (SELECT MIN(finish_pos) FROM race_results),
        (SELECT MIN(pos) FROM qualifying_results),
        (SELECT COUNT(*) FROM driver_championship_positions WHERE position = 1),
        (SELECT MIN(position) FROM driver_championship_positions);
END;
$$;

CREATE FUNCTION public.calculatepersonalpoints(championship uuid)
RETURNS TABLE(
    car_p_id uuid,
    team text,
    number text,
    manufacturer text,
    model text,
    raceclass text,
    points integer
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT
        cp.id AS car_p_id,
        t.name AS team,
        cp.number,
        m.name AS manufacturer,
        c.model,
        rc.name AS raceclass,
        pt.points
    FROM car_p cp
    JOIN car c ON cp.car = c.id
    JOIN (
        SELECT
            cp.id,
            SUM(
                CEIL(
                    CASE f.pos
                        WHEN 1 THEN 25
                        WHEN 2 THEN 18
                        WHEN 3 THEN 15
                        WHEN 4 THEN 12
                        WHEN 5 THEN 10
                        WHEN 6 THEN 8
                        WHEN 7 THEN 6
                        WHEN 8 THEN 4
                        WHEN 9 THEN 2
                        WHEN 10 THEN 1
                        ELSE 0
                    END *
                    CASE
                        WHEN r.type = 2 AND r.duration > 6 AND r.duration < 24 THEN 1.5
                        WHEN r.type = 2 AND r.duration >= 24 THEN 2
                        ELSE 1
                    END
                )
            )::INT AS points
        FROM finish f
        JOIN car_p cp ON f.car_p = cp.id
        JOIN race r ON f.race = r.id
        JOIN championship ch ON r.championship = ch.id
        WHERE ch.id = $1
        GROUP BY cp.id
    ) AS pt ON cp.id = pt.id
    JOIN manufacturer m ON c.manufacturer = m.id
    JOIN raceclass rc ON c.raceclass = rc.id
    JOIN team t ON cp.team = t.id
    ORDER BY pt.points DESC;
END;
$$;

CREATE FUNCTION public.calculateteamstats(p_team_id uuid)
RETURNS TABLE(
    total_races bigint,
    total_wins bigint,
    total_podiums bigint,
    total_points bigint,
    total_poles bigint,
    best_finish integer,
    best_qualifying integer,
    championship_wins bigint,
    best_championship_position bigint
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    WITH team_lineups AS (
        SELECT DISTINCT
            cp.id AS car_p_id,
            rc.name AS raceclass
        FROM car_p cp
        JOIN car c ON c.id = cp.car
        JOIN raceclass rc ON rc.id = c.raceclass
        WHERE cp.team = p_team_id
    ),
    race_results AS (
        SELECT
            r.id AS race_id,
            r.championship AS championship_id,
            cp.id AS car_p_id,
            f.pos AS finish_pos,
            CEIL(
                CASE f.pos
                    WHEN 1 THEN 25
                    WHEN 2 THEN 18
                    WHEN 3 THEN 15
                    WHEN 4 THEN 12
                    WHEN 5 THEN 10
                    WHEN 6 THEN 8
                    WHEN 7 THEN 6
                    WHEN 8 THEN 4
                    WHEN 9 THEN 2
                    WHEN 10 THEN 1
                    ELSE 0
                END *
                CASE
                    WHEN r.type = 2 AND r.duration > 6 AND r.duration < 24 THEN 1.5
                    WHEN r.type = 2 AND r.duration >= 24 THEN 2
                    ELSE 1
                END
            )::int AS points
        FROM team_lineups tl
        JOIN car_p cp ON cp.id = tl.car_p_id
        JOIN finish f ON f.car_p = cp.id
        JOIN race r ON r.id = f.race
    ),
    qualifying_results AS (
        SELECT q.pos
        FROM qualifying q
        JOIN team_lineups tl ON tl.car_p_id = q.car_p
    ),
    championships AS (
        SELECT DISTINCT r.championship AS championship_id
        FROM race r
    ),
    team_championships AS (
        SELECT DISTINCT championship_id
        FROM race_results
    ),
    championship_lineup_points AS (
        SELECT
            ch.championship_id,
            cpp.car_p_id,
            cpp.raceclass,
            cpp.points
        FROM championships ch
        CROSS JOIN LATERAL calculatepersonalpoints(ch.championship_id) cpp
    ),
    championship_team_points AS (
        SELECT
            clp.championship_id,
            cp.team,
            clp.raceclass,
            SUM(clp.points)::int AS points
        FROM championship_lineup_points clp
        JOIN car_p cp ON cp.id = clp.car_p_id
        GROUP BY clp.championship_id, cp.team, clp.raceclass
    ),
    championship_team_positions AS (
        SELECT
            championship_id,
            team,
            raceclass,
            ROW_NUMBER() OVER (
                PARTITION BY championship_id, raceclass
                ORDER BY points DESC, team
            ) AS position
        FROM championship_team_points
    ),
    team_positions AS (
        SELECT
            ctp.championship_id,
            ctp.raceclass,
            ctp.position
        FROM championship_team_positions ctp
        WHERE ctp.team = p_team_id
        AND ctp.championship_id IN (
            SELECT championship_id FROM team_championships
        )
    )
    SELECT
        (SELECT COUNT(DISTINCT race_id) FROM race_results),
        (SELECT COUNT(*) FROM race_results WHERE finish_pos = 1),
        (SELECT COUNT(*) FROM race_results WHERE finish_pos <= 3),
        (SELECT COALESCE(SUM(points), 0) FROM race_results),
        (SELECT COUNT(*) FROM qualifying_results WHERE pos = 1),
        (SELECT MIN(finish_pos) FROM race_results),
        (SELECT MIN(pos) FROM qualifying_results),
        (SELECT COUNT(*) FROM team_positions WHERE position = 1),
        (SELECT MIN(position) FROM team_positions);
END;
$$;

CREATE TABLE public.car (
    model text NOT NULL,
    year_of_production integer NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    manufacturer uuid,
    raceclass uuid
);

CREATE TABLE public.car_p (
    number text NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    car uuid,
    team uuid
);

CREATE TABLE public.championship (
    year integer NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    organizer uuid
);

CREATE TABLE public.driver (
    name text NOT NULL,
    photo bytea,
    nationality text NOT NULL,
    birthday date NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);

CREATE TABLE public.favourite_drivers (
    user_id uuid NOT NULL,
    driver uuid NOT NULL
);

CREATE TABLE public.favourite_teams (
    user_id uuid NOT NULL,
    team uuid NOT NULL
);

CREATE TABLE public.finish (
    pos integer NOT NULL,
    race uuid,
    car_p uuid
);

CREATE TABLE public.manufacturer (
    name text NOT NULL,
    country text NOT NULL,
    year_of_foundation integer NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);

CREATE TABLE public.organizer (
    name text NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);

CREATE TABLE public.qualifying (
    pos integer NOT NULL,
    car_p uuid,
    race uuid
);

CREATE TABLE public.race (
    name text NOT NULL,
    date date NOT NULL,
    type integer NOT NULL,
    duration integer NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    track uuid,
    championship uuid
);

CREATE TABLE public.raceclass (
    name text NOT NULL,
    docs text,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);

CREATE TABLE public.team (
    name text NOT NULL,
    logo bytea,
    country text,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);

CREATE TABLE public.team_p (
    car_p uuid,
    driver uuid
);

CREATE TABLE public.track (
    name text NOT NULL,
    country text NOT NULL,
    schema bytea,
    lap_length integer NOT NULL,
    turns integer NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    passwordhash character varying(256) NOT NULL,
    role text NOT NULL,
    created_at timestamp without time zone
);

CREATE TABLE public.db_sync_state (
    name text PRIMARY KEY,
    value text NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE public.sudoku_matrix_completions (
    user_id uuid NOT NULL,
    matrix_type text NOT NULL,
    completed_on date DEFAULT CURRENT_DATE NOT NULL
);

ALTER TABLE ONLY public.car_p
    ADD CONSTRAINT car_p_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.car
    ADD CONSTRAINT car_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.championship
    ADD CONSTRAINT championship_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.driver
    ADD CONSTRAINT driver_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.manufacturer
    ADD CONSTRAINT manufacturer_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.organizer
    ADD CONSTRAINT organizer_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.race
    ADD CONSTRAINT race_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.raceclass
    ADD CONSTRAINT raceclass_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.team
    ADD CONSTRAINT team_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.track
    ADD CONSTRAINT track_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.driver
    ADD CONSTRAINT unique_driver UNIQUE (name, nationality, birthday);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.sudoku_matrix_completions
    ADD CONSTRAINT sudoku_matrix_completions_unique UNIQUE (user_id, matrix_type, completed_on);

ALTER TABLE ONLY public.sudoku_matrix_completions
    ADD CONSTRAINT sudoku_matrix_completions_matrix_type_check
    CHECK (matrix_type IN ('drivers', 'teams'));

ALTER TABLE ONLY public.car
    ADD CONSTRAINT car_manufacturer_fkey
    FOREIGN KEY (manufacturer) REFERENCES public.manufacturer(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.car_p
    ADD CONSTRAINT car_p_car_fkey
    FOREIGN KEY (car) REFERENCES public.car(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.car_p
    ADD CONSTRAINT car_p_team_fkey
    FOREIGN KEY (team) REFERENCES public.team(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.car
    ADD CONSTRAINT car_raceclass_fkey
    FOREIGN KEY (raceclass) REFERENCES public.raceclass(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.championship
    ADD CONSTRAINT championship_organizer_fkey
    FOREIGN KEY (organizer) REFERENCES public.organizer(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.favourite_drivers
    ADD CONSTRAINT fdrivers_driver_fk
    FOREIGN KEY (driver) REFERENCES public.driver(id);

ALTER TABLE ONLY public.favourite_drivers
    ADD CONSTRAINT fdrivers_user_id_fk
    FOREIGN KEY (user_id) REFERENCES public.users(id);

ALTER TABLE ONLY public.finish
    ADD CONSTRAINT finish_car_p_fkey
    FOREIGN KEY (car_p) REFERENCES public.car_p(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.finish
    ADD CONSTRAINT finish_race_fkey
    FOREIGN KEY (race) REFERENCES public.race(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.favourite_teams
    ADD CONSTRAINT fteam_user_id_fk
    FOREIGN KEY (user_id) REFERENCES public.users(id);

ALTER TABLE ONLY public.favourite_teams
    ADD CONSTRAINT fteams_team_fk
    FOREIGN KEY (team) REFERENCES public.team(id);

ALTER TABLE ONLY public.qualifying
    ADD CONSTRAINT qualifying_car_p_fkey
    FOREIGN KEY (car_p) REFERENCES public.car_p(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.qualifying
    ADD CONSTRAINT qualifying_race_fkey
    FOREIGN KEY (race) REFERENCES public.race(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.sudoku_matrix_completions
    ADD CONSTRAINT sudoku_matrix_completions_user_id_fk
    FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.race
    ADD CONSTRAINT race_championship_fkey
    FOREIGN KEY (championship) REFERENCES public.championship(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.race
    ADD CONSTRAINT race_track_fkey
    FOREIGN KEY (track) REFERENCES public.track(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.team_p
    ADD CONSTRAINT team_p_car_p_fkey
    FOREIGN KEY (car_p) REFERENCES public.car_p(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.team_p
    ADD CONSTRAINT team_p_driver_fkey
    FOREIGN KEY (driver) REFERENCES public.driver(id) ON DELETE CASCADE;
