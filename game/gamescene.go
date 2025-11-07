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
	g.frameCount++

	if g.currentPiece == nil || g.frameCount%300 == 0 {
		g.currentPiece = g.choosePiece()
		g.currentPieceX = 0
		g.currentPieceY = 0
	}

	if i.IsRotateRight() {
		rotate(g.currentPiece.blocks, true)
	} else if i.IsRotateLeft() {
		rotate(g.currentPiece.blocks, false)
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

	currentPiece := g.currentPiece
	x := g.currentPieceX
	y := g.currentPieceY
	for i, row := range currentPiece.blocks {
		for j, blocked := range row {
			if blocked {
				drawBlock(r, currentPiece.blockType, (j+x)*blockWidth, (i+y)*blockHeight)
			}
		}
	}
}

func (g *GameScene) choosePiece() *Piece {
	n := int(BlockTypeMax)
	block := BlockType(rand.IntN(n))
	return Pieces[block]
}
