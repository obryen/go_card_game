package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/decks", CreateDeckHandler).Methods("POST")
	router.HandleFunc("/decks/{deck_id}", OpenDeckHandler).Methods("GET")
	router.HandleFunc("/decks/{deck_id}/cards", DrawCardsHandler).Methods("POST")
	http.ListenAndServe(":8000", router)

	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
