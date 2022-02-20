package main

import (
	"errors"
	"fmt"
)

type CardValue = int
type CardSuit = int

const (
	two CardValue = iota // card values
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
	clubs CardSuit = iota // card suits
	diamonds
	hearts
	spades
)

const (
	errorCodeValue            int = -5 // error codes for the different possible errors
	errorCodeSuit             int = -10
	errorCodeBothSuitAndValue int = -15
)

func main() {
	fmt.Println(compareCards(two, spades, two, spades)) // tests
	fmt.Println(compareCards(two, spades, two, hearts))
	fmt.Println(compareCards(two, hearts, two, spades))
	fmt.Println(compareCards(jack, clubs, two, spades))
	fmt.Println(compareCards(15, clubs, 222, spades))
	fmt.Println(compareCards(15, 123, 222, 123))
	fmt.Println(compareCards(nine, 5555, eight, spades))
	fmt.Println(compareCards(ten, 5555, ten, 123124))
	fmt.Println(compareCards(ace, clubs, ace, spades))
	fmt.Println(compareCards(queen, diamonds, queen, hearts))
	fmt.Println(compareCards(ace, 5555, ace, spades))
	fmt.Println(compareCards(eight, diamonds, jack, clubs))
	fmt.Println(compareCards(seven, diamonds, seven, diamonds))

}

func compareCards(cardOneVal int, cardOneSuit int, cardTwoVal int, cardTwoSuit int) int {

	if (!checkIfWithinRange(cardOneVal, two, ace) || !checkIfWithinRange(cardTwoVal, two, ace)) && // this if and these else if statements check if the data is valid
		(!checkIfWithinRange(cardOneSuit, clubs, spades) || !checkIfWithinRange(cardTwoSuit, clubs, spades)) { // check if both card value and card suit are valid
		fmt.Println(errors.New("Both card value and suit are invalid!"))
		return errorCodeBothSuitAndValue
	} else if !checkIfWithinRange(cardOneVal, two, ace) || !checkIfWithinRange(cardTwoVal, two, ace) { // check if only card value is valid
		fmt.Println(errors.New("Card value is invalid!"))
		return errorCodeValue
	} else if !checkIfWithinRange(cardOneSuit, clubs, spades) || !checkIfWithinRange(cardTwoSuit, clubs, spades) { // check if card suit is valid
		fmt.Println(errors.New("Card suit is invalid!"))
		return errorCodeSuit
	}

	if cardOneVal > cardTwoVal || (cardOneVal == cardTwoVal && cardOneSuit > cardTwoSuit) {
		return -1
	} else if cardOneVal < cardTwoVal || (cardOneVal == cardTwoVal && cardOneSuit < cardTwoSuit) {
		return 1
	} else {
		return 0
	}
}

func checkIfWithinRange(number int, min int, max int) bool { // repetetive check made into a function so that code isn`t repeated
	return number >= min && number <= max
}
