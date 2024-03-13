package main

import (
	"image"
	"image/color"

	"github.com/haashemi/painter"
	"github.com/haashemi/writer"
)

const (
	Text     = "A placeholder text showing how to write a text with gradient colors"
	FontPath = "./fonts/Vazirmatn-ExtraBold.ttf" // update it with a path to your font
)

func main() {
	// Load face from the font file
	face := writer.NewFaceFromFile(FontPath)
	defer face.Close()

	// Create a font from the face
	font := writer.NewFont(face, 40)
	defer font.Close()

	// Create a new image. (doesn't matter how.)
	img := painter.New(1280, 720)
	painter.DrawRadialGradient(img, img.Rect, color.NRGBA{45, 45, 45, 255}, color.NRGBA{20, 20, 20, 255}, image.Point{1150, 550})

	// Create a new writer instance.
	w, err := writer.NewWriter(font, Text, writer.Options{Bidi: true})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// Calculate the text bounds
	bounds := w.Bounds()
	x := (img.Rect.Dx() - bounds.Dx()) / 2
	y := (img.Rect.Dy() - bounds.Dy()) / 2

	// Create an image gradient for our text
	textImage := painter.New(bounds.Dx(), bounds.Dy())
	painter.DrawLinearGradient(textImage, textImage.Rect, 0, color.NRGBA{207, 139, 243, 255}, color.NRGBA{253, 185, 155, 255})

	// Write it on your image
	w.Write(img, image.Pt(x, y), textImage)

	// Save the image. (doesn't matter how.)
	painter.SavePNG(img, "gradient-text/result.png")
}
