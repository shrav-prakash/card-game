package main

import (
	"fmt"
	"sort"
)

type mappedCard struct {
	suit string
	number int
}

type hand []card

type mappedHand []mappedCard

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

func (h *hand) addToHand(c card) {
	*h = append(*h, c)
}

func compareHands(player hand, comp hand) int {
	ph := mapHand(player)
	ch := mapHand(comp)
	sort.SliceStable(ph, func(i, j int) bool { return ph[i].number < ph[j].number})
	sort.SliceStable(ch, func(i, j int) bool { return ch[i].number < ch[j].number})
	playerScore := calcScore(ph, true)
	compScore := calcScore(ch, false)
	if playerScore > compScore {
		return 1
	} else if playerScore < compScore {
		return -1
	} else {
		return tieBreaker(ph, ch, playerScore)
	}
}

func mapHand(h hand) mappedHand {
	numMap := map[string]int{
		"Two": 2,
		"Three": 3,
		"Four": 4,
		"Five": 5,
		"Six": 6,
		"Seven": 7,
		"Eight": 8,
		"Nine": 9,
		"Jack": 10,
		"Queen": 11,
		"King": 12,
		"Ace": 13,
	}
	var mh mappedHand
	for _, card := range h {
		mc := mappedCard{
			suit: card.suit,
			number: numMap[card.number],
		}
		mh = append(mh, mc)
	}
	return mh
}

func calcScore(h mappedHand, isPlayer bool) int {
	c1 := h[0]
	c2 := h[1]
	c3 := h[2]
	if (c1.number == c2.number - 1 && c2.number == c3.number - 1) || (c1.number == 13 && c2.number == 2 && c3.number == 3) {
		if(c1.suit == c2.suit && c2.suit == c3.suit) {
			if(isPlayer) {
				fmt.Println("\n\nYou have gotten a Straight Flush!")
			} else {
				fmt.Println("Dealer has gotten a Straight Flush!")
			}
			return 5
		} else {
			if(isPlayer) {
				fmt.Println("\n\nYou have gotten a Straight!")
			} else {
				fmt.Println("Dealer has gotten a Straight!")
			}
			return 3
		}
	} else if (c1.number == c2.number && c2.number == c3.number) {
		if(isPlayer) {
			fmt.Println("\n\nYou have gotten a Three of a Kind!")
		} else {
			fmt.Println("Dealer has gotten a Three of a Kind!")
		}
		return 4
	} else if(c1.suit == c2.suit && c2.suit == c3.suit) {
		if(isPlayer) {
			fmt.Println("\n\nYou have gotten a Flush!")
		} else {
			fmt.Println("Dealer has gotten a Flush!")
		}
		return 2
	} else if(c1.number == c2.number || c2.number == c3.number || c3.number == c1.number) {
		if(isPlayer) {
			fmt.Println("\n\nYou have gotten a Pair!")
		} else {
			fmt.Println("Dealer has gotten a Pair!")
		}
		return 1
	} else {
		if(isPlayer) {
			fmt.Println("\n\nYou have gotten a High Card!")
		} else {
			fmt.Println("Dealer has gotten a High Card!")
		}
		return 0
	}
}

func tieBreaker(ph mappedHand, ch mappedHand, score int) int {
	switch score {
	case 5:
		return checkHigher(ph[1], ch[1])
	case 4:
		return checkHigher(ph[0], ch[0])
	case 3:
		return checkHigher(ph[1], ch[1])
	case 2:
		res := checkHigher(ph[2], ch[2])
		if res == 0 {
			res = checkHigher(ph[1], ch[1])
			if res == 0 {
				res = checkHigher(ph[0], ch[0]) 
				if res == 0 {
					return 0
				}
			}
		}
		return res
	case 1:
		res := checkHigher(ph[1], ch[1])
		if res == 0 {
			var nuqp int
			if ph[0].number != ph[1].number {
				nuqp = ph[0].number
			} else {
				nuqp = ph[2].number
			}

			var nuqc int
			if ch[0].number != ch[1].number {
				nuqc = ch[0].number
			} else {
				nuqc = ch[2].number
			}

			if nuqp > nuqc {
				return 1
			} else if nuqp < nuqc {
				return -1
			} else {
				return 0
			}
			
		} 
		return res
	case 0:
		res := checkHigher(ph[2], ch[2])
		if res == 0 {
			res = checkHigher(ph[1], ch[1])
			if res == 0 {
				res = checkHigher(ph[0], ch[0]) 
				if res == 0 {
					return 0
				}
			}
		}
		return res
	}
	return 0
}

func checkHigher(c1 mappedCard, c2 mappedCard) int {
	if c1.number > c2.number {
			return 1
	} else if c1.number < c2.number {
		return -1
	} else {
		return 0
	}
}