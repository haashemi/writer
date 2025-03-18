package writer

import (
	"image"
	"image/draw"
	"math"

	"github.com/haashemi/writer/internal/drawer"
	"github.com/haashemi/writer/internal/hb"
	"github.com/mattn/go-pointer"
	"golang.org/x/image/vector"
)

type Writer struct {
	font *Font
	opts Options
	buf  hb.Buffer

	minX, minY float32
	textBounds image.Rectangle

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
	hb.Shape(w.font.font, w.buf, opts.Features)

	w.glyphs = hb.BufferGetGlyphInfos(w.buf)
	w.positions = hb.BufferGetGlyphPositions(w.buf)

	return
}

// Advance returns how far the text will go after drawing.
func (w *Writer) Advance() int {
	if w.textBounds.Empty() {
		return w.Bounds().Dx()
	}

	return w.textBounds.Dx()
}

// Bounds returns the after-drawing text bounds.
func (w *Writer) Bounds() image.Rectangle {
	if !w.textBounds.Empty() {
		return w.textBounds
	}

	state := new(drawer.DrawingState)
	statePointer := pointer.Save(state)
	defer pointer.Unref(statePointer)

	for i, gi := range w.glyphs {
		gp := w.positions[i]

		state.OffX = float32(gp.XOffset) / 64
		state.OffY = float32(gp.YOffset) / 64

		w.font.draw(gi.Codepoint, drawer.DrawFuncs, statePointer)

		state.PosX += float32(gp.XAdvance) / 64
		state.PosY += float32(gp.YAdvance) / 64
	}

	w.minY, w.minX = state.MinY, state.MinX
	w.textBounds = image.Rect(0, 0, int(math.Ceil(float64(state.MaxX-state.MinX))), int(math.Ceil(float64(state.MaxY-state.MinY))))
	return w.textBounds
}

// Write draws the text on the “image” at “at” with "color”.
func (w *Writer) Write(img draw.Image, at image.Point, color image.Image) {
	bounds := w.Bounds()

	state := &drawer.DrawingState{
		Vec:  vector.NewRasterizer(bounds.Dx(), bounds.Dy()),
		MinX: w.minX, MinY: w.minY,
	}

	statePointer := pointer.Save(state)
	defer pointer.Unref(statePointer)

	for i, gi := range w.glyphs {
		gp := w.positions[i]

		state.OffX = float32(gp.XOffset) / 64
		state.OffY = float32(gp.YOffset) / 64

		w.font.draw(gi.Codepoint, drawer.DrawFuncs, statePointer)

		state.PosX += float32(gp.XAdvance / 64)
		state.PosY += float32(gp.YAdvance / 64)
	}

	state.Vec.Draw(img, state.Vec.Bounds().Add(at), color, image.Point{})
}

// Close destroys the writer's buffer and frees the memory.
func (w *Writer) Close() error {
	hb.BufferDestroy(w.buf)
	return nil
}
