package writer

import (
	"io"
	"unsafe"

	"github.com/haashemi/go-harfbuzz/hb"
)

type FontExtents = hb.FontExtents

type Font struct {
	io.Closer

	font hb.Font
}

func NewFont(face *Face, size int32) *Font {
	font := &Font{font: hb.FontCreate(face.face)}
	font.SetSize(size)
	return font
}

func (f *Font) Extents() (FontExtents, bool) {
	return hb.FontGetHExtents(f.font)
}

func (f *Font) Size() int32 {
	_, h := hb.FontGetScale(f.font)
	return h / 64
}

func (f *Font) SetSize(size int32) {
	hb.FontSetScale(f.font, size*64, size*64)
}

func (f *Font) Close() {
	hb.FontDestroy(f.font)
}

func (f *Font) draw(glyph uint32, dfuncs hb.DrawFuncs, drawData unsafe.Pointer) {
	hb.FontDrawGlyph(f.font, glyph, dfuncs, drawData)
}
