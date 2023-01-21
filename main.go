package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	suitHeart  = 1
	suitClub   = 2
	suitDiamon = 3
	suitSpade  = 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("blackjack")
	// panic("wait this isn't ready to run yet!")
	// card1 := card{
	//   face: 2,
	//   suit:  suitClub,
	// }
	// fmt.Println(card1)

	deck := generateDeck()
	deck.shuffle()
	fmt.Println(deck.cards)
	fmt.Println(len(deck.cards))

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("blackjack")

	cardsBytesBuffer := bytes.NewBuffer(cardsPNGData)
	cardsImg, err := png.Decode(cardsBytesBuffer)
	if err != nil {
		log.Fatalf("failed to decode cards png data: %s", err.Error())
	}
	cardsImage := ebiten.NewImageFromImage(cardsImg)
	cardImages := newCardImages(cardsImage)

	screenWidth := 320
	screenHeight := 240
	theGame := NewGame(screenWidth, screenHeight, cardImages, deck)
	ebiten.RunGame(theGame)

}

type card struct {
	face int // ace thuoug king: 1-13
	suit int // clubs, diamonds, hearts, and spades: 1-4
}

func generateDeck() *cardPile {
	deck := make([]card, 0, 52)
	for suit := 1; suit <= 4; suit++ {
		for num := 1; num <= 13; num++ {
			aCard := card{
				face: num,
				suit: suit,
			}
			deck = append(deck, aCard)
		}
	}
	return &cardPile{
		cards: deck,
	}
}

func (c card) String() string {
	suitName := "unknown"
	switch c.suit {
	case 1:
		suitName = "hearts"
	case 2:
		suitName = "clubs"
	case 3:
		suitName = "diamonds"
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

	return fmt.Sprintf("%s of %s", faceName, suitName)
}

type cardPile struct {
	cards []card
}

func (c *cardPile) shuffle() {
	rand.Shuffle(len(c.cards), func(i, j int) { c.cards[i], c.cards[j] = c.cards[j], c.cards[i] })
}

func (c *cardPile) draw() card {
	drawCard := c.cards[0]
	c.cards = c.cards[1:]
	return drawCard
}
