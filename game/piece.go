package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	t = true
	f = false
)

// Pieces is the set of all the possible pieces.
var Pieces map[BlockType]*Piece

type Piece struct {
	blockType BlockType
	blocks    [][]bool
}

func init() {
	Pieces = map[BlockType]*Piece{
		BlockTypeI: {
			blockType: BlockTypeI,
			blocks: [][]bool{
				{f, f, f, f},
				{t, t, t, t},
				{f, f, f, f},
				{f, f, f, f},
			},
		},
		BlockTypeJ: {
			blockType: BlockTypeJ,
			blocks: [][]bool{
				{t, f, f},
				{t, t, t},
				{f, f, f},
			},
		},
		BlockTypeL: {
			blockType: BlockTypeL,
			blocks: [][]bool{
				{f, f, t},
				{t, t, t},
				{f, f, f},
			},
		},
		BlockTypeO: {
			blockType: BlockTypeO,
			blocks: [][]bool{
				{t, t},
				{t, t},
			},
		},
		BlockTypeS: {
			blockType: BlockTypeS,
			blocks: [][]bool{
				{f, t, t},
				{t, t, f},
				{f, f, f},
			},
		},
		BlockTypeT: {
			blockType: BlockTypeT,
			blocks: [][]bool{
				{f, t, f},
				{t, t, t},
				{f, f, f},
			},
		},
		BlockTypeZ: {
			blockType: BlockTypeZ,
			blocks: [][]bool{
				{t, t, f},
				{f, t, t},
				{f, f, f},
			},
		},
	}
}

// Draw renders the piece on the given image at the specified position.
func (p *Piece) Draw(r *ebiten.Image, x, y int) {
	for i, row := range p.blocks {
		for j, blocked := range row {
			if blocked {
				drawBlock(r, p.blockType, j*blockWidth+x, i*blockHeight+y)
			}
		}
	}
}

// Rotate rotates a square matrix 90 degrees.
// Direction depends on 'clockwise': true = clockwise, false = counterclockwise.
func (p *Piece) Rotate(clockwise bool) {
	n := len(p.blocks)
	matrix := make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if clockwise {
				matrix[j][n-1-i] = p.blocks[i][j]
			} else {
				matrix[n-1-j][i] = p.blocks[i][j]
			}
		}
	}

	p.blocks = matrix
}
