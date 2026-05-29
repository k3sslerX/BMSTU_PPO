package main

import (
	"RacingGuru/internal/auth"
	"RacingGuru/internal/controller/handlers"
	"RacingGuru/internal/controller/server"
	"RacingGuru/internal/logger"
	mongodbrepo "RacingGuru/internal/repository/mongodb"
	postgresqlrepo "RacingGuru/internal/repository/postgresql"
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	case "mongo":
		envName = "APP_MONGO_DATABASE_URL"
	default:
		return "", fmt.Errorf("unsupported DBMS %q; use postgres or mongo", dbms)
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
	case "mongo":
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
		if err != nil {
			return nil, nil, err
		}
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			_ = client.Disconnect(ctx)
			return nil, nil, err
		}
		dbName, err := mongoDatabaseName(dbURL)
		if err != nil {
			_ = client.Disconnect(ctx)
			return nil, nil, err
		}
		repo := mongodbrepo.NewRepository(client.Database(dbName))
		if err := repo.EnsureIndexes(ctx); err != nil {
			_ = client.Disconnect(ctx)
			return nil, nil, err
		}
		return repo, func() { _ = client.Disconnect(context.Background()) }, nil
	default:
		return nil, nil, fmt.Errorf("unsupported DBMS %q; use postgres or mongo", dbms)
	}
}

func normalizeDBMS(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "postgresql", "pg", "postgres":
		return "postgres"
	case "mongodb", "mongo":
		return "mongo"
	default:
		return strings.ToLower(strings.TrimSpace(value))
	}
}

func mongoDatabaseName(dbURL string) (string, error) {
	parsed, err := url.Parse(dbURL)
	if err != nil {
		return "", err
	}

	dbName := strings.Trim(parsed.Path, "/")
	if dbName != "" {
		return dbName, nil
	}

	dbName = strings.TrimSpace(os.Getenv("APP_MONGO_DATABASE_NAME"))
	if dbName != "" {
		return dbName, nil
	}

	return "", fmt.Errorf("mongo database name is required in APP_MONGO_DATABASE_URL path or APP_MONGO_DATABASE_NAME")
}
