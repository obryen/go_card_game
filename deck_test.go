package main

import (
	"testing"
)

func TestCreateDeck(t *testing.T) {
	service := NewMemoryDeckService()

	// Test creating a full deck
	deck, err := service.CreateDeck(false, "")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if deck.Remaining != 52 {
		t.Errorf("Expected 52 cards, got %d", deck.Remaining)
	}

	// Test creating a partial deck
	deck, err = service.CreateDeck(false, "AS,KH")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if deck.Remaining != 2 {
		t.Errorf("Expected 2 cards, got %d", deck.Remaining)
	}
}

func TestOpenDeck(t *testing.T) {
	service := NewMemoryDeckService()
	deck, _ := service.CreateDeck(false, "")

	// Test opening an existing deck
	openedDeck, err := service.OpenDeck(deck.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if openedDeck.ID != deck.ID {
		t.Errorf("Expected deck ID %s, got %s", deck.ID, openedDeck.ID)
	}

}

func TestDrawCards(t *testing.T) {
	service := NewMemoryDeckService()
	deck, _ := service.CreateDeck(false, "")

	// Test drawing cards from a deck
	drawnCards, err := service.DrawCards(deck.ID, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(drawnCards) != 5 {
		t.Errorf("Expected 5 drawn cards, got %d", len(drawnCards))
	}
	if deck.Remaining != 47 {
		t.Errorf("Expected 47 cards remaining, got %d", deck.Remaining)
	}

}
