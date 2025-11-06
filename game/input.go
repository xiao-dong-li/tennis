package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func (i *Input) Update() {}

func (i *Input) IsLeft() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func (i *Input) IsRight() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func (i *Input) IsDown() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

func (i *Input) IsHardDrop() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func (i *Input) IsRotateRight() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyX)
}

func (i *Input) IsRotateLeft() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyZ)
}

func (i *Input) IsPause() bool {
	return ebiten.IsKeyPressed(ebiten.KeyP)
}
