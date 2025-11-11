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
	fallCounter   int
	lines         int
}

func NewGameScene() *GameScene {
	return &GameScene{
		field: &entity.Field{},
	}
}

func (g *GameScene) Update(i *input.Input) {
	if g.currentPiece == nil {
		g.SpawnPiece(g.ChoosePiece())
		return
	}

	if i.IsRotateRight() || i.IsRotateLeft() {
		oldBlocks := game.CloneMatrix(g.currentPiece.Blocks)
		clockwise := i.IsRotateRight()
		g.currentPiece.Rotate(clockwise)
		if g.Collides(g.currentPieceX, g.currentPieceY) {
			g.currentPiece.Blocks = oldBlocks
		}
	}

	var moveX int
	if i.IsLeft() {
		moveX = -1
	} else if i.IsRight() {
		moveX = 1
	}
	if moveX != 0 && !g.Collides(g.currentPieceX+moveX, g.currentPieceY) {
		g.currentPieceX += moveX
	}

	g.fallCounter++
	gravityDrop := g.fallCounter >= max(game.DropIntervalBase-g.Level()*3, game.MinFallInterval)

	if i.IsDown() || gravityDrop {
		g.fallCounter = 0
		if !g.Collides(g.currentPieceX, g.currentPieceY+1) {
			g.currentPieceY++
		} else {
			g.field.AddPiece(g.currentPiece, g.currentPieceX, g.currentPieceY)
			g.lines += g.field.LineClear()
			g.SpawnPiece(g.nextPiece)
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

func (g *GameScene) ChoosePiece() *entity.Piece {
	n := int(entity.BlockTypeMax)
	block := entity.BlockType(rand.IntN(n) + 1)
	return entity.Pieces[block].Clone()
}

func (g *GameScene) Level() int {
	return g.lines / game.LinesPerLevel
}

// SpawnPiece initializes a new falling piece.
func (g *GameScene) SpawnPiece(piece *entity.Piece) {
	g.currentPiece = piece
	g.currentPieceX = (game.FieldBlockCountX - len(piece.Blocks)) / 2
	g.currentPieceY = 0

	if g.currentPiece.BlockType == entity.BlockTypeI {
		g.currentPieceY--
	}

	g.nextPiece = g.ChoosePiece()
}

// Collides checks whether the piece collides with walls or placed blocks.
func (g *GameScene) Collides(px, py int) bool {
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
