package main

import "fmt"

type deck []string

func (d deck) printDeck() {
	if len(d) == 0 {
		fmt.Println("Deck is empty!")
		return
	}
	for _, card := range d {
		fmt.Println(card)
	}
}

func createDeck() *deck {
	var cards deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNums := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardNum := range cardNums {
			cards = append(cards, cardNum + " of " + cardSuit)
		}
	}
	return &cards
}

func (d *deck) dealHand(size int) hand {
	var currHand hand
	var cards deck = *d

	if size >= len(*d) {
		currHand = createHand(*d)
		*d = make(deck, 0)
	} else {
		currHand = createHand(cards[:size])
		*d = cards[size:]
	}	

	return currHand
}