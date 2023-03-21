package main

import (
	"strings"
	"testing"
)

func TestNewFullDeck(t *testing.T) {
	deck := newFullDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length to be 52, but got %d", len(deck))
	}

	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING", "ACE"}

	cardCount := make(map[string]int)

	for _, card := range deck {
		if !contains(suits, card.Suit) {
			t.Errorf("Invalid suit found: %s", card.Suit)
		}
		if !contains(values, card.Value) {
			t.Errorf("Invalid value found: %s", card.Value)
		}
		cardCount[card.Code]++
		if cardCount[card.Code] > 1 {
			t.Errorf("Duplicate card found: %s", card.Code)
		}
	}
}

func TestNewPartialDeck(t *testing.T) {
	codes := []string{"2S", "3C", "4D", "5H", "KC"}
	partialDeck := newPartialDeck(codes)

	if len(partialDeck) != len(codes) {
		t.Errorf("Expected partial deck length to be %d, but got %d", len(codes), len(partialDeck))
	}

	for _, code := range codes {
		found := false
		for _, card := range partialDeck {
			if card.Code == strings.ToUpper(code) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected card not found in partial deck: %s", code)
		}
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
