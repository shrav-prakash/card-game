package main

import "fmt"

type hand []card

func (h hand) printHand() {
	if len(h) == 0 {
		fmt.Println("Hand is empty!")
		return
	}
	for _, card := range h {
		fmt.Println(card.number + " of " + card.suit)
	}
}

func createHand(d deck) hand {
	var newHand hand
	for _, c := range d {
		newC := card{
			number: c.number,
			suit: c.suit,
		}
		newHand = append(newHand, newC)
	}
	return newHand
}