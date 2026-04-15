package main

import (
	"RacingGuru/internal/bootstrap"
	"RacingGuru/internal/controller/handlers"
	"RacingGuru/internal/controller/server"
	"RacingGuru/internal/repository"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	//dbURL := os.Getenv("DATABASE_URL")
	//serverAddr := os.Getenv("SERVER_ADDR")
	logFileName := os.Getenv("LOG_FILE")
	logFile := os.Stdout
	var err error
	if logFileName != "" {
		logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
	serverAddr := "localhost:8080"
	dbURL := "postgresql://postgres:1337@localhost:5432/races"
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	repo := repository.NewRepository(pool)
	if err = bootstrap.EnsureAdmin(context.Background(), repo); err != nil {
		log.Fatal(err)
	}
	handler := handlers.NewHandler(repo)
	logger := log.New(logFile, "", log.LstdFlags)
	s := server.NewServer(serverAddr, handler.Routes(), logger)
	log.Fatal(s.Run())
}
