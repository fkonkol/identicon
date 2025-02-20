package identicon

import (
	"image/color"
	"math"
)

type HSL struct {
	H float64
	S float64
	L float64
}

func (hsl HSL) ToRGBA() color.RGBA {
	h, s, l := hsl.H, hsl.S, hsl.L

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r1, g1, b1 float64
	switch {
	case 0 <= h && h < 60:
		r1, g1, b1 = c, x, 0
	case 60 <= h && h < 120:
		r1, g1, b1 = x, c, 0
	case 120 <= h && h < 180:
		r1, g1, b1 = 0, c, x
	case 180 <= h && h < 240:
		r1, g1, b1 = 0, x, c
	case 240 <= h && h < 300:
		r1, g1, b1 = x, 0, c
	case 300 <= h && h < 360:
		r1, g1, b1 = c, 0, x
	}

	r := uint8(math.Round((r1 + m) * 255))
	g := uint8(math.Round((g1 + m) * 255))
	b := uint8(math.Round((b1 + m) * 255))

	return color.RGBA{R: r, G: g, B: b, A: 255}
}
