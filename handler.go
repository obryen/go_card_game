package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type createDeckResponse struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

// Initialize a single instance of MemoryDeckService
var memoryDeckService = NewMemoryDeckService()

// DeckHandler handles the requests for creating and opening decks
func DeckHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		shuffled := r.URL.Query().Get("shuffled") == "true"
		cardsParam := r.URL.Query().Get("cards")

		deck, err := memoryDeckService.CreateDeck(shuffled, cardsParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		responsePayload := createDeckResponse{
			ID:        deck.ID,
			Shuffled:  deck.Shuffled,
			Remaining: deck.Remaining,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responsePayload)

	case "GET":
		vars := mux.Vars(r)
		id := vars["id"]
		if deck, err := memoryDeckService.OpenDeck(id); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(deck)
		} else {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		vars := mux.Vars(r)
		id := vars["id"]
		countStr := r.URL.Query().Get("count")
		count, err := strconv.Atoi(countStr)
		if err != nil {
			http.Error(w, "Invalid count value", http.StatusBadRequest)
			return
		}
		if cards, err := memoryDeckService.DrawCards(id, count); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string][]Card{"cards": cards})
		} else {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
