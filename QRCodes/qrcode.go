package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello QR Code")

	file, err := os.Create("qrcode.png")
	must(err)
	defer file.Close()

	err = GenerateQRCode(file, "555-2368", Version(1))
	must(err)
}

// Version determines qr version and size
type Version int8

// GenerateQRCode creates a QR code!
func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := version.PatternSize()
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

func (v Version) PatternSize() int {
	return 4*int(v) + 17
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
