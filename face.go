package writer

import (
	"io"

	"github.com/haashemi/go-harfbuzz/hb"
)

// Just to make sure if [Face] implements [io.Closer].
var _ io.Closer = &Face{}

// Face holds a harfbuzz Face object.
type Face struct {
	io.Closer

	blob hb.Blob
	face hb.Face
}

// Close destroys the harfbuzz Blob and Face objects and frees the memory.
func (f *Face) Close() error {
	hb.FaceDestroy(f.face)
	hb.BlobDestroy(f.blob)
	return nil
}

// NewFace returns a newly initialized harfbuzz Face object from “data” bytes.
func NewFace(data []byte) *Face {
	blob := hb.BlobCreate(data, hb.MemoryModeDuplicate, nil, nil)
	face := hb.FaceCreate(blob, 0)

	return &Face{
		blob: blob,
		face: face,
	}
}

// NewFaceFromFile returns a newly initialized harfbuzz Face object from “filename” file.
func NewFaceFromFile(filename string) *Face {
	blob := hb.BlobCreateFromFile(filename)
	face := hb.FaceCreate(blob, 0)

	return &Face{
		blob: blob,
		face: face,
	}
}
