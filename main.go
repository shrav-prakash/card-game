package main

import "fmt"

func main() {
	cards := createDeck()
	hand := cards.dealHand(10)
	fmt.Println("Current Hand:")
	hand.printHand()
	fmt.Println("\n\nExisting Deck:")
	cards.printDeck()
}