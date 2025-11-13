package render

import (
	"image/color"

	"github.com/xiao-dong-li/tennis/game"
)

type BlinkStyle int32

const (
	BlinkGold BlinkStyle = iota
	BlinkBlue
	BlinkNeon
	BlinkRed
	BlinkAqua
	BlinkSilver
	BlinkWhite
)

// BlinkColor returns a dynamic color that "breathes"/flashes over time.
// Supported styles: "gold", "blue", "neon", "red", "aqua", "silver".
func BlinkColor(style BlinkStyle) color.RGBA {
	phase := game.CalcAlpha()

	switch style {
	case BlinkGold:
		return color.RGBA{
			R: uint8(180 + 60*phase),
			G: uint8(140 + 80*phase),
			B: uint8(40 + 40*phase),
			A: 255,
		}
	case BlinkBlue:
		return color.RGBA{
			R: uint8(60 * phase),
			G: uint8(160 * phase),
			B: uint8(255),
			A: 255,
		}
	case BlinkNeon:
		return color.RGBA{
			R: uint8(200 + 55*phase),
			G: uint8(60 + 20*phase),
			B: uint8(200 + 55*phase),
			A: 255,
		}
	case BlinkRed:
		return color.RGBA{
			R: uint8(200 + 55*phase),
			G: uint8(80 + 40*phase),
			B: uint8(50),
			A: 255,
		}
	case BlinkAqua:
		return color.RGBA{
			R: uint8(50 * phase),
			G: uint8(255),
			B: uint8(180 + 75*phase),
			A: 255,
		}
	case BlinkSilver:
		return color.RGBA{
			R: uint8(180 + 75*phase),
			G: uint8(180 + 75*phase),
			B: uint8(200 + 55*phase),
			A: 255,
		}
	default:
		return color.RGBA{
			R: uint8(200 + 55*phase),
			G: uint8(200 + 55*phase),
			B: uint8(200 + 55*phase),
			A: 255,
		}
	}
}
