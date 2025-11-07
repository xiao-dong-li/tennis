package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	initialDelay = 10 // Delay before auto-repeat starts
	repeatRate   = 3  // Repeat every few frames
)

type Input struct{}

func (i *Input) Update() {}

// isKeyRepeated returns true if the key should trigger a repeat action.
// It fires on the first press and then at fixed intervals after an initial delay.
func (i *Input) isKeyRepeated(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d == 1 || (d > initialDelay && (d-initialDelay)%repeatRate == 0)
}

func (i *Input) IsLeft() bool {
	return i.isKeyRepeated(ebiten.KeyLeft)
}

func (i *Input) IsRight() bool {
	return i.isKeyRepeated(ebiten.KeyRight)
}

func (i *Input) IsDown() bool {
	return i.isKeyRepeated(ebiten.KeyDown)
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
