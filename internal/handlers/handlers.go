package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Dmitrii30002/Quote-library/internal/models"
	"github.com/Dmitrii30002/Quote-library/internal/repository"
	"github.com/Dmitrii30002/Quote-library/pkg/database"
)

func PostQuote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	type Message struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	var m Message

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if m.Author == "" {
		log.Println("[DEBUG] author is empty")
		http.Error(w, "author is empty", http.StatusBadRequest)
		return
	}

	if m.Quote == "" {
		log.Println("[DEBUG] quote is empty")
		http.Error(w, "quote is empty", http.StatusBadRequest)
		return
	}
	repA := repository.NewAuthorRepository(database.DB)
	a := &models.Author{Name: m.Author}
	err = repA.Create(a)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	a, err = repA.GetByName(m.Author)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	repQ := repository.NewQuoteRepository(database.DB)
	q := &models.Quote{Text: m.Quote, Author_ID: a.ID}
	err = repQ.Create(q)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rep := repository.NewQuoteRepository(database.DB)
	quotes, err := rep.GetAll()
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(quotes)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rep := repository.NewQuoteRepository(database.DB)
	quote, err := rep.GetRandom()
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
	w.WriteHeader(http.StatusOK)
}

func GetQuotesByAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	author := r.URL.Query().Get("author")
	if author == "" {
		log.Println("[DEBUG] Empty author")
		http.Error(w, "Parametr author can't be empty", http.StatusBadRequest)
		return
	}

	rep := repository.NewQuoteRepository(database.DB)
	quotes, err := rep.GetByAuthorName(author)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(quotes)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteQuoteByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	path := strings.TrimPrefix(r.URL.Path, "/quotes/")
	id, err := strconv.Atoi(path)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	rep := repository.NewQuoteRepository(database.DB)
	err = rep.Delete(id)
	if err != nil {
		log.Printf("[DEBUG] %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
}
