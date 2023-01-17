package main

import (
  "fmt"
)

const (
  suitClub = 1
  suitDiamon = 2
  suitHeart = 3
  suitSpade = 4
)

func main() {
  fmt.Println("blackjack")
  card1 := card{
    face: 2,
    suit:  suitClub,
  }
  fmt.Println(card1)
}

type card struct {
  face int // ace thuoug king: 1-13
  suit int // clubs, diamonds, hearts, and spades: 1-4
}

func (c card) String() string {
  suitName := "clubs" // 1
  if c.suit == 2 {
    suitName = "diamonds" // 2
  } else if c.suit == 3 {
    suitName = "hearts" // 3
  } else {
    suitName = "spades" // 4
  }

  faceName := fmt.Sprintf("%d", c.face)
  if c.face == 1 {
    faceName = "Ace"
  } else if c.face == 11 {
    faceName = "Jack"
  } else if c.face == 12 {
    faceName = "Queen"
  } else if c.face == 13 {
    faceName = "King"
  }

  return fmt.Sprintf("%s of %s",faceName, suitName)
}

type cardpile struct{
  cards []card
}
