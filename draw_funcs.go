package writer

// #cgo pkg-config: harfbuzz
// #include "draw_funcs.h"
import "C"
import (
	"unsafe"

	"github.com/haashemi/go-harfbuzz/hb"
	"github.com/mattn/go-pointer"
	"golang.org/x/image/vector"
)

// drawingState holds the essential information of drawing glyphs.
type drawingState struct {
	vec *vector.Rasterizer

	posX, posY float32
	offX, offY float32
	fontSize   float32

	topY, bottomY float32
}

// processY sets the highest and the lowest Y values passed to the draw methods
// to the current drawingState.
func (ds *drawingState) processY(ys ...float32) {
	for _, y := range ys {
		if y > ds.bottomY {
			ds.bottomY = y
		}

		if ds.topY == 0 || y < ds.topY {
			ds.topY = y
		}
	}
}

// drawFuncs contains the internal drawing methods used to pass to the harfbuzz
// hb_font_draw method for drawing the glyphs.
var drawFuncs = hb.DrawFuncsCreate()

// This methods adds all drawing methods to drawFuncs variable.
func init() {
	hb.DrawFuncsSetMoveToFunc(drawFuncs, (hb.DrawMoveToFunc)(C.move_to), nil, nil)
	hb.DrawFuncsSetLineToFunc(drawFuncs, (hb.DrawLineToFunc)(C.line_to), nil, nil)
	hb.DrawFuncsSetQuadraticToFunc(drawFuncs, (hb.DrawQuadraticToFunc)(C.quadratic_to), nil, nil)
	hb.DrawFuncsSetCubicToFunc(drawFuncs, (hb.DrawCubicToFunc)(C.cubic_to), nil, nil)
	hb.DrawFuncsSetClosePathFunc(drawFuncs, (hb.DrawClosePathFunc)(C.close_path), nil, nil)
}

//export drawMoveTo
func drawMoveTo(drawData unsafe.Pointer, toX, toY C.float) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x = w.posX + w.offX + float32(toX)/64
		y = w.fontSize + w.posY - w.offY - (float32(toY) / 64)
	)

	w.processY(y)
	if w.vec != nil {
		w.vec.MoveTo(x, y)
	}
}

//export drawLineTo
func drawLineTo(drawData unsafe.Pointer, toX, toY C.float) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x = w.posX + w.offX + float32(toX)/64
		y = w.fontSize + w.posY - w.offY - (float32(toY) / 64)
	)

	w.processY(y)
	if w.vec != nil {
		w.vec.LineTo(x, y)
	}
}

//export drawQuadraticTo
func drawQuadraticTo(drawData unsafe.Pointer, controlX, controlY, toX, toY C.float) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x1 = w.posX + w.offX + float32(controlX)/64
		y1 = w.fontSize + w.posY - w.offY - (float32(controlY) / 64)
		x2 = w.posX + w.offX + float32(toX)/64
		y2 = w.fontSize + w.posY - w.offY - (float32(toY) / 64)
	)

	w.processY(y1, y2)
	if w.vec != nil {
		w.vec.QuadTo(x1, y1, x2, y2)
	}
}

//export drawCubicTo
func drawCubicTo(drawData unsafe.Pointer, control1X, control1Y, control2X, control2Y, toX, toY C.float) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x1 = w.posX + w.offX + float32(control1X)/64
		y1 = w.fontSize + w.posY - w.offY - (float32(control1Y) / 64)
		x2 = w.posX + w.offX + float32(control2X)/64
		y2 = w.fontSize + w.posY - w.offY - (float32(control2Y) / 64)
		x3 = w.posX + w.offX + float32(toX)/64
		y3 = w.fontSize + w.posY - w.offY - (float32(toY) / 64)
	)

	w.processY(y1, y2, y3)
	if w.vec != nil {
		w.vec.CubeTo(x1, y1, x2, y2, x3, y3)
	}
}

//export drawClosePath
func drawClosePath(drawData unsafe.Pointer) {
	vec := pointer.Restore(drawData).(*drawingState).vec
	if vec != nil {
		vec.ClosePath()
	}
}
