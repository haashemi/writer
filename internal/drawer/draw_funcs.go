package drawer

// #cgo pkg-config: harfbuzz
// #include "draw_funcs.h"
import "C"
import (
	"unsafe"

	"github.com/haashemi/writer/internal/hb"
	"github.com/mattn/go-pointer"
)

// DrawFuncs contains the internal drawing methods used to pass to the harfbuzz
// hb_font_draw method for drawing the glyphs.
var DrawFuncs = hb.DrawFuncsCreate()

// This methods adds all drawing methods to DrawFuncs variable.
func init() {
	hb.DrawFuncsSetMoveToFunc(DrawFuncs, (hb.DrawMoveToFunc)(C.move_to), nil, nil)
	hb.DrawFuncsSetLineToFunc(DrawFuncs, (hb.DrawLineToFunc)(C.line_to), nil, nil)
	hb.DrawFuncsSetQuadraticToFunc(DrawFuncs, (hb.DrawQuadraticToFunc)(C.quadratic_to), nil, nil)
	hb.DrawFuncsSetCubicToFunc(DrawFuncs, (hb.DrawCubicToFunc)(C.cubic_to), nil, nil)
	hb.DrawFuncsSetClosePathFunc(DrawFuncs, (hb.DrawClosePathFunc)(C.close_path), nil, nil)
}

//export drawMoveTo
func drawMoveTo(drawData unsafe.Pointer, toX, toY float32) {
	w := pointer.Restore(drawData).(*DrawingState)

	var (
		x = w.PosX + w.OffX + (toX / 64)
		y = w.PosY - w.OffY - (toY / 64)
	)

	if w.Vec != nil {
		w.Vec.MoveTo(x-w.MinX, y-w.MinY)
	} else {
		w.processY(y)
		w.processX(x)
	}
}

//export drawLineTo
func drawLineTo(drawData unsafe.Pointer, toX, toY float32) {
	w := pointer.Restore(drawData).(*DrawingState)

	var (
		x = w.PosX + w.OffX + (toX / 64)
		y = w.PosY - w.OffY - (toY / 64)
	)

	if w.Vec != nil {
		w.Vec.LineTo(x-w.MinX, y-w.MinY)
	} else {
		w.processY(y)
		w.processX(x)
	}
}

//export drawQuadraticTo
func drawQuadraticTo(drawData unsafe.Pointer, controlX, controlY, toX, toY float32) {
	w := pointer.Restore(drawData).(*DrawingState)

	var (
		x1 = w.PosX + w.OffX + (controlX / 64)
		y1 = w.PosY - w.OffY - (controlY / 64)
		x2 = w.PosX + w.OffX + (toX / 64)
		y2 = w.PosY - w.OffY - (toY / 64)
	)

	if w.Vec != nil {
		w.Vec.QuadTo(x1-w.MinX, y1-w.MinY, x2-w.MinX, y2-w.MinY)
	} else {
		w.processY(y1, y2)
		w.processX(x1, x2)
	}
}

//export drawCubicTo
func drawCubicTo(drawData unsafe.Pointer, control1X, control1Y, control2X, control2Y, toX, toY float32) {
	w := pointer.Restore(drawData).(*DrawingState)

	var (
		x1 = w.PosX + w.OffX + (control1X / 64)
		y1 = w.PosY - w.OffY - (control1Y / 64)
		x2 = w.PosX + w.OffX + (control2X / 64)
		y2 = w.PosY - w.OffY - (control2Y / 64)
		x3 = w.PosX + w.OffX + (toX / 64)
		y3 = w.PosY - w.OffY - (toY / 64)
	)

	if w.Vec != nil {
		w.Vec.CubeTo(x1-w.MinX, y1-w.MinY, x2-w.MinX, y2-w.MinY, x3-w.MinX, y3-w.MinY)
	} else {
		w.processY(y1, y2, y3)
		w.processX(x1, x2, x3)
	}
}

//export drawClosePath
func drawClosePath(drawData unsafe.Pointer) {
	vec := pointer.Restore(drawData).(*DrawingState).Vec
	if vec != nil {
		vec.ClosePath()
	}
}
