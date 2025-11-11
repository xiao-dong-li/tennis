package render

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
	"github.com/xiao-dong-li/tennis/game"
)

const fontSize = 8

var (
	imageBackground *ebiten.Image
	imageWindows    = ebiten.NewImage(game.ScreenWidth, game.ScreenHeight)
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
	x, y := FieldWindowPosition()
	drawWindow(x, y, game.FieldWidth, game.FieldHeight)

	// Windows: Next
	x, y = NextLabelPosition()
	drawTextWithShadow("NEXT", x, y)
	x, y = NextWindowPosition()
	drawWindow(x, y, game.BlockWidth*5, game.BlockHeight*5)

	// Windows: Score
	x, y = ScoreLabelPosition()
	drawTextWindow("SCORE", x, y)

	// Windows: Level
	x, y = LevelLabelPosition()
	drawTextWindow("LEVEL", x, y)

	// Windows: Lines
	x, y = LinesLabelPosition()
	drawTextWindow("LINES", x, y)
}

// DrawSceneBackground draws the overall scene background including
// the background image and static window frames.
func DrawSceneBackground(r *ebiten.Image) {
	// draw background image
	drawBackground(r)

	// draw window overlays
	r.DrawImage(imageWindows, nil)
}

// drawBackground draws the background image.
func drawBackground(r *ebiten.Image) {
	bgWidth := imageBackground.Bounds().Dx()
	bgHeight := imageBackground.Bounds().Dy()

	scaleX := game.ScreenWidth / float64(bgWidth)
	scaleY := game.ScreenHeight / float64(bgHeight)

	op := &colorm.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	clr := colorm.ColorM{}
	clr.Translate(0.22, 0.22, 0.22, 0)
	clr.ChangeHSV(0, 0.3, 1)

	colorm.DrawImage(r, imageBackground, clr, op)
}

func FieldWindowPosition() (x, y int) {
	return 20, 20
}

func NextLabelPosition() (x, y int) {
	x, y = FieldWindowPosition()
	return 2*x + game.FieldWidth, y
}

func NextWindowPosition() (x, y int) {
	x, y = NextLabelPosition()
	return x, y + game.BlockHeight
}

func ScoreLabelPosition() (x, y int) {
	x, y = NextWindowPosition()
	return x, y + game.BlockWidth*5 + game.TopMargin
}

func LevelLabelPosition() (x, y int) {
	x, y = ScoreLabelPosition()
	return x, y + game.BlockHeight*3 + game.TopMargin
}

func LinesLabelPosition() (x, y int) {
	x, y = LevelLabelPosition()
	return x, y + game.BlockHeight*3 + game.TopMargin
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
	fieldX, _ := FieldWindowPosition()
	drawWindow(x, y+game.BlockHeight, game.ScreenWidth-x-fieldX, game.BlockHeight*2)
}
