package sceen

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xiao-dong-li/tennis/game"
	"github.com/xiao-dong-li/tennis/game/entity"
	"github.com/xiao-dong-li/tennis/game/input"
	"github.com/xiao-dong-li/tennis/game/render"
)

type GameScene struct {
	field         *entity.Field
	currentPiece  *entity.Piece
	currentPieceX int
	currentPieceY int
	nextPiece     *entity.Piece
}

func NewGameScene() *GameScene {
	return &GameScene{
		field: &entity.Field{},
	}
}

func (g *GameScene) Update(i *input.Input) {
	if g.currentPiece == nil {
		g.spawnPiece(g.choosePiece())
		return
	}

	if i.IsRotateRight() || i.IsRotateLeft() {
		oldBlocks := game.CloneMatrix(g.currentPiece.Blocks)
		clockwise := i.IsRotateRight()
		g.currentPiece.Rotate(clockwise)
		if g.collides(g.currentPieceX, g.currentPieceY) {
			g.currentPiece.Blocks = oldBlocks
		}
	}

	var moveX int
	if i.IsLeft() {
		moveX = -1
	} else if i.IsRight() {
		moveX = 1
	}
	if moveX != 0 && !g.collides(g.currentPieceX+moveX, g.currentPieceY) {
		g.currentPieceX += moveX
	}
	if i.IsDown() {
		if !g.collides(g.currentPieceX, g.currentPieceY+1) {
			g.currentPieceY++
		} else {
			g.field.AddPiece(g.currentPiece, g.currentPieceX, g.currentPieceY)
			g.spawnPiece(g.nextPiece)
		}
	}
}

func (g *GameScene) Draw(r *ebiten.Image) {
	// Draw static background and window frame
	render.DrawSceneBackground(r)

	// Draw field blocks
	fieldX, fieldY := render.FieldWindowPosition()
	g.field.Draw(r, fieldX, fieldY)

	// Draw current falling piece
	x := fieldX + g.currentPieceX*game.BlockWidth
	y := fieldY + g.currentPieceY*game.BlockHeight
	g.currentPiece.Draw(r, x, y)

	// Draw next piece preview
	nextX, nextY := render.NextWindowPosition()
	x = nextX + game.BlockWidth*(5-len(g.nextPiece.Blocks))/2
	y = nextY + game.BlockHeight*(5-len(g.nextPiece.Blocks))/2
	g.nextPiece.Draw(r, x, y)
}

func (g *GameScene) choosePiece() *entity.Piece {
	n := int(entity.BlockTypeMax)
	block := entity.BlockType(rand.IntN(n) + 1)
	return entity.Pieces[block].Clone()
}

// spawnPiece initializes a new falling piece.
func (g *GameScene) spawnPiece(piece *entity.Piece) {
	g.currentPiece = piece
	g.currentPieceX = (game.FieldBlockCountX - len(piece.Blocks)) / 2
	g.currentPieceY = 0

	if g.currentPiece.BlockType == entity.BlockTypeI {
		g.currentPieceY--
	}

	g.nextPiece = g.choosePiece()
}

// collides checks whether the piece collides with walls or placed blocks.
func (g *GameScene) collides(px, py int) bool {
	for i, row := range g.currentPiece.Blocks {
		for j, blocked := range row {
			if !blocked {
				continue
			}
			x, y := px+j, py+i
			// Out of field bounds
			if x < 0 || x >= game.FieldBlockCountX || y < 0 || y >= game.FieldBlockCountY {
				return true
			}
			// Collides with placed block
			if g.field.Blocks[y][x] > entity.BlockTypeNone {
				return true
			}
		}
	}
	return false
}
