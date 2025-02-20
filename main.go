package main

import (
	"crypto/md5"
	"os"

	"github.com/fkonkol/identicon/identicon"
)

func saveIdenticonToFile(icon *identicon.Identicon, name string) error {
	file, err := os.Create(name + ".png")
	if err != nil {
		return err
	}

	bytes, err := icon.Bytes()
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	return err
}

func main() {
	name := "fkonkol"
	sum := md5.Sum([]byte(name))

	icon, err := identicon.New(
		identicon.WithSource(sum[:]),
		identicon.WithPadding(100),
		identicon.WithSize(1000),
	)
	if err != nil {
		panic(err)
	}

	if err := saveIdenticonToFile(icon, name); err != nil {
		panic(err)
	}
}
