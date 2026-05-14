package main

import (
	"RacingGuru/internal/auth"
	"RacingGuru/internal/controller/handlers"
	"RacingGuru/internal/controller/server"
	"RacingGuru/internal/logger"
	mysqlrepo "RacingGuru/internal/repository/mysql"
	postgresqlrepo "RacingGuru/internal/repository/postgresql"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()
	dbms := os.Getenv("DBMS")
	serverAddr := os.Getenv("SERVER_ADDR")
	logFileName := os.Getenv("LOG_FILE")
	logLevel := os.Getenv("LOG_LEVEL")

	logOutput := os.Stdout
	var err error
	var logFile *os.File
	if logFileName != "" {
		logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
			os.Exit(1)
		}
		defer func() {
			_ = logFile.Close()
		}()
		logOutput = logFile
	}

	appLogger, err := logger.New(logLevel, logOutput)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	if err = auth.ValidateJWTSecretKey(); err != nil {
		appLogger.Errorf("jwt configuration failed: %v", err)
		os.Exit(1)
	}

	dbURL, err := databaseURLForDBMS(dbms)
	if err != nil {
		appLogger.Errorf("database configuration failed: %v", err)
		os.Exit(1)
	}

	normalizedDBMS := normalizeDBMS(dbms)
	appLogger.Info("connecting to database", "dbms", normalizedDBMS)
	repo, closeRepo, err := openRepository(ctx, normalizedDBMS, dbURL)
	if err != nil {
		appLogger.Errorf("database connection failed: %v", err)
		os.Exit(1)
	}
	defer closeRepo()

	handler := handlers.NewHandler(repo, appLogger)
	s := server.NewServer(serverAddr, handler.Routes(), appLogger)
	if err = s.Run(); err != nil {
		appLogger.Errorf("server stopped with error: %v", err)
		os.Exit(1)
	}
}

func databaseURLForDBMS(dbms string) (string, error) {
	envName := ""
	switch normalizeDBMS(dbms) {
	case "postgres":
		envName = "APP_POSTGRES_DATABASE_URL"
	case "mysql":
		envName = "APP_MYSQL_DATABASE_URL"
	default:
		return "", fmt.Errorf("unsupported DBMS %q; use postgres or mysql", dbms)
	}

	dbURL := strings.TrimSpace(os.Getenv(envName))
	if dbURL == "" {
		return "", fmt.Errorf("%s is required", envName)
	}

	return dbURL, nil
}

func openRepository(ctx context.Context, dbms, dbURL string) (handlers.Repo, func(), error) {
	if strings.TrimSpace(dbURL) == "" {
		return nil, nil, fmt.Errorf("database URL is required")
	}

	switch normalizeDBMS(dbms) {
	case "postgres":
		pool, err := pgxpool.Connect(ctx, dbURL)
		if err != nil {
			return nil, nil, err
		}
		return postgresqlrepo.NewRepository(pool), pool.Close, nil
	case "mysql":
		db, err := sql.Open("mysql", dbURL)
		if err != nil {
			return nil, nil, err
		}
		if err := configureSQLPool(db); err != nil {
			_ = db.Close()
			return nil, nil, err
		}
		if err := db.PingContext(ctx); err != nil {
			_ = db.Close()
			return nil, nil, err
		}
		return mysqlrepo.NewRepository(db), func() { _ = db.Close() }, nil
	default:
		return nil, nil, fmt.Errorf("unsupported DBMS %q; use postgres or mysql", dbms)
	}
}

func normalizeDBMS(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "postgresql", "pg", "postgres":
		return "postgres"
	case "mysql", "mariadb":
		return "mysql"
	default:
		return strings.ToLower(strings.TrimSpace(value))
	}
}

func configureSQLPool(db *sql.DB) error {
	maxOpen, err := envInt("DB_MAX_OPEN_CONNS", 25)
	if err != nil {
		return err
	}
	maxIdle, err := envInt("DB_MAX_IDLE_CONNS", 25)
	if err != nil {
		return err
	}
	maxLifetime, err := envDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	db.SetConnMaxLifetime(maxLifetime)

	return nil
}

func envInt(name string, fallback int) (int, error) {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("%s must be integer: %w", name, err)
	}
	if value < 0 {
		return 0, fmt.Errorf("%s must be non-negative", name)
	}

	return value, nil
}

func envDuration(name string, fallback time.Duration) (time.Duration, error) {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return fallback, nil
	}

	value, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("%s must be duration, e.g. 5m: %w", name, err)
	}
	if value < 0 {
		return 0, fmt.Errorf("%s must be non-negative", name)
	}

	return value, nil
}
