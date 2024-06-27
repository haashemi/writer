package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFontSize(t *testing.T) {
	face := &Face{} // Mock Face object
	size := int32(12)

	font := NewFont(face, size)
	fontSize := font.Size()

	assert.Equal(t, size, fontSize)
}

func TestSetFontReSize(t *testing.T) {
	face := &Face{} // Mock Face object
	size := int32(12)

	font := NewFont(face, size)
	newSize := int32(16)
	font.SetSize(newSize)

	assert.Equal(t, newSize, font.Size())
}
