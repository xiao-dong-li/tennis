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
	score         int
	lines         int
	gameOver      bool
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

	if g.gameOver {
		return
	}

	// Handle rotation
	g.handleRotation(i)

	// Handle horizontal movement
	g.handleMovement(i)

	// Handle falling
	g.handleFalling(i)
}

func (g *GameScene) Draw(r *ebiten.Image) {
	// Draw static background and window frame
	render.DrawSceneBackground(r)

	// Draw stats panel
	render.DrawStatsPanel(r, g.score, g.Level(), g.lines)

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

	if g.gameOver {
		render.DrawGameOver(r)
	}
}

// handleRotation processes clockwise and counterclockwise rotation input.
func (g *GameScene) handleRotation(i *input.Input) {
	if !i.IsRotateRight() && !i.IsRotateLeft() {
		return
	}

	oldBlocks := game.CloneMatrix(g.currentPiece.Blocks)
	clockwise := i.IsRotateRight()
	g.currentPiece.Rotate(clockwise)

	if g.Collides(g.currentPieceX, g.currentPieceY) {
		g.currentPiece.Blocks = oldBlocks
	}
}

// handleMovement processes left and right movement input.
func (g *GameScene) handleMovement(i *input.Input) {
	var dx int
	if i.IsLeft() {
		dx = -1
	} else if i.IsRight() {
		dx = 1
	}

	if dx != 0 && !g.Collides(g.currentPieceX+dx, g.currentPieceY) {
		g.currentPieceX += dx
	}
}

// handleFalling processes gravity, soft drop, and hard drop.
func (g *GameScene) handleFalling(i *input.Input) {
	g.fallCounter++

	// Hard drop
	if i.IsHardDrop() {
		g.fallCounter = 0
		for !g.Collides(g.currentPieceX, g.currentPieceY+1) {
			g.currentPieceY++
		}
		g.LockPiece()
		return
	}

	// Gravity or soft drop
	gravityDrop := g.fallCounter >= max(game.BaseDropInterval-g.Level()*3, game.MinDropInterval)
	if i.IsDown() || gravityDrop {
		g.fallCounter = 0
		if !g.Collides(g.currentPieceX, g.currentPieceY+1) {
			g.currentPieceY++
		} else {
			g.LockPiece()
		}
	}
}

// Level calculates the current level based on cleared lines.
func (g *GameScene) Level() int {
	return g.lines / game.LinesPerLevel
}

// ChoosePiece returns a random new piece.
func (g *GameScene) ChoosePiece() *entity.Piece {
	n := int(entity.BlockTypeMax)
	block := entity.BlockType(rand.IntN(n) + 1)
	return entity.Pieces[block].Clone()
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

	if g.Collides(g.currentPieceX, g.currentPieceY) {
		g.gameOver = true
	}
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

// LockPiece merges the current piece into the field, clears lines, updates score,
// and spawns the next piece.
func (g *GameScene) LockPiece() {
	g.field.Merge(g.currentPiece, g.currentPieceX, g.currentPieceY)
	cleared := g.field.LineClear()

	// Update score first (based on old level)
	g.UpdateScore(cleared)

	// Then update total cleared lines
	g.lines += cleared
	g.SpawnPiece(g.nextPiece)
}

// UpdateScore updates the score based on the number of cleared lines and current level.
func (g *GameScene) UpdateScore(lines int) {
	level := g.Level() + 1
	switch lines {
	case 1:
		g.score += 40 * level
	case 2:
		g.score += 100 * level
	case 3:
		g.score += 300 * level
	case 4:
		g.score += 1200 * level
	}
}
