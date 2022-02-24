package main

import (
	"errors"
	"fmt"
)

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

func main() {
	var cardDeck []Card
	cardDeck = append(cardDeck, NewCard(two, spades))
	cardDeck = append(cardDeck, NewCard(ace, hearts))
	cardDeck = append(cardDeck, NewCard(ten, diamonds))
	cardDeck = append(cardDeck, NewCard(jack, clubs))
	cardDeck = append(cardDeck, NewCard(king, spades))
	cardDeck = append(cardDeck, NewCard(nine, hearts))
	cardDeck = append(cardDeck, NewCard(seven, spades))
	cardDeck = append(cardDeck, NewCard(queen, diamonds))
	cardDeck = append(cardDeck, NewCard(ace, spades))
	for _, card := range cardDeck {
		fmt.Println(cardInText(card))
	}
	println()
	println("Biggest card in the deck is: ", cardInText(maxCard(cardDeck)))

}

func compareCards(cardOne Card, cardTwo Card) int {
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

func maxCard(cards []Card) Card {
	var temp Card
	temp.Value, temp.Suite = copyCard(cards[0])
	for _, card := range cards {
		if compareCards(temp, card) == 1 {
			temp.Value, temp.Suite = copyCard(card)
		}
	}
	return temp
}

func copyCard(source Card) (CardValue, CardSuit) {
	return source.Value, source.Suite
}

func checkIfCardIsValid(cardToCheck Card) bool {
	return cardToCheck.Suite >= 1 && cardToCheck.Suite <= 4 && cardToCheck.Value >= 2 && cardToCheck.Value <= 14
}
func NewCard(value int, suite int) Card {

	if suite < 1 && suite > 4 && value < 2 && value > 14 {
		err := errors.New("Invalid input!")
		fmt.Println(err)
		return Card{-5, -5}
	}
	return Card{value, suite}
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
		result = "ten of"
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
