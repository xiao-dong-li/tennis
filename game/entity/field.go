package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game"
)

type Field struct {
	Blocks [game.FieldBlockCountY][game.FieldBlockCountX]BlockType
}

// AddPiece merges the piece into the field grid.
func (f *Field) AddPiece(p *Piece, px, py int) {
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
