package cardgame

import (
	"errors"
	"fmt"
)

type CardValue = int
type CardSuit = int

const (
	Two CardValue = iota + 2 // card values
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	Clubs CardSuit = iota + 1 // card suits
	Diamonds
	Hearts
	Spades
)

type Card struct {
	Value int
	Suite int
}

type CardComparator func(cOne Card, cTwo Card) int

func CheckIfCardIsValid(cardToCheck Card) bool {
	return cardToCheck.Suite >= 1 && cardToCheck.Suite <= 4 && cardToCheck.Value >= 2 && cardToCheck.Value <= 14
}

func MaxCard(cards []Card, comparatorFunc CardComparator) Card {
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
func CardInText(card Card) string {
	var result string
	switch card.Value {
	case Two:
		result = "two of "
	case Three:
		result = "three of "
	case Four:
		result = "four of "
	case Five:
		result = "five of "
	case Six:
		result = "six of "
	case Seven:
		result = "seven of "
	case Eight:
		result = "eight of "
	case Nine:
		result = "nine of "
	case Ten:
		result = "ten of "
	case Jack:
		result = "jack of "
	case Queen:
		result = "queen of "
	case King:
		result = "king of "
	case Ace:
		result = "ace of "
	}

	switch card.Suite {
	case Clubs:
		result += "clubs"
	case Diamonds:
		result += "diamonds"
	case Hearts:
		result += "hearts"
	case Spades:
		result += "spades"
	}
	return result
}

func CardCompare(cardOne Card, cardTwo Card) int {
	var result int
	if CheckIfCardIsValid(cardOne) && CheckIfCardIsValid(cardTwo) {
		if cardOne.Value > cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suite > cardTwo.Suite) {
			result = -1
		} else if cardOne.Value < cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suite < cardTwo.Suite) {
			result = 1
		} else {
			result = 0
		}
		return result
	}
	fmt.Println(errors.New("invalid input"))
	return -5
}
