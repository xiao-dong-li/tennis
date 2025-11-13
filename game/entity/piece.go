package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game"
)

// Pieces defines all available Tetris pieces.
var Pieces map[BlockType]*Piece

// Piece represents a Tetris block piece.
type Piece struct {
	BlockType BlockType
	Blocks    [][]bool
}

func init() {
	const (
		t = true
		f = false
	)
	Pieces = map[BlockType]*Piece{
		BlockTypeI: newPiece(
			BlockTypeI, [][]bool{
				{f, f, f, f},
				{t, t, t, t},
				{f, f, f, f},
				{f, f, f, f},
			},
		),
		BlockTypeJ: newPiece(
			BlockTypeJ, [][]bool{
				{t, f, f},
				{t, t, t},
				{f, f, f},
			},
		),
		BlockTypeL: newPiece(
			BlockTypeL, [][]bool{
				{f, f, t},
				{t, t, t},
				{f, f, f},
			},
		),
		BlockTypeO: newPiece(
			BlockTypeO, [][]bool{
				{t, t},
				{t, t},
			},
		),
		BlockTypeS: newPiece(
			BlockTypeS, [][]bool{
				{f, t, t},
				{t, t, f},
				{f, f, f},
			},
		),
		BlockTypeT: newPiece(
			BlockTypeT, [][]bool{
				{f, t, f},
				{t, t, t},
				{f, f, f},
			},
		),
		BlockTypeZ: newPiece(
			BlockTypeZ, [][]bool{
				{t, t, f},
				{f, t, t},
				{f, f, f},
			},
		),
	}
}

func newPiece(blockType BlockType, blocks [][]bool) *Piece {
	return &Piece{
		BlockType: blockType,
		Blocks:    blocks,
	}
}

// Clone returns a deep copy of the current Piece.
func (p *Piece) Clone() *Piece {
	return &Piece{
		BlockType: p.BlockType,
		Blocks:    game.CloneMatrix(p.Blocks),
	}
}

// Rotate rotates the piece 90 degrees.
// Direction depends on 'clockwise': true = clockwise, false = counterclockwise.
func (p *Piece) Rotate(clockwise bool) {
	n := len(p.Blocks)
	matrix := make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if clockwise {
				matrix[j][n-1-i] = p.Blocks[i][j]
			} else {
				matrix[n-1-j][i] = p.Blocks[i][j]
			}
		}
	}

	p.Blocks = matrix
}

// Draw renders the piece on the target image at the given position (x, y).
func (p *Piece) Draw(r *ebiten.Image, x, y int) {
	for i, row := range p.Blocks {
		for j, blocked := range row {
			if blocked {
				drawBlock(r, p.BlockType, j*game.BlockWidth+x, i*game.BlockHeight+y)
			}
		}
	}
}
