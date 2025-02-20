package main

import (
	"crypto/md5"
	"image/png"
	"os"

	"github.com/fkonkol/identicon/identicon"
)

func main() {
	name := "fkonkole"
	sum := md5.Sum([]byte(name))

	icon, err := identicon.New(
		identicon.WithSource(sum[:]),
		identicon.WithPadding(100),
		identicon.WithSize(1000),
	)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(name + ".png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(file, icon.Image())
	if err != nil {
		panic(err)
	}
}
