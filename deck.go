package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards,omitempty"`
}

type MemoryDeckService struct {
	decks map[string]*Deck
}

func NewMemoryDeckService() *MemoryDeckService {
	return &MemoryDeckService{
		decks: make(map[string]*Deck),
	}
}

func (m *MemoryDeckService) CreateDeck(shuffled bool, cardsParam string) (*Deck, error) {
	var cards []Card

	if cardsParam == "" {
		cards = newFullDeck()
	} else {
		cardsParamArr := strings.Split(cardsParam, ",")
		cards = newPartialDeck(cardsParamArr)
	}

	if shuffled {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	}

	deck := &Deck{
		ID:        fmt.Sprintf("%d", rand.Int()),
		Shuffled:  shuffled,
		Remaining: len(cards),
		Cards:     cards,
	}

	m.decks[deck.ID] = deck

	return deck, nil
}

func (m *MemoryDeckService) OpenDeck(id string) (*Deck, error) {
	deck, found := m.decks[id]
	if !found {
		return nil, errors.New("deck not found")
	}

	return deck, nil
}

func (m *MemoryDeckService) DrawCards(deckID string, count int) ([]Card, error) {
	deck, err := m.OpenDeck(deckID)
	if err != nil {
		return nil, err
	}

	if count > deck.Remaining {
		return nil, errors.New("not enough cards in the deck")
	}

	drawnCards := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	deck.Remaining = len(deck.Cards)

	return drawnCards, nil
}
