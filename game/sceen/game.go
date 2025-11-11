package sceen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game"
	"github.com/xiao-dong-li/tennis/game/input"
)

type Game struct {
	sceneManager *SceneManager
	input        input.Input
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
	return game.ScreenWidth, game.ScreenHeight
}
