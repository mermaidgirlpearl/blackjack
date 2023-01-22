package main

import "testing"

func TestCardsSumAces(t *testing.T) {
	hand := &cardPile{}
	hand.add(card{1, suitHeart})
	hand.add(card{1, suitHeart})
	if hand.sum() != 12 {
		t.Error()
	}
}

func TestCardsSumTens(t *testing.T) {
	hand := &cardPile{}
	hand.add(card{12, suitHeart}) // Queen
	hand.add(card{13, suitHeart}) // King
	if hand.sum() != 20 {
		t.Error()
	}
}
