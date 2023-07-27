package main

import "fmt"

func main() {
	cards := createDeck()
	cards.shuffle()
	var playerHand hand
	var dealerHand hand
	i := 0
	for i < 3 {
		c := cards.dealHand(1)[0]
		playerHand.addToHand(c)
		c = cards.dealHand(1)[0]
		dealerHand.addToHand(c)
		i += 1
	}
	fmt.Println("\n\nYour hand: ")
	playerHand.printHand();
	fmt.Println("\n\nDealer's hand: ")
	dealerHand.printHand()
	res := compareHands(playerHand, dealerHand)
	if(res == 1) {
		fmt.Println("\n\nPlayer has won!")
	} else if (res == -1) {
		fmt.Println("\n\nDealer has won!")
	} else {
		fmt.Println("\n\nIt is a tie!")
	}
}