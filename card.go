package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Card represents a card in the game
type Card struct {
	Type  int
	Color color.Color
	X, Y  float64
}

// Constants for card types
const (
	Rock = iota
	Paper
	Scissors
)

// NewCard creates a new card with a specific type and position
func NewCard(cardType int, x, y float64) Card {
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
		Type:  cardType,
		Color: cardColor,
		X:     x,
		Y:     y,
	}
}

// Draw draws the card on the screen
func (c *Card) Draw(screen *ebiten.Image) {
	// Define the size of the card
	cardWidth, cardHeight := 100, 150

	// Create an image for the card
	cardImage := ebiten.NewImage(cardWidth, cardHeight)

	// Fill the card image with the card's color
	cardImage.Fill(c.Color)

	// Draw the card image onto the screen at the specified position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.X, c.Y)
	screen.DrawImage(cardImage, op)
}
