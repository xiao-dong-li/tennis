package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game"
	"github.com/xiao-dong-li/tennis/game/sceen"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth*2, game.ScreenHeight*2)
	ebiten.SetWindowTitle("Tetris")
	if err := ebiten.RunGame(&sceen.Game{}); err != nil {
		log.Fatal(err)
	}
}
