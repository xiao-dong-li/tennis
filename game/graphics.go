package game

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
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
	imageBackground *ebiten.Image
	imageWindows    = ebiten.NewImage(ScreenWidth, ScreenHeight)
	fontSource      *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontSource = s

	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	imageBackground = ebiten.NewImageFromImage(img)

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

func drawBackground(r *ebiten.Image) {
	bgWidth := imageBackground.Bounds().Dx()
	bgHeight := imageBackground.Bounds().Dy()

	scaleX := ScreenWidth / float64(bgWidth)
	scaleY := ScreenHeight / float64(bgHeight)

	op := &colorm.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	clr := colorm.ColorM{}
	clr.Translate(0.22, 0.22, 0.22, 0)
	clr.ChangeHSV(0, 0.3, 1)

	colorm.DrawImage(r, imageBackground, clr, op)
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

// drawWindow draws a semi-transparent rectangular window at the given position.
func drawWindow(x, y, width, height int) {
	vector.FillRect(imageWindows, float32(x), float32(y), float32(width), float32(height), color.RGBA{R: 0, G: 0, B: 0, A: 192}, false)
}

// drawTextWithShadow draws a string with a simple drop shadow effect.
func drawTextWithShadow(str string, x, y int) {
	face := &text.GoTextFace{
		Source: fontSource,
		Size:   fontSize,
	}

	// shadow layer
	shadowOp := &text.DrawOptions{}
	shadowOp.GeoM.Translate(float64(x)+1, float64(y)+1)
	shadowOp.ColorScale.ScaleWithColor(color.RGBA{R: 0, G: 0, B: 0, A: 128})
	text.Draw(imageWindows, str, face, shadowOp)

	// text layer
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color.RGBA{R: 64, G: 64, B: 255, A: 255})
	text.Draw(imageWindows, str, face, op)
}

func drawTextWindow(str string, x, y int) {
	drawTextWithShadow(str, x, y)
	y += blockHeight
	fieldX, _ := fieldWindowPosition()
	drawWindow(x, y, ScreenWidth-x-fieldX, blockHeight*2)
}
