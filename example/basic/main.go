package main

import (
	"image"
	"image/png"
	"os"

	"github.com/haashemi/writer"
)

func main() {
	// Load face from the font file
	face := writer.NewFaceFromFile("./font.ttf")
	defer face.Close()

	// Create a font from the face
	font := writer.NewFont(face, 45)
	defer font.Close()

	// Create a new writer instance.
	w, _ := writer.NewWriter(font, "Hello World!", writer.Options{})
	defer w.Close()

	// Create a new image. (doesn't matter how.)
	img := image.NewNRGBA(w.Bounds())

	// Write it on your image
	w.Write(img, w.Bounds(), image.White)

	// Save the image. (doesn't matter how.)
	f, _ := os.Create("basic/result.png")
	png.Encode(f, img)
}
