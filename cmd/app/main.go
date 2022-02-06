package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gopherzz/cucumberdb"
	"github.com/gopherzz/webchat/internal/repository"
	"github.com/gopherzz/webchat/internal/services"
	"github.com/gopherzz/webchat/internal/webchat"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
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
	if err := godotenv.Load(); err != nil {
		return err
	}

	port := fmt.Sprintf(":%s", getEnv("PORT", "8080"))
	databasePath := getEnv("DATABASE_PATH", "webchat.db")

	db := cucumberdb.New()
	if err := db.Load(databasePath); err != nil {
		return err
	}

	repo := repository.NewRepository(db)
	services := services.NewServices(repo)

	poll := webchat.NewPoll(services)
	go poll.Run()

	http.HandleFunc("/messages", poll.GetMessages)
	http.HandleFunc("/send", poll.SendMessage)

	log.Printf("listening on %s", port)
	return http.ListenAndServe(port, nil)
}
