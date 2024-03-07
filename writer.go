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

	return
}

// TODO: Calculate Arabic vowels
func (w *Writer) Bounds() (rect image.Rectangle) {
	var width int32
	for _, gi := range hb.BufferGetGlyphPositions(w.buf) {
		width += gi.XAdvance / 64
	}
	rect.Max.X = int(width)

	e, ok := w.font.Extents()
	if ok {
		rect.Max.Y = int(e.Ascender/64 - e.Descender/64)
	}

	return rect
}

func (w *Writer) Write(img draw.Image, rect image.Rectangle, color image.Image) {
	gis := hb.BufferGetGlyphInfos(w.buf)
	gps := hb.BufferGetGlyphPositions(w.buf)

	state := new(drawingState)
	state.vec = vector.NewRasterizer(rect.Dx(), rect.Dy())
	state.fontSize = float32(w.font.Size())
	statePointer := pointer.Save(state)

	for i := range len(gis) {
		gi := gis[i]
		gp := gps[i]

		state.offX = float32(gp.XOffset) / 64
		state.offY = float32(gp.YOffset) / 64

		w.font.draw(gi.Codepoint, drawFuncs, statePointer)

		state.posX += float32(gp.XAdvance / 64)
		state.posY += float32(gp.YAdvance / 64)
	}

	state.vec.Draw(img, rect, color, image.Point{})
}

func (w *Writer) Close() error {
	hb.BufferDestroy(w.buf)
	return nil
}
