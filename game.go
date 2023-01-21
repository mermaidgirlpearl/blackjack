package main

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed cards.png
var cardsPNGData []byte

type game struct {
	screenWidth  int
	screenHeight int
	cardImages   *cardImages

	cardSprites []*cardSprite

	deck *cardPile
}

func NewGame(screenWidth, screenHeight int, cardImages *cardImages, deck *cardPile) *game {
	rand.Seed(time.Now().UnixNano())

	return &game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		cardImages:   cardImages,
		deck:         deck,
	}
}

func (g *game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		// faceValue := rand.Intn(13) + 2 // 2-14, 2 through King and Ace
		// suit := rand.Intn(4) + 1       // 1-4

		drawCard := g.deck.draw()
		faceValue := drawCard.face
		suit := drawCard.suit
		fmt.Println("drew card ", drawCard)

		log.Printf("new card: %d of %d", faceValue, suit)
		cardImage := g.cardImages.getFaceUpImage(faceValue, suit)
		sprite := &cardSprite{
			image: cardImage,
		}
		g.cardSprites = append(g.cardSprites, sprite)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	const cardsPerRow = 9
	const cardWidth = 35
	const cardHeight = 47

	debugY := 0
	if len(g.cardSprites) > 0 {
		debugY = (((len(g.cardSprites) - 1) / 9) + 1) * cardHeight
	}
	ebitenutil.DebugPrintAt(screen, "press 'D' to Draw", 0, debugY)
	// opts.GeoM.Skew(0, .25)
	// opts.GeoM.Scale(1, .80)
	// opts.GeoM.

	x := 0.0
	y := 0.0
	// length := len(g.cardSprites)
	for i, c := range g.cardSprites {
		_ = i
		// offset :=
		// alpha := length - i
		// opts.ColorM.Translate(0, 0, 0, -(float64(alpha) / 100))
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(x, y)
		screen.DrawImage(c.image, opts)
		// log.Printf("i: %v, x: %v", i, x)
		x += cardWidth
		if x >= cardsPerRow*cardWidth {
			x = 0.0
			y += cardHeight
		}
	}

	// screen.DrawImage(g.cardImages.getFaceUpImage(2, 2), opts)
}

func (g *game) Layout(w, h int) (int, int) { return g.screenWidth, g.screenHeight }
