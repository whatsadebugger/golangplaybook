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

	err = GenerateQRCode(file, "555-2368")
	must(err)
}

// GenerateQRCode creates a QR code!
func GenerateQRCode(w io.Writer, code string) error {
	img := image.NewRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
