package carddraw

import "week3Lecture8Task1/week3/lecture8/task1/cardgame"

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	var cards []cardgame.Card
	var temp = dealer.Deal()
	for temp != nil {
		cards = append(cards, *temp)
		temp = dealer.Deal()
	}
	return cards
}
