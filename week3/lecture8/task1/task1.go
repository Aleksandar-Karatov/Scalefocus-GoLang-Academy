package main

import (
	"fmt"
	"week3Lecture8Task1/week3/lecture8/task1/carddraw"
	"week3Lecture8Task1/week3/lecture8/task1/cardgame"
)

func main() {
	fmt.Println()
	var deck cardgame.Deck
	deck.New()
	deck.Shuffle()

	fmt.Println(cardgame.PrintAllCardsInText(carddraw.DrawAllCards(&deck)))
}
