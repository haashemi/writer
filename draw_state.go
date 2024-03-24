package writer

import (
	"golang.org/x/image/vector"
)

// drawingState holds the essential information of drawing glyphs.
type drawingState struct {
	vec *vector.Rasterizer

	posX, posY float32
	offX, offY float32

	minY, maxY float32
	minX, maxX float32
}

// processY sets the highest and the lowest Y values passed to the draw methods
// to the current drawingState.
func (ds *drawingState) processY(ys ...float32) {
	for _, y := range ys {
		if y > ds.maxY {
			ds.maxY = y
		}

		if ds.minY == 0 || y < ds.minY {
			ds.minY = y
		}
	}
}

func (ds *drawingState) processX(xs ...float32) {
	for _, x := range xs {
		if x > ds.maxX {
			ds.maxX = x
		}

		if ds.minX == 0 || x < ds.minX {
			ds.minX = x
		}
	}
}
