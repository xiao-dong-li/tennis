package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/xiao-dong-li/tennis/game"
)

type Input struct{}

func (i *Input) Update() {}

// isKeyRepeated returns true if the key should trigger a repeat action.
// It fires on the first press and then at fixed intervals after an initial delay.
func (i *Input) isKeyRepeated(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d == 1 || (d > game.InputInitialDelay && (d-game.InputInitialDelay)%game.InputRepeatRate == 0)
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

func (i *Input) IsSpace() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
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
	return inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}
