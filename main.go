package main

import "fmt"

// import "fmt"

func main() {
	cards := createDeck()
	
	hand := cards.dealHand(10)
	fmt.Println("Current Hand:")
	hand.printHand()
	cards.storeDeck("deck1")
	fmt.Println("\n\nExisting Deck:")
	fmt.Println(readDeckFromFile("deck1"))
}