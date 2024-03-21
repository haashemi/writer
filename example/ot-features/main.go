package main

import (
	"image"
	"image/color"

	"github.com/haashemi/painter"
	"github.com/haashemi/writer"
)

// Load the font faces.
var (
	geist         = writer.NewFaceFromFile("./fonts/GeistMonoNerdFontMono-Regular.otf")
	dancingScript = writer.NewFaceFromFile("./fonts/DancingScript-Medium.ttf")
)

// Create fonts from faces.
var (
	titleFont = writer.NewFont(geist, 80)
	textFont  = writer.NewFont(dancingScript, 130)
)

// A slice of font features we want to enable/use in this example
var OTFeatures = []writer.Feature{
	writer.StandardLigatures(),
	writer.ContextualAlternatives(),
	writer.Fractions(),
	writer.StylisticAlternatives(),
}

func main() {
	// Create a new image. (doesn't matter how.)
	img := painter.New(1080, 1080)
	painter.DrawRadialGradient(img, img.Rect, color.NRGBA{30, 30, 30, 255}, color.NRGBA{10, 10, 10, 255}, image.Point{img.Rect.Dx() / 2, img.Rect.Dy()})

	// Write the header text normally.
	w, _ := writer.NewWriter(titleFont, "Without OT Features", writer.Options{})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 100), image.White)

	// Write a the sample text without any font features.
	w, _ = writer.NewWriter(textFont, "My Test Text 1/2", writer.Options{})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 250), image.White)

	// Write the next header normally.
	w, _ = writer.NewWriter(titleFont, "With OT Features", writer.Options{})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 580), image.White)

	// Write the sample text, but using a few font features!
	w, _ = writer.NewWriter(textFont, "My Test Text 1/2", writer.Options{Features: OTFeatures})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 730), image.White)

	// Save the image. (doesn't matter how.)
	painter.SavePNG(img, "ot-features/result.png")
}
