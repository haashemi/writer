package writer

// #cgo pkg-config: harfbuzz
// #include "draw_funcs.h"
import "C"
import (
	"unsafe"

	"github.com/haashemi/go-harfbuzz/hb"
	"github.com/mattn/go-pointer"
)

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
func drawMoveTo(drawData unsafe.Pointer, toX, toY float32) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x = w.posX + w.offX + (toX / 64)
		y = w.posY - w.offY - (toY / 64)
	)

	if w.vec != nil {
		w.vec.MoveTo(x-w.minX, y-w.minY)
	} else {
		w.processY(y)
		w.processX(x)
	}
}

//export drawLineTo
func drawLineTo(drawData unsafe.Pointer, toX, toY float32) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x = w.posX + w.offX + (toX / 64)
		y = w.posY - w.offY - (toY / 64)
	)

	if w.vec != nil {
		w.vec.LineTo(x-w.minX, y-w.minY)
	} else {
		w.processY(y)
		w.processX(x)
	}
}

//export drawQuadraticTo
func drawQuadraticTo(drawData unsafe.Pointer, controlX, controlY, toX, toY float32) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x1 = w.posX + w.offX + (controlX / 64)
		y1 = w.posY - w.offY - (controlY / 64)
		x2 = w.posX + w.offX + (toX / 64)
		y2 = w.posY - w.offY - (toY / 64)
	)

	if w.vec != nil {
		w.vec.QuadTo(x1-w.minX, y1-w.minY, x2-w.minX, y2-w.minY)
	} else {
		w.processY(y1, y2)
		w.processX(x1, x2)
	}
}

//export drawCubicTo
func drawCubicTo(drawData unsafe.Pointer, control1X, control1Y, control2X, control2Y, toX, toY float32) {
	w := pointer.Restore(drawData).(*drawingState)

	var (
		x1 = w.posX + w.offX + (control1X / 64)
		y1 = w.posY - w.offY - (control1Y / 64)
		x2 = w.posX + w.offX + (control2X / 64)
		y2 = w.posY - w.offY - (control2Y / 64)
		x3 = w.posX + w.offX + (toX / 64)
		y3 = w.posY - w.offY - (toY / 64)
	)

	if w.vec != nil {
		w.vec.CubeTo(x1-w.minX, y1-w.minY, x2-w.minX, y2-w.minY, x3-w.minX, y3-w.minY)
	} else {
		w.processY(y1, y2, y3)
		w.processX(x1, x2, x3)
	}
}

//export drawClosePath
func drawClosePath(drawData unsafe.Pointer) {
	vec := pointer.Restore(drawData).(*drawingState).vec
	if vec != nil {
		vec.ClosePath()
	}
}
