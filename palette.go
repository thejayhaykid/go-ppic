package ppic

import "image/color"

// Palette represents a pair of colors to use in image generation.
type Palette struct {
	Foreground color.Color
	Background color.Color
}

// Palette returns a color palette with the first color being the background and the second being the foreground.
func (p Palette) Palette() color.Palette {
	return color.Palette{p.Background, p.Foreground}
}

// Inverse returns the inverse of this color palette (i.e. foreground and background swapped).
func (p Palette) Inverse() Palette {
	return Palette{p.Background, p.Foreground}
}

// DefaultPalette is the default black and white color palette.
var DefaultPalette = Palette{Foreground: color.Black, Background: color.White}

// Palettes is a map of color palettes provided by the library.
var Palettes = map[string]Palette{
	"Default": DefaultPalette,
	"Candy": {
		Foreground: color.RGBA{R: 0xF0, G: 0x54, B: 0x4F, A: 0xFF},
		Background: color.RGBA{R: 0x40, G: 0x4E, B: 0x4D, A: 0xFF},
	},
}
