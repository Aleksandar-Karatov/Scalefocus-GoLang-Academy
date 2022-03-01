package main

import (
	"errors"
	"fmt"
)

func main() {
	var cardDeck []Card
	cardDeck = append(cardDeck, Card{two, spades})
	cardDeck = append(cardDeck, Card{ace, hearts})
	cardDeck = append(cardDeck, Card{ten, diamonds})
	cardDeck = append(cardDeck, Card{jack, clubs})
	cardDeck = append(cardDeck, Card{king, spades})
	cardDeck = append(cardDeck, Card{nine, hearts})
	cardDeck = append(cardDeck, Card{seven, spades})
	cardDeck = append(cardDeck, Card{queen, diamonds})
	cardDeck = append(cardDeck, Card{ace, spades})
	// for _, card := range cardDeck {
	// 	fmt.Println(cardInText(card))
	// }
	// fmt.Println()

	var f CardComparator
	f = func(cardOne Card, cardTwo Card) int {
		var result int
		if checkIfCardIsValid(cardOne) && checkIfCardIsValid(cardTwo) {
			if cardOne.Value > cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suite > cardTwo.Suite) {
				result = -1
			} else if cardOne.Value < cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suite < cardTwo.Suite) {
				result = 1
			} else {
				result = 0
			}
			return result
		}
		fmt.Println(errors.New("Invalid input!"))
		return -5
	}
	fmt.Println("Anonymously: ")
	fmt.Println("Biggest card in the deck is: ", cardInText(maxCard(cardDeck, f)))
	fmt.Println("With a reference: ")
	fmt.Println("Biggest card in the deck is: ", cardInText(maxCard(cardDeck, CardCompare)))

}

type CardValue = int
type CardSuit = int

const (
	two CardValue = iota + 2 // card values
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

const (
	clubs CardSuit = iota + 1 // card suits
	diamonds
	hearts
	spades
)

type Card struct {
	Value int
	Suite int
}

type CardComparator func(cOne Card, cTwo Card) int

func checkIfCardIsValid(cardToCheck Card) bool {
	return cardToCheck.Suite >= 1 && cardToCheck.Suite <= 4 && cardToCheck.Value >= 2 && cardToCheck.Value <= 14
}

func maxCard(cards []Card, comparatorFunc CardComparator) Card {
	var temp = cards[0]
	for _, card := range cards {
		check := comparatorFunc(temp, card)
		if check == 1 {
			temp = card
		} else if check == -5 {
			return Card{-5, -5}
		}
	}
	return temp
}
func cardInText(card Card) string {
	var result string
	switch card.Value {
	case two:
		result = "two of "
	case three:
		result = "three of "
	case four:
		result = "four of "
	case five:
		result = "five of "
	case six:
		result = "six of "
	case seven:
		result = "seven of "
	case eight:
		result = "eight of "
	case nine:
		result = "nine of "
	case ten:
		result = "ten of "
	case jack:
		result = "jack of "
	case queen:
		result = "queen of "
	case king:
		result = "king of "
	case ace:
		result = "ace of "
	}

	switch card.Suite {
	case clubs:
		result += "clubs"
	case diamonds:
		result += "diamonds"
	case hearts:
		result += "hearts"
	case spades:
		result += "spades"
	}
	return result
}

func CardCompare(cardOne Card, cardTwo Card) int {
	var result int
	if checkIfCardIsValid(cardOne) && checkIfCardIsValid(cardTwo) {
		if cardOne.Value > cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suite > cardTwo.Suite) {
			result = -1
		} else if cardOne.Value < cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suite < cardTwo.Suite) {
			result = 1
		} else {
			result = 0
		}
		return result
	}
	fmt.Println(errors.New("Invalid input!"))
	return -5
}
