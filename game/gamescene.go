package game

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	field         [fieldBlockCountY][fieldBlockCountX]BlockType
	currentPiece  *Piece
	currentPieceX int
	currentPieceY int
	nextPiece     *Piece
}

func NewGameScene() *GameScene {
	return &GameScene{}
}

func (g *GameScene) Update(i *Input) {
	if g.currentPiece == nil {
		g.initCurrentPiece(g.choosePiece())
		return
	}

	oldBlocks := make([][]bool, len(g.currentPiece.blocks))
	copy(oldBlocks, g.currentPiece.blocks)

	if i.IsRotateRight() {
		g.currentPiece.Rotate(t)
		if g.collides(g.currentPieceX, g.currentPieceY) {
			g.currentPiece.blocks = oldBlocks
		}
	}
	if i.IsRotateLeft() {
		g.currentPiece.Rotate(f)
		if g.collides(g.currentPieceX, g.currentPieceY) {
			g.currentPiece.blocks = oldBlocks
		}
	}
	if i.IsLeft() && !g.collides(g.currentPieceX-1, g.currentPieceY) {
		g.currentPieceX--
	}
	if i.IsRight() && !g.collides(g.currentPieceX+1, g.currentPieceY) {
		g.currentPieceX++
	}
	if i.IsDown() {
		if !g.collides(g.currentPieceX, g.currentPieceY+1) {
			g.currentPieceY++
		} else {
			g.mergePieceToField()
			g.initCurrentPiece(g.nextPiece)
		}
	}
}

func (g *GameScene) Draw(r *ebiten.Image) {
	// Draw static background and window frame
	drawBackground(r)
	r.DrawImage(imageWindows, nil)

	// Draw field blocks
	fieldX, fieldY := fieldWindowPosition()
	g.drawField(r, fieldX, fieldY)

	// Draw current falling piece
	x := fieldX + g.currentPieceX*blockWidth
	y := fieldY + g.currentPieceY*blockHeight
	g.currentPiece.Draw(r, x, y)

	// Draw next piece preview
	nextX, nextY := nextWindowPosition()
	x = nextX + blockWidth*(5-len(g.nextPiece.blocks))/2
	y = nextY + blockHeight*(5-len(g.nextPiece.blocks))/2
	g.nextPiece.Draw(r, x, y)
}

// initCurrentPiece sets the given piece as the current falling piece.
func (g *GameScene) initCurrentPiece(piece *Piece) {
	g.currentPiece = piece
	g.currentPieceX = (fieldBlockCountX - len(g.currentPiece.blocks)) / 2
	g.currentPieceY = 0

	if g.currentPiece.blockType == BlockTypeI {
		g.currentPieceY--
	}

	g.nextPiece = g.choosePiece()
}

func (g *GameScene) choosePiece() *Piece {
	n := int(BlockTypeMax)
	block := BlockType(rand.IntN(n) + 1)
	return Pieces[block]
}

// collides checks whether the piece collides with walls or placed blocks.
func (g *GameScene) collides(px, py int) bool {
	for i, row := range g.currentPiece.blocks {
		for j, blocked := range row {
			if !blocked {
				continue
			}

			x := px + j
			y := py + i

			// Out of field bounds
			if x < 0 || x >= fieldBlockCountX || y < 0 || y >= fieldBlockCountY {
				return true
			}

			// Collides with placed block
			if g.field[y][x] > BlockTypeNone {
				return true
			}
		}
	}
	return false
}

// mergePieceToField merges the current piece into the field grid.
func (g *GameScene) mergePieceToField() {
	for i, row := range g.currentPiece.blocks {
		for j, blocked := range row {
			if !blocked {
				continue
			}
			x := g.currentPieceX + j
			y := g.currentPieceY + i

			if x >= 0 && x < fieldBlockCountX && y >= 0 && y < fieldBlockCountY {
				g.field[y][x] = g.currentPiece.blockType
			}
		}
	}
}

// drawField renders all placed blocks in the field grid.
func (g *GameScene) drawField(r *ebiten.Image, x, y int) {
	for i, row := range g.field {
		for j, block := range row {
			if block == BlockTypeNone {
				continue
			}
			drawBlock(r, block, j*blockWidth+x, i*blockHeight+y)
		}
	}
}
