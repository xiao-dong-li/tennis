package entity

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game"
)

type Field struct {
	Blocks [game.FieldBlockCountY][game.FieldBlockCountX]BlockType
}

// Merge the piece into the field grid.
func (f *Field) Merge(p *Piece, px, py int) {
	for i, row := range p.Blocks {
		for j, blocked := range row {
			if !blocked {
				continue
			}
			x, y := px+j, py+i
			if x >= 0 && x < game.FieldBlockCountX && y >= 0 && y < game.FieldBlockCountY {
				f.Blocks[y][x] = p.BlockType
			}
		}
	}
}

// LineClear removes all full lines from the field.
func (f *Field) LineClear() int {
	var fullLines int
	for y, row := range f.Blocks {
		if !slices.Contains(row[:], BlockTypeNone) {
			fullLines++
			// Shift all rows above down by one
			copy(f.Blocks[1:y+1], f.Blocks[0:y])
			// Clear the top row
			for x := 0; x < game.FieldBlockCountX; x++ {
				f.Blocks[0][x] = BlockTypeNone
			}
		}
	}
	return fullLines
}

// Draw renders all placed blocks in the field grid.
func (f *Field) Draw(r *ebiten.Image, x, y int) {
	for i, row := range f.Blocks {
		for j, block := range row {
			if block > BlockTypeNone {
				drawBlock(r, block, j*game.BlockWidth+x, i*game.BlockHeight+y)
			}
		}
	}
}
