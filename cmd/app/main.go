package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	repo "github.com/gopherzz/webchat/internal/repository"
	"github.com/gopherzz/webchat/internal/services"
	"github.com/gopherzz/webchat/internal/webchat"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func getEnv(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}

func run() error {
	port := fmt.Sprintf(":%s", getEnv("PORT", "8080"))
	databasePath := getEnv("DATABASE_PATH", "webchat.db")

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}

	repo := repo.NewRepository(db)
	services := services.NewServices(repo)

	hub := webchat.NewHub(services)
	go hub.Run()

	http.HandleFunc("/", webchat.WebSocketHandler(hub))

	log.Printf("listening on %s", port)
	return http.ListenAndServe(port, nil)
}
