package main

import (
	"RacingGuru/internal/controller/handlers"
	"RacingGuru/internal/controller/server"
	"RacingGuru/internal/repository"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	//dbURL := os.Getenv("DATABASE_URL")
	//serverAddr := os.Getenv("SERVER_ADDR")
	serverAddr := "localhost:8080"
	dbURL := "postgresql://postgres:1337@localhost:5432/races"
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer pool.Close()
	repo := repository.NewRepository(pool)
	handler := handlers.NewHandler(repo)
	logger := log.Default()
	s := server.NewServer(serverAddr, handler.Routes(), logger)
	log.Fatal(s.Run())
}
