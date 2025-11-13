package sceen

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images/blocks"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/xiao-dong-li/tennis/game"
	"github.com/xiao-dong-li/tennis/game/render"
)

var imageBackground *ebiten.Image

type TitleScene struct {
	offset float64 // current scroll offset
	speed  float64 // scroll speed
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(blocks.Background_png))
	if err != nil {
		log.Fatal(err)
	}
	imageBackground = ebiten.NewImageFromImage(img)
}

func NewTitleScene() Scene {
	return &TitleScene{
		speed: 0.3,
	}
}

func (t *TitleScene) Update(g *GameState) {
	t.offset += t.speed

	if g.Input.IsSpace() {
		g.SceneManager.GoTo(NewGameScene())
	}
}

func (t *TitleScene) Draw(r *ebiten.Image) {
	drawTitleBackground(r, t.offset)

	// Draw game title
	render.DrawTextWithShadow(
		r,
		strings.ToUpper(game.Title),
		game.TitleFontSize,
		game.ScreenWidth/2,
		game.TitleFontSize,
		color.RGBA{B: 128, A: 255},
		text.AlignCenter,
		text.AlignStart,
	)

	render.DrawTextWithShadow(
		r,
		"BY XIAODONG LI",
		game.FontSize,
		game.ScreenWidth/2,
		game.ScreenHeight-game.TitleFontSize*2,
		render.BlinkColor(render.BlinkGold),
		text.AlignCenter,
		text.AlignEnd,
	)

	// Draw start prompt
	render.DrawTextWithShadow(
		r,
		"PRESS SPACE TO START",
		game.FontSize,
		game.ScreenWidth/2,
		game.ScreenHeight-game.TitleFontSize,
		color.RGBA{R: 128, A: 255},
		text.AlignCenter,
		text.AlignEnd,
	)
}

// drawTitleBackground draws a looping, scrolling background for the title screen.
func drawTitleBackground(r *ebiten.Image, offset float64) {
	w, h := imageBackground.Bounds().Dx(), imageBackground.Bounds().Dy()

	// Wrap offset so it repeats smoothly
	ox := int(offset) % w
	oy := int(offset) % h

	// Draw the tiled background
	for y := -h; y < game.ScreenHeight+h; y += h {
		for x := -w; x < game.ScreenWidth+w; x += w {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x+ox), float64(y+oy))
			r.DrawImage(imageBackground, op)
		}
	}
}
