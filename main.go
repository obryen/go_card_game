package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/deck/new", DeckHandler).Methods("POST")
	router.HandleFunc("/deck/{id}", DeckHandler).Methods("GET")
	router.HandleFunc("/deck/{id}/draw", CardHandler).Methods("POST")

	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
