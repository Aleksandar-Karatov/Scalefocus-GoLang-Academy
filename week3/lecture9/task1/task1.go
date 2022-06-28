package main

import (
	"fmt"
	"log"
	"week3Lecture9Task1/week3/lecture9/task1/carddraw"
	"week3Lecture9Task1/week3/lecture9/task1/cardgame"
)

func main() {
	var deck cardgame.Deck
	deck.New()
	deck.Shuffle()
	allcards, err := carddraw.DrawAllCards(&deck)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cardgame.PrintAllCardsInText(allcards))

	var deck2 cardgame.Deck
	allcards, err = carddraw.DrawAllCards(&deck2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cardgame.PrintAllCardsInText(allcards))
}
