package drawer

import (
	"golang.org/x/image/vector"
)

// DrawingState holds the essential information of drawing glyphs.
type DrawingState struct {
	Vec *vector.Rasterizer

	PosX, PosY float32
	OffX, OffY float32

	MinY, MaxY float32
	MinX, MaxX float32
}

// processY sets the highest and the lowest Y values passed to the draw methods
// to the current DrawingState.
func (ds *DrawingState) processY(ys ...float32) {
	for _, y := range ys {
		if y > ds.MaxY {
			ds.MaxY = y
		}

		if ds.MinY == 0 || y < ds.MinY {
			ds.MinY = y
		}
	}
}

func (ds *DrawingState) processX(xs ...float32) {
	for _, x := range xs {
		if x > ds.MaxX {
			ds.MaxX = x
		}

		if ds.MinX == 0 || x < ds.MinX {
			ds.MinX = x
		}
	}
}
