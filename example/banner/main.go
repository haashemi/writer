package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/haashemi/painter"
	"github.com/haashemi/writer"
)

func main() {
	face := writer.NewFaceFromFile("./font.ttf")
	defer face.Close()

	img := painter.New(1280, 720)
	painter.DrawRadialGradient(img, img.Rect, color.NRGBA{45, 45, 45, 255}, color.NRGBA{20, 20, 20, 255}, image.Point{img.Rect.Dx() / 2, img.Rect.Dy() / 2})

	writeTitle(face, img)
	writeDescription(face, img)
	writeRepoURL(face, img)

	painter.SavePNG(img, "banner/result.png")
}

func writeTitle(face *writer.Face, img draw.Image) {
	font := writer.NewFont(face, 200)
	defer font.Close()

	w, err := writer.NewWriter(font, "WRITER", writer.Options{})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	bounds := w.Bounds()

	textImage := painter.New(bounds.Dx(), bounds.Dy())
	painter.DrawLinearGradient(textImage, textImage.Rect, 200, color.NRGBA{207, 139, 243, 255}, color.NRGBA{253, 185, 155, 255})

	w.Write(img, image.Pt((img.Bounds().Dx()-bounds.Dx())/2, 230), textImage)
}

func writeDescription(face *writer.Face, img draw.Image) {
	font := writer.NewFont(face, 35)
	defer font.Close()

	w, err := writer.NewWriter(font, "Write any text on any image easily!", writer.Options{})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	bounds := w.Bounds()

	w.Write(img, image.Pt((img.Bounds().Dx()-bounds.Dx())/2, 430), image.NewUniform(color.NRGBA{220, 160, 200, 255}))
}

func writeRepoURL(face *writer.Face, img draw.Image) {
	font := writer.NewFont(face, 25)
	defer font.Close()

	w, err := writer.NewWriter(font, "github.com/haashemi/writer", writer.Options{})
	if err != nil {
		panic(err)
	}
	defer w.Close()

	bounds := w.Bounds()

	w.Write(img, image.Pt(10, img.Bounds().Dy()-bounds.Dy()-5), image.White)
}
