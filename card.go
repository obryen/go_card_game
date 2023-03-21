package main

import (
	"fmt"
	"strings"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func newFullDeck() []Card {
	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING", "ACE"}

	var cards []Card

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, Card{
				Value: value,
				Suit:  suit,
				Code:  fmt.Sprintf("%s%s", string(value[0]), string(suit[0])),
			})
		}
	}

	return cards
}

func newPartialDeck(codes []string) []Card {
	var cards []Card
	fullDeck := newFullDeck()

	for _, code := range codes {
		for _, card := range fullDeck {
			if card.Code == strings.ToUpper(code) {
				cards = append(cards, card)
				break
			}
		}
	}

	return cards
}
