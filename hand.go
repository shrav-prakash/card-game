package main

import "fmt"

type hand []string

func (h hand) printHand() {
	if len(h) == 0 {
		fmt.Println("Hand is empty!")
		return
	}
	for _, card := range h {
		fmt.Println(card)
	}
}

func createHand(d deck) hand {
	var newHand hand
	for _, card := range d {
		newHand = append(newHand, card)
	}
	return newHand
}