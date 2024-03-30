package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/haashemi/painter"
	"github.com/haashemi/writer"
)

var (
	DancingScript = writer.NewFaceFromFile("./fonts/DancingScript-Medium.ttf")
	GeistMono     = writer.NewFaceFromFile("./fonts/GeistMonoNerdFontMono-Regular.otf")
	Vazirmatn     = writer.NewFaceFromFile("./fonts/Vazirmatn-ExtraBold.ttf")
)

func main() {
	img := painter.New(1280, 550)
	painter.DrawLinearGradient(img, img.Rect, math.Pi/2, color.NRGBA{}, color.NRGBA{20, 20, 20, 255})

	writeTitle(img)
	writeDescription(img)
	writeRepoURL(img)

	painter.SavePNG(img, "banner/result.png")
}

func writeTitle(img draw.Image) {
	font := writer.NewFont(Vazirmatn, 200)
	defer font.Close()

	w, err := writer.NewWriter(font, "WRITER", writer.Options{})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	textImage := painter.New(w.Bounds().Dx(), w.Bounds().Dy())
	painter.DrawLinearGradient(textImage, textImage.Rect, 200, color.NRGBA{207, 139, 243, 255}, color.NRGBA{253, 185, 155, 255})

	w.Write(img, image.Pt((img.Bounds().Dx()-w.Bounds().Dx())/2, 150), textImage)
}

func writeDescription(img draw.Image) {
	font := writer.NewFont(DancingScript, 50)
	defer font.Close()

	w, err := writer.NewWriter(font, "Write any text on any image easily!", writer.Options{})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	w.Write(img, image.Pt((img.Bounds().Dx()-w.Bounds().Dx())/2, 350), image.NewUniform(color.NRGBA{220, 160, 200, 255}))
}

func writeRepoURL(img draw.Image) {
	font := writer.NewFont(GeistMono, 25)
	defer font.Close()

	w, err := writer.NewWriter(font, "github.com/haashemi/writer", writer.Options{})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	w.Write(img, image.Pt((img.Bounds().Dx()-w.Bounds().Dx())/2, img.Bounds().Dy()-w.Bounds().Dy()-5), image.White)
}
