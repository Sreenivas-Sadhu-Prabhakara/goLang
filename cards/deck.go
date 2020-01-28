package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}
	cardSuits := []string{"hearts", "clubs", "spades", "diamonds"}

	for _, suit := range cardSuits {

		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
