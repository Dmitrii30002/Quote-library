package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dmitrii30002/Quote-library/config"
	"github.com/Dmitrii30002/Quote-library/internal/handlers"
	"github.com/Dmitrii30002/Quote-library/internal/migrations"
	"github.com/Dmitrii30002/Quote-library/pkg/database"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			if _, ok := r.URL.Query()["author"]; ok {
				handlers.GetQuotesByAuthor(w, r)
			} else {
				handlers.PostQuote(w, r)
			}
		case http.MethodGet:
			handlers.GetQuotes(w, r)
		}
	})

	mux.HandleFunc("/quotes/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			handlers.DeleteQuoteByID(w, r)
		}
	})

	mux.HandleFunc("/quotes/random", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetQuotes(w, r)
		}
	})

	cfg, err := config.New(".env", "config/config.json")
	if err != nil {
		log.Fatalf("[ERROR] Config loading error: %v", err)
	}

	db, err := database.New(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.DataBase.User,
		cfg.DataBase.Pwd,
		cfg.DataBase.Name,
		cfg.DataBase.Host,
		cfg.DataBase.Port,
		cfg.DataBase.Sslmode,
	))
	if err != nil {
		log.Fatalf("[ERROR] DB connection error: %v", err)
	}

	err = migrations.Migrate(db)
	if err != nil {
		log.Fatalf("[ERROR] Migration error: %v", err)
	}

	log.Printf("[INFO] Server was launched on %s:%s", cfg.Server.Host, cfg.Server.Port)
	err = http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, mux)
	if err != nil {
		log.Fatalf("[ERROR] Server startup error: %v", err)
	}
}
