package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Coordinate struct {
	x int
	y int
}

type Game struct {
	PlayerSelection Coordinate
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		g.PlayerSelection.x = x
		g.PlayerSelection.y = y

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for cardType := 0; cardType <= 2; cardType++ {
		c := NewCard(cardType)
		c.Draw(screen, Coordinate{cardType * c.Width, 0})
	}

	textToRender := "Pick a card"
	if g.PlayerSelection.x != 0 {
	}
	text.Draw(screen, textToRender, basicfont.Face7x13, 50, 300, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
