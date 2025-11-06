package game

import "github.com/hajimehoshi/ebiten/v2"

type TitleScene struct {
}

func NewTitleScene() *TitleScene {
	return &TitleScene{}
}

func (t *TitleScene) Update(i *Input) {
}

func (t *TitleScene) Draw(r *ebiten.Image) {
}
