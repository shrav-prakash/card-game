package main

import "fmt"

func main() {
	cards := createDeck()
	hand := cards.dealHand(5)
	fmt.Println("Current Hand:")
	hand.printDeck()
	fmt.Println("\n\nExisting Deck:")
	cards.printDeck()
}