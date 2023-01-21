package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// type card struct {
// 	face int
// 	suit int
// }

type cardSprite struct {
	image *ebiten.Image
	// location:
	x int
	y int
}

func (s *cardSprite) in(x, y int) bool {
	return s.image.At(x-s.x, y-s.y).(color.RGBA).A > 0
}
