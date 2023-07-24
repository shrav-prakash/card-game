package main

import "fmt"

type deck []string

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func createDeck() deck {
	var cards deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNums := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardNum := range cardNums {
			cards = append(cards, cardNum + " of " + cardSuit)
		}
	}
	return cards
}