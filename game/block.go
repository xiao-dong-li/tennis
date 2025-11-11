package game

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images/blocks"
)

const (
	blockWidth       = 10
	blockHeight      = 10
	fieldBlockCountX = 10
	fieldBlockCountY = 20
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
// x, y specify the drawing position in pixels (top-left corner).
func drawBlock(r *ebiten.Image, block BlockType, x, y int) {
	x0 := (int(block) - 1) * blockWidth
	x1 := x0 + blockWidth
	img := imageBlocks.SubImage(image.Rect(x0, 0, x1, blockHeight)).(*ebiten.Image)

	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	clr := colorm.ColorM{}

	colorm.DrawImage(r, img, clr, op)
}
