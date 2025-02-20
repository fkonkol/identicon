package identicon_test

import (
	"crypto/md5"
	"testing"

	"github.com/fkonkol/identicon/identicon"
)

func BenchmarkNewIdenticon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		identifier := "fkonkol"
		checksum := md5.Sum([]byte(identifier))

		identicon.New(
			identicon.WithSource(checksum[:]),
		)
	}
}

func BenchmarkIdenticonEncodePNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		identifier := "fkonkol"
		checksum := md5.Sum([]byte(identifier))

		icon, _ := identicon.New(
			identicon.WithSource(checksum[:]),
		)

		icon.Bytes()
	}
}
