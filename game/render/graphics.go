package render

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/xiao-dong-li/tennis/game"
)

var (
	imageBackground *ebiten.Image
	imageWindows    = ebiten.NewImage(game.ScreenWidth, game.ScreenHeight)
	imageGameOver   = ebiten.NewImage(game.ScreenWidth, game.ScreenHeight)
	fontSource      *text.GoTextFaceSource
	labelColor      = color.RGBA{R: 64, G: 64, B: 255, A: 255}
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
	drawWindow(imageWindows, x, y, game.FieldWidth, game.FieldHeight)

	// Windows: Next
	x, y = NextLabelPosition()
	DrawTextWithShadow(imageWindows, "NEXT", game.FontSize, x, y, labelColor, text.AlignStart, text.AlignStart)
	x, y = NextWindowPosition()
	drawWindow(imageWindows, x, y, game.BlockWidth*5, game.BlockHeight*5)

	// Windows: Score
	x, y = ScoreLabelPosition()
	drawTextWindow(imageWindows, "SCORE", x, y)

	// Windows: Level
	x, y = LevelLabelPosition()
	drawTextWindow(imageWindows, "LEVEL", x, y)

	// Windows: Lines
	x, y = LinesLabelPosition()
	drawTextWindow(imageWindows, "LINES", x, y)

	// Game Over
	imageGameOver.Fill(color.RGBA{A: 128})
	DrawTextWithShadow(
		imageGameOver,
		"GAME OVER\n\nPRESS SPACE",
		game.FontSize,
		game.ScreenWidth/2,
		game.ScreenHeight/2,
		color.White,
		text.AlignCenter,
		text.AlignCenter,
	)
}

// DrawSceneBackground draws the overall scene background including
// the background image and static window frames.
func DrawSceneBackground(r *ebiten.Image) {
	// Draw background image
	drawBackground(r)

	// Draw window overlays
	r.DrawImage(imageWindows, nil)
}

// DrawStatsPanel renders the score, level, and line count on the right-side UI panel.
func DrawStatsPanel(r *ebiten.Image, score, level, lines int) {
	fieldX, _ := FieldWindowPosition()
	x := game.ScreenWidth - fieldX - 5

	// Draw score
	_, y := ScoreLabelPosition()
	DrawTextWithShadow(
		r,
		strconv.Itoa(score),
		game.FontSize,
		x,
		y+game.BlockHeight*2,
		color.White,
		text.AlignEnd,
		text.AlignCenter,
	)

	// Draw level
	_, y = LevelLabelPosition()
	DrawTextWithShadow(
		r,
		strconv.Itoa(level),
		game.FontSize,
		x,
		y+game.BlockHeight*2,
		color.White,
		text.AlignEnd,
		text.AlignCenter,
	)

	// Draw lines
	_, y = LinesLabelPosition()
	DrawTextWithShadow(
		r,
		strconv.Itoa(lines),
		game.FontSize,
		x,
		y+game.BlockHeight*2,
		color.White,
		text.AlignEnd,
		text.AlignCenter,
	)
}

func DrawGameOver(r *ebiten.Image) {
	r.DrawImage(imageGameOver, nil)
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
	return x, y + game.BlockHeight*5 + game.TopMargin
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
func drawWindow(r *ebiten.Image, x, y, width, height int) {
	vector.FillRect(r, float32(x), float32(y), float32(width), float32(height), color.RGBA{A: 192}, false)
}

// DrawTextWithShadow draws a string with a simple drop shadow effect.
func DrawTextWithShadow(
	r *ebiten.Image,
	str string,
	s, x, y int,
	clr color.Color,
	primaryAlign, secondaryAlign text.Align,
) {
	face := &text.GoTextFace{
		Source: fontSource,
		Size:   float64(s),
	}

	// Shadow layer
	shadowOp := &text.DrawOptions{}
	shadowOp.GeoM.Translate(float64(x)+1, float64(y)+1)
	shadowOp.ColorScale.ScaleWithColor(color.RGBA{A: 128})
	shadowOp.LineSpacing = game.FontSize
	shadowOp.PrimaryAlign = primaryAlign
	shadowOp.SecondaryAlign = secondaryAlign
	text.Draw(r, str, face, shadowOp)

	// Text layer
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(clr)
	op.LineSpacing = game.FontSize
	op.PrimaryAlign = primaryAlign
	op.SecondaryAlign = secondaryAlign
	text.Draw(r, str, face, op)
}

func drawTextWindow(r *ebiten.Image, str string, x, y int) {
	DrawTextWithShadow(r, str, game.FontSize, x, y, labelColor, text.AlignStart, text.AlignStart)
	fieldX, _ := FieldWindowPosition()
	drawWindow(r, x, y+game.BlockHeight, game.ScreenWidth-x-fieldX, game.BlockHeight*2)
}
