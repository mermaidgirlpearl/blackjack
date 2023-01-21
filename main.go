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
  // panic("wait this isn't ready to run yet!")
  // card1 := card{
  //   face: 2,
  //   suit:  suitClub,
  // }
  // fmt.Println(card1)

  deck := generateDeck()
  fmt.Println(deck)
  fmt.Println(len(deck))
}

type card struct {
  face int // ace thuoug king: 1-13
  suit int // clubs, diamonds, hearts, and spades: 1-4
}

func generateDeck() []card   {
  deck := make([]card, 0, 52)
  for suit := 1; suit <=4; suit++ {
    for num:=1; num<=13; num++{
      aCard := card{
        face: num,
        suit: suit,
      }
      deck=append(deck, aCard)
    }
  }
  return deck
}

func (c card) String() string {
  suitName := "unknown"
  switch c.suit {
  case 1:
    suitName = "clubs"
  case 2:
    suitName = "diamonds"
  case 3:
    suitName = "hearts"
  case 4:
    suitName = "spades"
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
