package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cardWidth  = 35
	cardHeight = 47
)

type cardImages struct {
	allCardsImage *ebiten.Image
}

func newCardImages(allCardsImage *ebiten.Image) *cardImages {
	return &cardImages{
		allCardsImage: allCardsImage,
	}
}

func (c *cardImages) getFaceUpImage(faceValue, suitPictureValue int) *ebiten.Image {
	// face := 2
	// suit := 2 // hearts

	// mapping the card structs face and suit values with that of the png file
	facePictureValue := faceValue
	if faceValue == 1 {
		facePictureValue = 14
	}

	if facePictureValue < 2 || facePictureValue > 14 {
		panic(fmt.Sprintf("invalid card face value: %d", facePictureValue)) // TODO: do this with proper error handling
	}
	if suitPictureValue < 1 || suitPictureValue > 4 { // 1-4: clubs, diamonds, hearts, spades
		panic(fmt.Sprintf("invalid card suit: %d", suitPictureValue)) // TODO: do this with proper error handling
	}

	xStart := (facePictureValue - 1) * cardWidth
	yStart := (suitPictureValue * cardHeight) - cardHeight

	sub := c.allCardsImage.SubImage(image.Rect(xStart, yStart, xStart+cardWidth, yStart+cardHeight))
	return sub.(*ebiten.Image)
}

func (c *cardImages) getFaceDownImage() *ebiten.Image {
	xStart := 0
	yStart := 0
	sub := c.allCardsImage.SubImage(image.Rect(xStart, yStart, xStart+cardWidth, yStart+cardHeight))
	return sub.(*ebiten.Image)
}
