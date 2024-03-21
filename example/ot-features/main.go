package main

import (
	"image"
	"image/color"

	"github.com/haashemi/painter"
	"github.com/haashemi/writer"
)

var (
	geist         = writer.NewFaceFromFile("./fonts/GeistMonoNerdFontMono-Regular.otf")
	dancingScript = writer.NewFaceFromFile("./fonts/DancingScript-Medium.ttf")
)

var (
	titleFont = writer.NewFont(geist, 80)
	textFont  = writer.NewFont(dancingScript, 130)
)

func main() {
	img := painter.New(1080, 1080)
	painter.DrawRadialGradient(img, img.Rect, color.NRGBA{30, 30, 30, 255}, color.NRGBA{10, 10, 10, 255}, image.Point{img.Rect.Dx() / 2, img.Rect.Dy()})

	w, _ := writer.NewWriter(titleFont, "Without OT Features", writer.Options{})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 100), image.White)

	w, _ = writer.NewWriter(textFont, "My Test Text 1/2", writer.Options{})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 250), image.White)

	w, _ = writer.NewWriter(titleFont, "With OT Features", writer.Options{})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 580), image.White)

	w, _ = writer.NewWriter(textFont, "My Test Text 1/2", writer.Options{Features: []writer.Feature{writer.StandardLigatures(), writer.ContextualAlternatives(), writer.Fractions(), writer.StylisticAlternatives()}})
	w.Write(img, image.Pt((1080-w.Bounds().Dx())/2, 730), image.White)

	painter.SavePNG(img, "ot-features/result.png")
}
