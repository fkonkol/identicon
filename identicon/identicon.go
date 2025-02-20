package identicon

import (
	"image"
	"image/color"
	"iter"
	"math/rand/v2"
)

type Identicon struct {
	source []byte
	size   int
}

func New(source []byte) Identicon {
	return Identicon{
		source: source,
		size:   448,
	}
}

// TODO: Always generate pretty colors 🦄 Should probably be HSL/HSB based.
func (i *Identicon) Foreground() color.RGBA {
	r := uint8(128 + rand.IntN(129))
	g := uint8(128 + rand.IntN(129))
	b := uint8(128 + rand.IntN(129))
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func (i *Identicon) Rect(img *image.RGBA, x0, y0, x1, y1 int, color color.RGBA) {
	for x := x0; x < x0+x1; x++ {
		for y := y0; y < y0+y1; y++ {
			img.SetRGBA(x, y, color)
		}
	}
}

func (i *Identicon) Nibbles() iter.Seq[uint8] {
	return func(yield func(uint8) bool) {
		for _, s := range i.source {
			hi := s & 0xf0 >> 4
			lo := s & 0x0f

			if !yield(hi) || !yield(lo) {
				return
			}
		}
	}
}

func (i *Identicon) Pixels() [5][5]bool {
	pixels := [5][5]bool{}

	next, stop := iter.Pull(i.Nibbles())
	defer stop()

	for row := 0; row < 5; row++ {
		for col := 0; col < 5/2+1; col++ {
			nibble, _ := next()
			paint := nibble&1 == 0

			pixels[row][col], pixels[row][5-col-1] = paint, paint
		}
	}

	return pixels
}

func (i *Identicon) Image() *image.RGBA {
	padding := 64
	background := color.RGBA{R: 240, G: 240, B: 240, A: 255}
	foreground := i.Foreground()

	image := image.NewRGBA(image.Rect(0, 0, i.size, i.size))
	i.Rect(image, 0, 0, i.size, i.size, background)

	pixels := i.Pixels()
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if pixels[y][x] {
				i.Rect(image, x*64+padding, y*64+padding, 64, 64, foreground)
			}
		}
	}

	return image
}
