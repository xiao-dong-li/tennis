package game

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	currentPiece  *Piece
	currentPieceX int
	currentPieceY int
	frameCount    int64
}

func NewGameScene() *GameScene {
	return &GameScene{}
}

func (g *GameScene) Update(i *Input) {
	if g.currentPiece == nil {
		g.initCurrentPiece()
	}

	if i.IsRotateRight() {
		g.currentPiece.Rotate(t)
	} else if i.IsRotateLeft() {
		g.currentPiece.Rotate(f)
	} else if i.IsLeft() {
		g.currentPieceX--
	} else if i.IsRight() {
		g.currentPieceX++
	} else if i.IsDown() {
		g.currentPieceY++
	}
}

func (g *GameScene) Draw(r *ebiten.Image) {
	drawBackground(r)
	r.DrawImage(imageWindows, nil)

	fieldX, fieldY := fieldWindowPosition()
	x := fieldX + g.currentPieceX*blockWidth
	y := fieldY + g.currentPieceY*blockHeight
	g.currentPiece.Draw(r, x, y)
}

// initCurrentPiece initializes a new falling piece.
// It selects a random piece, centers it horizontally,
// and slightly shifts up if it's an I-type block.
func (g *GameScene) initCurrentPiece() {
	g.currentPiece = g.choosePiece()
	g.currentPieceX = (fieldBlockCountX - len(g.currentPiece.blocks)) / 2
	g.currentPieceY = 0

	if g.currentPiece.blockType == BlockTypeI {
		g.currentPieceY--
	}
}

func (g *GameScene) choosePiece() *Piece {
	n := int(BlockTypeMax)
	block := BlockType(rand.IntN(n))
	return Pieces[block]
}
