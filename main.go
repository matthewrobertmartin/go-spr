package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Coordinate struct {
	x int `default:"-1"`
	y int `default:"-1"`
}

type Game struct {
	PlayerSelection Coordinate
	PlayerChoice    Card
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		g.PlayerSelection.x = x
		g.PlayerSelection.y = y

	} else {
		g.PlayerSelection.x = 0
		g.PlayerSelection.y = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//var textToRender string
	for cardType := 0; cardType <= 2; cardType++ {
		c := NewCard(cardType, Coordinate{cardType*CardWidth + 640/3 - 50, 480 - CardHeight})
		c.Draw(screen)
		if c.IsClicked(g.PlayerSelection) {
			g.PlayerChoice = c
			//textToRender = fmt.Sprintf("you have clicked %s", c.String())
			//text.Draw(screen, textToRender, basicfont.Face7x13, 50, 300, color.White)
			botChoice := rand.Intn(2)
			player := fmt.Sprintf("player clicked %d       ", c.Type)
			bot := fmt.Sprintf("bot clicked %d", botChoice)
			text.Draw(screen, player+bot, basicfont.Face7x13, 50, 200, color.White)
			if (g.PlayerChoice.Type-botChoice+3)%3 == 1 {
				text.Draw(screen, "Player wins", basicfont.Face7x13, 50, 300, color.White)
			} else {
				text.Draw(screen, "Bot wins", basicfont.Face7x13, 50, 300, color.White)
			}
		}
	}
	text.Draw(screen, fmt.Sprintf("%d, %d", g.PlayerSelection.x, g.PlayerSelection.y), basicfont.Face7x13, 50, 350, color.White)
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
