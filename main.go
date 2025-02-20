package main

import (
	"crypto/md5"
	"image/png"
	"os"

	"github.com/fkonkol/identicon/identicon"
)

func main() {
	name := "fkonkol"
	sum := md5.Sum([]byte(name))

	icon := identicon.New(sum[:])

	file, err := os.Create(name + ".png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(file, icon.Image())
	if err != nil {
		panic(err)
	}
}
