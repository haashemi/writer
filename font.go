package writer

import (
	"io"
	"unsafe"

	"github.com/haashemi/writer/internal/hb"
)

// FontExtents is just an alias to harfbuzz's FontExtents type.
type FontExtents = hb.FontExtents

// Just to make sure if [Font] implements [io.Closer].
var _ io.Closer = &Font{}

// Font holds a harfbuzz Font object.
type Font struct {
	io.Closer

	font hb.Font
}

// NewFont returns a new Font from the “face” with size of “size”.
func NewFont(face *Face, size int32) *Font {
	font := &Font{font: hb.FontCreate(face.face)}
	font.SetSize(size)
	return font
}

// Extents returns the FontExtents of the Font.
func (f *Font) Extents() (FontExtents, bool) {
	return hb.FontGetHExtents(f.font)
}

// Size returns the font size.
//
// TODO: It may not be accurate if scaled manually. A better way should be found.
func (f *Font) Size() int32 {
	_, h := hb.FontGetScale(f.font)
	return h / 64
}

// SetSize updates the font size.
func (f *Font) SetSize(size int32) {
	hb.FontSetScale(f.font, size*64, size*64)
}

// Close destroys the font and frees the memory.
func (f *Font) Close() error {
	hb.FontDestroy(f.font)
	return nil
}

func (f *Font) draw(glyph uint32, dfuncs hb.DrawFuncs, drawData unsafe.Pointer) {
	hb.FontDrawGlyph(f.font, glyph, dfuncs, drawData)
}
