package sceen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game/input"
)

type TitleScene struct {
}

func NewTitleScene() *TitleScene {
	return &TitleScene{}
}

func (t *TitleScene) Update(i *input.Input) {
}

func (t *TitleScene) Draw(r *ebiten.Image) {
}
