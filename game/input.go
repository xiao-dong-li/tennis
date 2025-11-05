package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct{}

func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		log.Println("KeyLeft")
	}
}
