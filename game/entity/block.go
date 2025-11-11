package entity

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images/blocks"
	"github.com/xiao-dong-li/tennis/game"
)

var imageBlocks *ebiten.Image

type BlockType int32

const (
	BlockTypeNone BlockType = iota
	BlockTypeI
	BlockTypeJ
	BlockTypeL
	BlockTypeO
	BlockTypeS
	BlockTypeT
	BlockTypeZ
	BlockTypeMax = BlockTypeZ
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(blocks.Blocks_png))
	if err != nil {
		log.Fatal(err)
	}
	imageBlocks = ebiten.NewImageFromImage(img)
}

// drawBlock draws a single block of the given type onto the target image.
// (x, y) specifies the top-left pixel position.
func drawBlock(dst *ebiten.Image, block BlockType, x, y int) {
	if block == BlockTypeNone {
		return
	}

	srcX := (int(block) - 1) * game.BlockWidth
	srcRect := image.Rect(srcX, 0, srcX+game.BlockWidth, game.BlockHeight)
	srcImg := imageBlocks.SubImage(srcRect).(*ebiten.Image)

	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	colorm.DrawImage(dst, srcImg, colorm.ColorM{}, op)
}
