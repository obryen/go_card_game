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

// POST request to create a new deck
func CreateDeckHandler(w http.ResponseWriter, r *http.Request) {
	shuffled := r.FormValue("shuffled") == "true"
	cardsParam := r.FormValue("cards")

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
}

// GET request to open an existing deck
func OpenDeckHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["deck_id"]
	if deck, err := memoryDeckService.OpenDeck(id); err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deck)
	} else {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

// POST request to draw cards from a deck
func DrawCardsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["deck_id"]
	countStr := r.FormValue("count")
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
}

// // Define the routes using the Gorilla Mux router
// func main() {
//     router := mux.NewRouter()
//     router.HandleFunc("/decks", CreateDeckHandler).Methods("POST")
//     router.HandleFunc("/decks/{deck_id}", OpenDeckHandler).Methods("GET")
//     router.HandleFunc("/decks/{deck_id}/cards", DrawCardsHandler).Methods("POST")
//     http.ListenAndServe(":8000", router)
// }
