package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/haashemi/writer"
)

const (
	Text     = "حتی متن bi-directional با عدد (1234) هم درست می نویسه!"
	FontPath = "./font.ttf" // update it with a path to your font
)

func main() {
	// Load face from the font file
	face := writer.NewFaceFromFile(FontPath)
	defer face.Close()

	// Create a font from the face
	font := writer.NewFont(face, 45)
	defer font.Close()

	// Create a new image. (doesn't matter how.)
	img := image.NewNRGBA(image.Rect(0, 0, 1280, 720))
	draw.Draw(img, img.Rect, image.Black, image.Point{}, draw.Over)

	// Create a new writer instance.
	w, err := writer.NewWriter(font, Text, writer.Options{Bidi: true})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// Get the text bounds and calculate the position to write the text at.
	bounds := w.Bounds()
	x := (img.Rect.Dx() - bounds.Dx()) / 2
	y := (img.Rect.Dy() - bounds.Dy()) / 2

	// Write it on your image
	w.Write(img, image.Pt(x, y), image.White)

	// Save the image. (doesn't matter how.)
	f, err := os.Create("bidi-text/result.png")
	if err != nil {
		panic(err)
	}

	if err = png.Encode(f, img); err != nil {
		panic(err)
	}
}
