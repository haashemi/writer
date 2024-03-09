package writer

import (
	"image"
	"image/draw"

	"github.com/haashemi/go-harfbuzz/hb"
	"github.com/mattn/go-pointer"
	"golang.org/x/image/vector"
)

type Writer struct {
	font *Font
	opts Options
	buf  hb.Buffer

	bounds    image.Rectangle
	glyphs    []hb.GlyphInfo
	positions []hb.GlyphPosition
}

func NewWriter(font *Font, text string, opts Options) (w *Writer, err error) {
	w = &Writer{
		font: font,
		opts: opts,
		buf:  hb.BufferCreate(),
	}

	if w.opts.Bidi {
		text, err = bidiText(text)
		if err != nil {
			return nil, err
		}
	}

	hb.BufferAddUTF8(w.buf, text)
	hb.BufferGuessSegmentProperties(w.buf)
	hb.Shape(w.font.font, w.buf, nil)

	// TODO: Let's support these with Options
	// []hb.Feature{
	// 	{Tag: hb.TagFromString("calt"), Value: 1, Start: 0, End: 4294967295},
	// 	{Tag: hb.TagFromString("liga"), Value: 1, Start: 0, End: 4294967295},
	// }

	w.glyphs = hb.BufferGetGlyphInfos(w.buf)
	w.positions = hb.BufferGetGlyphPositions(w.buf)

	return
}

// Advance returns how far the text will go after drawing.
func (w *Writer) Advance() int32 {
	if !w.bounds.Empty() {
		return int32(w.bounds.Dx())
	}

	var advance int32
	for _, gp := range hb.BufferGetGlyphPositions(w.buf) {
		advance += gp.XAdvance
	}

	return advance / 64
}

// Bounds returns the after-drawing text bounds.
func (w *Writer) Bounds() image.Rectangle {
	if !w.bounds.Empty() {
		return w.bounds
	}

	state := new(drawingState)
	state.fontSize = float32(w.font.Size())
	statePointer := pointer.Save(state)

	for i, gi := range w.glyphs {
		gp := w.positions[i]

		state.offX = float32(gp.XOffset) / 64
		state.offY = float32(gp.YOffset) / 64

		w.font.draw(gi.Codepoint, drawFuncs, statePointer)

		state.posX += float32(gp.XAdvance / 64)
		state.posY += float32(gp.YAdvance / 64)
	}

	w.bounds = image.Rect(0, int(-state.topY), int(state.posX), int(state.bottomY-state.topY))
	return w.bounds
}

// Write draws the text on the “image” at “at” with "color”.
func (w *Writer) Write(img draw.Image, at image.Point, color image.Image) {
	bounds := w.Bounds()

	state := new(drawingState)
	state.vec = vector.NewRasterizer(bounds.Dx(), bounds.Dy())
	state.fontSize = float32(w.font.Size())
	statePointer := pointer.Save(state)

	for i, gi := range w.glyphs {
		gp := w.positions[i]

		state.offX = float32(gp.XOffset) / 64
		state.offY = float32(gp.YOffset) / 64

		w.font.draw(gi.Codepoint, drawFuncs, statePointer)

		state.posX += float32(gp.XAdvance / 64)
		state.posY += float32(gp.YAdvance / 64)
	}

	state.vec.Draw(img, bounds.Add(at), color, image.Point{})
}

func (w *Writer) Close() error {
	hb.BufferDestroy(w.buf)
	return nil
}
