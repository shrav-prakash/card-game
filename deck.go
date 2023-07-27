package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type card struct {
	suit string
	number string
}

type deck []card

func (d deck) printDeck() {
	if len(d) == 0 {
		fmt.Println("Deck is empty!")
		return
	}
	for _, card := range d {
		fmt.Println(card.number + " of " + card.suit)
	}
}

func createDeck() deck {
	var cards deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNums := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardNum := range cardNums {
			c := card{
				suit: cardSuit,
				number: cardNum,
			}
			cards = append(cards, c)
		}
	}
	return cards
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

func (d deck) storeDeck(filename string) error {
	var byteDeck []byte = []byte(d.deckToString())
	
	return os.WriteFile("decks/" + filename, byteDeck, 0666)

}

func readDeckFromFile(filename string) deck {
	cards, err := os.ReadFile("decks/" + filename)
	if err == nil {
		var newDeck []string = strings.Split(string(cards), ",")
		var d deck
		for _, c := range newDeck {
			newCard := card{
				suit: strings.Split(c, " of ")[1],
				number: strings.Split(c, " of ")[0],
			}
			d = append(d, newCard)
		}
		return d
	}
	fmt.Println("Error occured during file read: ", err)
	os.Exit(1)
	return make(deck, 0)
}

func (d deck) deckToString() string {
	var cards []string
	for _, card := range d {
		cards = append(cards, card.number + " of " + card.suit)
	}
	return strings.Join(cards, ",")
}

func (d deck) shuffle() {
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i] 
	}
}