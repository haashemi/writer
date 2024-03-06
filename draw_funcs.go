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

type drawingState struct {
	vec *vector.Rasterizer

	posX, posY float32
	offX, offY float32
	fontSize   float32
}

var drawFuncs = hb.DrawFuncsCreate()

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
	w.vec.MoveTo(w.posX+w.offX+float32(toX)/64, w.fontSize+w.posY-w.offY-(float32(toY)/64))
}

//export drawLineTo
func drawLineTo(drawData unsafe.Pointer, toX, toY C.float) {
	w := pointer.Restore(drawData).(*drawingState)
	w.vec.LineTo(w.posX+w.offX+float32(toX)/64, w.fontSize+w.posY-w.offY-(float32(toY)/64))
}

//export drawQuadraticTo
func drawQuadraticTo(drawData unsafe.Pointer, controlX, controlY, toX, toY C.float) {
	w := pointer.Restore(drawData).(*drawingState)
	w.vec.QuadTo(
		w.posX+w.offX+float32(controlX)/64, w.fontSize+w.posY-w.offY-(float32(controlY)/64),
		w.posX+w.offX+float32(toX)/64, w.fontSize+w.posY-w.offY-(float32(toY)/64),
	)
}

//export drawCubicTo
func drawCubicTo(drawData unsafe.Pointer, control1X, control1Y, control2X, control2Y, toX, toY C.float) {
	w := pointer.Restore(drawData).(*drawingState)
	w.vec.CubeTo(
		w.posX+w.offX+float32(control1X)/64, w.fontSize+w.posY-w.offY-(float32(control1Y)/64),
		w.posX+w.offX+float32(control2X)/64, w.fontSize+w.posY-w.offY-(float32(control2Y)/64),
		w.posX+w.offX+float32(toX)/64, w.fontSize+w.posY-w.offY-(float32(toY)/64),
	)
}

//export drawClosePath
func drawClosePath(drawData unsafe.Pointer) {
	pointer.Restore(drawData).(*drawingState).vec.ClosePath()
}
