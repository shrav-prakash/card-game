package main

import "fmt"

func main() {
	cards := createDeck()
	cards.shuffle()
	hand := cards.dealHand(10)
	fmt.Println("Current Hand:")
	hand.printHand()
	cards.storeDeck("deck1")
	fmt.Println("\n\nExisting Deck:")
	fmt.Println(readDeckFromFile("deck1"))
}