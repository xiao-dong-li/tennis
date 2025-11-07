package game

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	fieldWidth  = blockWidth * fieldBlockCountX
	fieldHeight = blockHeight * fieldBlockCountY
	topMargin   = (fieldHeight-blockHeight*6)/3 - blockHeight*3 // Top padding area above score, level, and lines
	fontSize    = 8
)

var (
	imageWindows = ebiten.NewImage(ScreenWidth, ScreenHeight)
	fontSource   *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontSource = s

	// Windows: Field
	x, y := fieldWindowPosition()
	drawWindow(x, y, fieldWidth, fieldHeight)

	// Windows: Next
	x, y = nextLabelPosition()
	drawTextWithShadow("NEXT", x, y)
	x, y = nextWindowPosition()
	drawWindow(x, y, blockWidth*5, blockHeight*5)

	// Windows: Score
	x, y = scoreLabelPosition()
	drawTextWindow("SCORE", x, y)

	// Windows: Level
	x, y = levelLabelPosition()
	drawTextWindow("LEVEL", x, y)

	// Windows: Lines
	x, y = linesLabelPosition()
	drawTextWindow("LINES", x, y)
}

func fieldWindowPosition() (x, y int) {
	return 20, 20
}

func nextLabelPosition() (x, y int) {
	x, y = fieldWindowPosition()
	return 2*x + fieldWidth, y
}

func nextWindowPosition() (x, y int) {
	x, y = nextLabelPosition()
	return x, y + blockHeight
}

func scoreLabelPosition() (x, y int) {
	x, y = nextWindowPosition()
	return x, y + blockWidth*5 + topMargin
}

func levelLabelPosition() (x, y int) {
	x, y = scoreLabelPosition()
	return x, y + blockHeight*3 + topMargin
}

func linesLabelPosition() (x, y int) {
	x, y = levelLabelPosition()
	return x, y + blockHeight*3 + topMargin
}

func drawWindow(x, y, width, height int) {
	vector.FillRect(imageWindows, float32(x), float32(y), float32(width), float32(height), color.RGBA{R: 255, G: 255, B: 255, A: 100}, false)
}

func drawTextWithShadow(str string, x, y int) {
	face := &text.GoTextFace{
		Source: fontSource,
		Size:   fontSize,
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	text.Draw(imageWindows, str, face, op)
}

func drawTextWindow(str string, x, y int) {
	drawTextWithShadow(str, x, y)
	y += blockHeight
	fieldX, _ := fieldWindowPosition()
	drawWindow(x, y, ScreenWidth-x-fieldX, blockHeight*2)
}
