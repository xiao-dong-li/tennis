package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input Input
}

func (g *Game) Update() error {
	g.input.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBlock(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
