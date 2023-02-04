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

	playerHandSprites []*cardSprite
	dealerHandSprites []*cardSprite

	deck       *cardPile
	playerHand *cardPile
	dealerHand *cardPile
	hold       bool
	wallet     int
}

func NewGame(screenWidth, screenHeight int, cardImages *cardImages, deck *cardPile) *game {
	rand.Seed(time.Now().UnixNano())

	theGame := &game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		cardImages:   cardImages,
		deck:         deck,
		playerHand:   newCardPile(),
		dealerHand:   newCardPile(),
		wallet:       1000,
	}

	// draw 2 for each dealer and player
	theGame.playerHand.add(theGame.deck.draw())
	theGame.dealerHand.add(theGame.deck.draw())
	theGame.playerHand.add(theGame.deck.draw())
	theGame.dealerHand.add(theGame.deck.draw())

	for _, playerCard := range theGame.playerHand.cards {
		cardImage := theGame.cardImages.getFaceUpImage(playerCard.face, playerCard.suit)
		sprite := &cardSprite{
			image: cardImage,
		}
		theGame.playerHandSprites = append(theGame.playerHandSprites, sprite)
	}
	for i, dealerCard := range theGame.dealerHand.cards {
		cardImage := theGame.cardImages.getFaceUpImage(dealerCard.face, dealerCard.suit)
		if i == 0 {
			cardImage = theGame.cardImages.getFaceDownImage()
		}
		sprite := &cardSprite{
			image: cardImage,
		}
		theGame.dealerHandSprites = append(theGame.dealerHandSprites, sprite)
	}

	return theGame
}

func (g *game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if g.playerHand.sum() > 21 || g.hold {
		if inpututil.IsKeyJustPressed(ebiten.KeyN) {
			g.hold = false
			g.playerHand = newCardPile()
			g.dealerHand = newCardPile()
			g.playerHandSprites = make([]*cardSprite, 0)
			g.dealerHandSprites = make([]*cardSprite, 0)

			// draw 2 for each dealer and player
			g.playerHand.add(g.deck.draw())
			g.dealerHand.add(g.deck.draw())
			g.playerHand.add(g.deck.draw())
			g.dealerHand.add(g.deck.draw())

			for _, playerCard := range g.playerHand.cards {
				cardImage := g.cardImages.getFaceUpImage(playerCard.face, playerCard.suit)
				sprite := &cardSprite{
					image: cardImage,
				}
				g.playerHandSprites = append(g.playerHandSprites, sprite)
			}
			for i, dealerCard := range g.dealerHand.cards {
				cardImage := g.cardImages.getFaceUpImage(dealerCard.face, dealerCard.suit)
				if i == 0 {
					cardImage = g.cardImages.getFaceDownImage()
				}
				sprite := &cardSprite{
					image: cardImage,
				}
				g.dealerHandSprites = append(g.dealerHandSprites, sprite)
			}
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.hold = true
		for g.dealerHand.sum() < 17 {
			drawCard := g.deck.draw()
			g.dealerHand.add(drawCard)
			cardImage := g.cardImages.getFaceUpImage(drawCard.face, drawCard.suit)
			sprite := &cardSprite{
				image: cardImage,
			}
			g.dealerHandSprites = append(g.dealerHandSprites, sprite)
		}

		// handle wallet changes here
		dealerSum := g.dealerHand.sum()
		playerSum := g.playerHand.sum()
		if dealerSum > 21 || playerSum > dealerSum {
			g.wallet += 10
		} else if playerSum == dealerSum {

		} else {
			g.wallet -= 10
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		// faceValue := rand.Intn(13) + 2 // 2-14, 2 through King and Ace
		// suit := rand.Intn(4) + 1       // 1-4

		drawCard := g.deck.draw()
		g.playerHand.add(drawCard)
		faceValue := drawCard.face
		suit := drawCard.suit
		fmt.Println("drew card ", drawCard)

		log.Printf("new card: %d of %d", faceValue, suit)
		cardImage := g.cardImages.getFaceUpImage(faceValue, suit)
		sprite := &cardSprite{
			image: cardImage,
		}
		g.playerHandSprites = append(g.playerHandSprites, sprite)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	const cardsPerRow = 9
	const cardWidth = 35
	const cardHeight = 47

	debugY := cardHeight + cardHeight

	playerSum := g.playerHand.sum()
	dealerSum := g.dealerHand.sum()
	winState := "lost"
	if playerSum > 21 {
		winState = "bust"
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Dealer: %d\nPlayer: %d You %s! press 'N' for a new deal", dealerSum, playerSum, winState), 0, debugY)
	} else {
		if g.hold {
			if dealerSum > 21 || playerSum > dealerSum {
				winState = "won"
			} else if playerSum == dealerSum {
				winState = "tied"
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Dealer: %d\nPlayer: %d You %s! press 'N' for a new deal", dealerSum, playerSum, winState), 0, debugY)
		} else {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Dealer: ?\nPlayer: %d press 'D' to Draw, 'H' to Hold", playerSum), 0, debugY)
		}
	}

	debugY += cardHeight
	walletMessage := fmt.Sprintf("wallet: %d", g.wallet)
	ebitenutil.DebugPrintAt(screen, walletMessage, 0, debugY)

	x := 0.0
	y := 0.0

	if g.hold {
		g.dealerHandSprites = make([]*cardSprite, 0)
		for _, dealerCard := range g.dealerHand.cards {
			cardImage := g.cardImages.getFaceUpImage(dealerCard.face, dealerCard.suit)
			sprite := &cardSprite{
				image: cardImage,
			}
			g.dealerHandSprites = append(g.dealerHandSprites, sprite)
		}
	}

	for _, c := range g.dealerHandSprites {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(x, y)
		screen.DrawImage(c.image, opts)
		x += cardWidth
		if x >= cardsPerRow*cardWidth {
			x = 0.0
			y += cardHeight
		}
	}

	x = 0.0
	y = cardHeight
	for _, c := range g.playerHandSprites {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(x, y)
		screen.DrawImage(c.image, opts)
		x += cardWidth
		if x >= cardsPerRow*cardWidth {
			x = 0.0
			y += cardHeight
		}
	}
}

func (g *game) Layout(w, h int) (int, int) { return g.screenWidth, g.screenHeight }
