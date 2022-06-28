package cardgame_test

import (
	"fmt"
	"testing"
	"week6Lecture17Task/week6/lecture17/cardgame"
)

func TestMaxCard(t *testing.T) {
	var cardDeck []cardgame.Card
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Two, cardgame.Spades})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Ace, cardgame.Hearts})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Ten, cardgame.Diamonds})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Jack, cardgame.Clubs})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.King, cardgame.Spades})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Nine, cardgame.Hearts})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Seven, cardgame.Spades})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Queen, cardgame.Diamonds})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Ace, cardgame.Spades})

	maxCard := cardgame.MaxCard(cardDeck, cardgame.CardCompare)
	aceOfSpades := cardgame.Card{cardgame.Ace, cardgame.Spades}
	if maxCard != aceOfSpades {
		t.Error("Max card was incorrect, got" + cardgame.CardInText(maxCard) + ", want:" + "ace of spades")
	}
	cardDeck = append(cardDeck, cardgame.Card{1234, 1234})
	maxCard = cardgame.MaxCard(cardDeck, cardgame.CardCompare)
	errCard := cardgame.Card{-5, -5}
	if maxCard != errCard {
		t.Error("Error card was incorrect, got" + cardgame.CardInText(maxCard) + ", want: err: Card{-5 -5}")
	}

}
func TestCardCompare(t *testing.T) {
	var card1 cardgame.Card = cardgame.Card{cardgame.King, cardgame.Spades}
	var card2 cardgame.Card = cardgame.Card{cardgame.King, cardgame.Spades}
	var card3 cardgame.Card = cardgame.Card{cardgame.Ace, cardgame.Spades}
	var card4 cardgame.Card = cardgame.Card{cardgame.Two, cardgame.Spades}
	var card5 cardgame.Card = cardgame.Card{cardgame.King, cardgame.Hearts}
	var card6 cardgame.Card = cardgame.Card{124123, 123142}
	var card7 cardgame.Card = cardgame.Card{-123142, -12314}
	var c1c2 = cardgame.CardCompare(card1, card2)
	if c1c2 != 0 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c1c2) + ", want: 0")
	}
	var c1c3 = cardgame.CardCompare(card1, card3)
	if c1c3 != 1 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c1c3) + ", want: 1")
	}
	var c1c4 = cardgame.CardCompare(card1, card4)
	if c1c4 != -1 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c1c4) + ", want: -1")
	}
	var c1c5 = cardgame.CardCompare(card1, card5)
	if c1c5 != -1 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c1c5) + ", want: -1")
	}
	var c1c6 = cardgame.CardCompare(card1, card6)
	if c1c6 != -5 {
		t.Error("Value was incorrect, got" + fmt.Sprint(c1c6) + ", want: -5")
	}
	var c1c7 = cardgame.CardCompare(card1, card7)
	if c1c7 != -5 {
		t.Error("Value was incorrect, got" + fmt.Sprint(c1c7) + ", want: -5")
	}

	var c2c1 = cardgame.CardCompare(card2, card1)
	var c3c1 = cardgame.CardCompare(card3, card1)
	var c4c1 = cardgame.CardCompare(card4, card1)
	var c5c1 = cardgame.CardCompare(card5, card1)
	var c6c1 = cardgame.CardCompare(card6, card1)
	var c7c1 = cardgame.CardCompare(card7, card1)

	if c2c1 != 0 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c2c1) + ", want: 0")
	}
	if c3c1 != -1 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c3c1) + ", want: -1")
	}
	if c4c1 != 1 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c4c1) + ", want: 1")
	}
	if c5c1 != 1 {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(c5c1) + ", want: 1")
	}
	if c6c1 != -5 {
		t.Error("Value was incorrect, got" + fmt.Sprint(c6c1) + ", want: -5")
	}
	if c7c1 != -5 {
		t.Error("Value was incorrect, got" + fmt.Sprint(c7c1) + ", want: -5")
	}
}
func TestCardToText(t *testing.T) {
	var card cardgame.Card = cardgame.Card{cardgame.King, cardgame.Spades}
	var cardInText = cardgame.CardInText(card)
	if cardInText != "king of spades" {
		t.Error("String was incorrect, got" + cardInText + ", want:" + "king of spades")
	}

	var cardDeck []cardgame.Card
	var cardDeckNames []string
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Two, cardgame.Spades})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Three, cardgame.Diamonds})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Four, cardgame.Clubs})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Five, cardgame.Spades})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Six, cardgame.Hearts})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Seven, cardgame.Diamonds})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Eight, cardgame.Clubs})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Nine, cardgame.Spades})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Ten, cardgame.Spades})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Jack, cardgame.Diamonds})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Queen, cardgame.Hearts})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.King, cardgame.Clubs})
	cardDeck = append(cardDeck, cardgame.Card{cardgame.Ace, cardgame.Hearts})

	cardDeckNames = append(cardDeckNames, "two of spades")
	cardDeckNames = append(cardDeckNames, "three of diamonds")
	cardDeckNames = append(cardDeckNames, "four of clubs")
	cardDeckNames = append(cardDeckNames, "five of spades")
	cardDeckNames = append(cardDeckNames, "six of hearts")
	cardDeckNames = append(cardDeckNames, "seven of diamonds")
	cardDeckNames = append(cardDeckNames, "eight of clubs")
	cardDeckNames = append(cardDeckNames, "nine of spades")
	cardDeckNames = append(cardDeckNames, "ten of spades")
	cardDeckNames = append(cardDeckNames, "jack of diamonds")
	cardDeckNames = append(cardDeckNames, "queen of hearts")
	cardDeckNames = append(cardDeckNames, "king of clubs")
	cardDeckNames = append(cardDeckNames, "ace of hearts")

	for i := 0; i < 13; i++ {
		currentCardInText := cardgame.CardInText(cardDeck[i])
		if currentCardInText != cardDeckNames[i] {

			t.Error("String was incorrect, got" + currentCardInText + ", want:" + cardDeckNames[i])

		}
	}

}
func TestCheckIfCardIsValid(t *testing.T) {
	var card1 cardgame.Card = cardgame.Card{cardgame.King, cardgame.Spades}
	check1 := cardgame.CheckIfCardIsValid(card1)
	var card2 cardgame.Card = cardgame.Card{12312, 123142}
	check2 := cardgame.CheckIfCardIsValid(card2)

	var card3 cardgame.Card = cardgame.Card{-12312, -123142}
	check3 := cardgame.CheckIfCardIsValid(card3)

	if !check1 {
		t.Error("String was incorrect, got" + fmt.Sprint(check1) + ", want: TRUE")
	}
	if check2 {
		t.Error("String was incorrect, got" + fmt.Sprint(check2) + ", want: FALSE")

	}
	if check3 {
		t.Error("String was incorrect, got" + fmt.Sprint(check3) + ", want: FALSE")

	}

}

func TestCard(t *testing.T) {
	card := cardgame.Card{cardgame.Ace, cardgame.Clubs}
	cardWithInt := cardgame.Card{14, 1}
	if card != cardWithInt {
		t.Error("Value setting incorrect, got" + cardgame.CardInText(card) + ", want: ace of clubs")
	}
}
