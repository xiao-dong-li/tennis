package game

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images/blocks"
)

var imageBlocks *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(blocks.Blocks_png))
	if err != nil {
		log.Fatal(err)
	}
	imageBlocks = ebiten.NewImageFromImage(img)
}

type BlockType int32

const (
	BlockTypeI BlockType = iota
	BlockTypeJ
	BlockTypeL
	BlockTypeO
	BlockTypeS
	BlockTypeT
	BlockTypeZ
	BlockTypeMax
)

const (
	blockWidth  = 10
	blockHeight = 10
)

// drawBlock draws a single block of the given type onto the target image.
// x, y specify the drawing position in pixels (top-left corner).
func drawBlock(r *ebiten.Image, block BlockType, x, y int) {
	x0 := int(block) * blockWidth
	x1 := x0 + blockWidth
	img := imageBlocks.SubImage(image.Rect(x0, 0, x1, blockHeight)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	r.DrawImage(img, op)
}
