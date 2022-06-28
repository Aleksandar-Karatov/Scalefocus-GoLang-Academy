package carddraw

import (
	"week3Lecture9Task1/week3/lecture9/task1/cardgame"
)

type dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	var cards []cardgame.Card
	temp, err := dealer.Deal()
	for temp != nil {
		cards = append(cards, *temp)
		temp, err = dealer.Deal()
	}
	if err != nil {
		if dealer.Done() {
			return cards, nil
		} else {
			return nil, err
		}
	}
	return cards, nil

}
