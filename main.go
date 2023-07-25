package main

import "fmt"

func main() {
	cards := createDeck()
	cards.shuffle()
	hand := cards.dealHand(10)
	fmt.Println("Starting Deck:")
	cards.printDeck();
	fmt.Println("\n\nCurrent Hand:")
	hand.printHand()
	cards.storeDeck("deck1")
	fmt.Println("\n\nExisting Deck:")
	fmt.Println(readDeckFromFile("deck1"))
}