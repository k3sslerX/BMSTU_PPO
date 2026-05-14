CREATE TABLE manufacturer (
    name TEXT NOT NULL,
    country TEXT NOT NULL,
    year_of_foundation INT NOT NULL,
    id CHAR(36) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE raceclass (
    name TEXT NOT NULL,
    docs TEXT,
    id CHAR(36) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE team (
    name TEXT NOT NULL,
    logo LONGBLOB,
    country TEXT,
    id CHAR(36) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE organizer (
    name TEXT NOT NULL,
    id CHAR(36) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE track (
    name TEXT NOT NULL,
    country TEXT NOT NULL,
    `schema` LONGBLOB,
    lap_length INT NOT NULL,
    turns INT NOT NULL,
    id CHAR(36) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE driver (
    name TEXT NOT NULL,
    photo LONGBLOB,
    nationality TEXT NOT NULL,
    birthday DATE NOT NULL,
    id CHAR(36) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT unique_driver UNIQUE (name(255), nationality(255), birthday)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE car (
    model TEXT NOT NULL,
    year_of_production INT NOT NULL,
    id CHAR(36) NOT NULL,
    manufacturer CHAR(36),
    raceclass CHAR(36),
    PRIMARY KEY (id),
    CONSTRAINT car_manufacturer_fkey FOREIGN KEY (manufacturer) REFERENCES manufacturer(id) ON DELETE CASCADE,
    CONSTRAINT car_raceclass_fkey FOREIGN KEY (raceclass) REFERENCES raceclass(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE championship (
    year INT NOT NULL,
    id CHAR(36) NOT NULL,
    organizer CHAR(36),
    PRIMARY KEY (id),
    CONSTRAINT championship_organizer_fkey FOREIGN KEY (organizer) REFERENCES organizer(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE car_p (
    number TEXT NOT NULL,
    id CHAR(36) NOT NULL,
    car CHAR(36),
    team CHAR(36),
    PRIMARY KEY (id),
    CONSTRAINT car_p_car_fkey FOREIGN KEY (car) REFERENCES car(id) ON DELETE CASCADE,
    CONSTRAINT car_p_team_fkey FOREIGN KEY (team) REFERENCES team(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE race (
    name TEXT NOT NULL,
    date DATE NOT NULL,
    type INT NOT NULL,
    duration INT NOT NULL,
    id CHAR(36) NOT NULL,
    track CHAR(36),
    championship CHAR(36),
    PRIMARY KEY (id),
    CONSTRAINT race_track_fkey FOREIGN KEY (track) REFERENCES track(id) ON DELETE CASCADE,
    CONSTRAINT race_championship_fkey FOREIGN KEY (championship) REFERENCES championship(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE team_p (
    car_p CHAR(36),
    driver CHAR(36),
    CONSTRAINT team_p_car_p_fkey FOREIGN KEY (car_p) REFERENCES car_p(id) ON DELETE CASCADE,
    CONSTRAINT team_p_driver_fkey FOREIGN KEY (driver) REFERENCES driver(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE users (
    id CHAR(36) NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    passwordhash VARCHAR(256) NOT NULL,
    role TEXT NOT NULL,
    created_at DATETIME,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE db_sync_state (
    name VARCHAR(128) NOT NULL,
    value VARCHAR(128) NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE sudoku_matrix_completions (
    user_id CHAR(36) NOT NULL,
    matrix_type VARCHAR(32) NOT NULL,
    completed_on DATE DEFAULT (CURRENT_DATE) NOT NULL,
    CONSTRAINT sudoku_matrix_completions_matrix_type_check CHECK (matrix_type IN ('drivers', 'teams')),
    UNIQUE KEY sudoku_matrix_completions_unique (user_id, matrix_type, completed_on),
    CONSTRAINT sudoku_matrix_completions_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE favourite_drivers (
    user_id CHAR(36) NOT NULL,
    driver CHAR(36) NOT NULL,
    CONSTRAINT fdrivers_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fdrivers_driver_fk FOREIGN KEY (driver) REFERENCES driver(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE favourite_teams (
    user_id CHAR(36) NOT NULL,
    team CHAR(36) NOT NULL,
    CONSTRAINT fteam_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fteams_team_fk FOREIGN KEY (team) REFERENCES team(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE finish (
    pos INT NOT NULL,
    race CHAR(36),
    car_p CHAR(36),
    CONSTRAINT finish_race_fkey FOREIGN KEY (race) REFERENCES race(id) ON DELETE CASCADE,
    CONSTRAINT finish_car_p_fkey FOREIGN KEY (car_p) REFERENCES car_p(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE qualifying (
    pos INT NOT NULL,
    car_p CHAR(36),
    race CHAR(36),
    CONSTRAINT qualifying_car_p_fkey FOREIGN KEY (car_p) REFERENCES car_p(id) ON DELETE CASCADE,
    CONSTRAINT qualifying_race_fkey FOREIGN KEY (race) REFERENCES race(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
