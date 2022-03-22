package main

import (
	"fmt"
	"math/rand"
	"time"
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

type Deck struct {
	data []Card
}

//shuffle
//deal
//new

func (r *Deck) New() {
	for val := two; val <= ace; val++ {
		for suit := clubs; suit <= spades; suit++ {
			r.data = append(r.data, Card{val, suit})
		}
	}
}
func (r *Deck) Shuffle() {
	rand.Seed(time.Now().UnixMilli())
	var temp []Card
	for _, v := range r.data {
		temp = append(temp, v)
	}

	for i := 0; i < len(r.data); i++ {
		ind := rand.Intn(len(temp))
		r.data[i] = temp[ind]
		temp = RemoveAtIndex(temp, ind)
	}
}
func (r *Deck) Deal() Card {
	result := r.data[len(r.data)-1]
	r.data = r.data[0 : len(r.data)-1]
	return result
}
func RemoveAtIndex(cards []Card, index int) []Card {
	return append(cards[:index], cards[index+1:]...)
}
func main() {
	var d Deck
	d.New()

	d.Shuffle()

	fmt.Println(cardInText(d.Deal()))
	fmt.Println(cardInText(d.Deal()))
	fmt.Println(cardInText(d.Deal()))
	fmt.Println(cardInText(d.Deal()))
	fmt.Println(cardInText(d.Deal()))
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
