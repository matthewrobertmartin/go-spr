package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Card represents a card in the game
type Card struct {
	Type          int
	Color         color.Color
	Width, Height int
}

// Constants for card types
const (
	Rock     = 0
	Paper    = 1
	Scissors = 2
)

// NewCard creates a new card with a specific type and position
func NewCard(cardType int) Card {
	var cardColor color.Color
	switch cardType {
	case Rock:
		cardColor = color.Gray{Y: 128} // Grey color for Rock
	case Paper:
		cardColor = color.White // White color for Paper
	case Scissors:
		cardColor = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red color
	}
	return Card{
		Type:   cardType,
		Color:  cardColor,
		Width:  100,
		Height: 150,
	}
}

func (c *Card) Draw(screen *ebiten.Image, coord Coordinate) {
	cardImage := ebiten.NewImage(c.Width, c.Height)

	// Fill the card image with the card's color
	cardImage.Fill(c.Color)

	// Draw the card image onto the screen at the specified position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(coord.x), float64(coord.y))
	screen.DrawImage(cardImage, op)
}
