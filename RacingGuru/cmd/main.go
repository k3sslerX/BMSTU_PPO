package main

import (
	"RacingGuru/internal/auth"
	"RacingGuru/internal/controller/handlers"
	"RacingGuru/internal/controller/server"
	"RacingGuru/internal/logger"
	"RacingGuru/internal/repository"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
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

	appLogger.Info("connecting to database")
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		appLogger.Errorf("database connection failed: %v", err)
		os.Exit(1)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool)
	handler := handlers.NewHandler(repo, appLogger)
	s := server.NewServer(serverAddr, handler.Routes(), appLogger)
	if err = s.Run(); err != nil {
		appLogger.Errorf("server stopped with error: %v", err)
		os.Exit(1)
	}
}
