package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 256
	ScreenHeight = 240
)

type Game struct {
	sceneManager *SceneManager
	input        Input
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = NewSceneManager()
		g.sceneManager.GoTo()
	}

	g.input.Update()
	g.sceneManager.Update(&g.input)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
