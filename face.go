package writer

import (
	"io"

	"github.com/haashemi/go-harfbuzz/hb"
)

type Face struct {
	io.Closer

	blob hb.Blob
	face hb.Face
}

func (f *Face) Close() {
	hb.FaceDestroy(f.face)
	hb.BlobDestroy(f.blob)
}

func NewFace(data []byte) *Face {
	blob := hb.BlobCreate(data, hb.MemoryModeDuplicate, nil, nil)
	face := hb.FaceCreate(blob, 0)

	return &Face{
		blob: blob,
		face: face,
	}
}

func NewFaceFromFile(filename string) *Face {
	blob := hb.BlobCreateFromFile(filename)
	face := hb.FaceCreate(blob, 0)

	return &Face{
		blob: blob,
		face: face,
	}
}
