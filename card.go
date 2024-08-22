package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Card represents a card in the game
type Card struct {
	Type          int
	Color         color.Color
	Width, Height int
	Position      Coordinate
}

// Constants for card types
const (
	Rock     = 0
	Paper    = 1
	Scissors = 2
)

// Constants for card size
const (
	CardWidth  = 100
	CardHeight = 150
)

// NewCard creates a new card with a specific type and position
func NewCard(cardType int, cardPosition Coordinate) Card {
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
		Type:     cardType,
		Color:    cardColor,
		Width:    CardWidth,
		Height:   CardHeight,
		Position: cardPosition,
	}
}

func (c *Card) Draw(screen *ebiten.Image) {
	cardImage := ebiten.NewImage(c.Width, c.Height)

	// Fill the card image with the card's color
	cardImage.Fill(c.Color)

	// Draw the card image onto the screen at the specified position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.Position.x), float64(c.Position.y))
	screen.DrawImage(cardImage, op)
}

func (c *Card) IsClicked(coord Coordinate) bool {
	return coord.x >= c.Position.x && coord.x <= c.Position.x+c.Width && coord.y >= c.Position.y && coord.y <= c.Position.y+c.Height
}

// String returns a string representation of the card type
func (c *Card) String() string {
	switch c.Type {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return "Unknown"
	}
}
