package hb

// #include <hb.h>
import "C"
import "unsafe"

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-t
type DrawFuncs *C.hb_draw_funcs_t

// DrawState holds the current drawing state.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-state-t
type DrawState struct {
	PathOpen bool // Whether there is an open path
	// in CGO, path_open is an int32, but it's a boolean. so we skip the next
	// three bytes to be able to use a boolean for our bindings and also make
	// sure that the C and Go structs are compatible with each other
	skipped    [3]byte     // [private] reserved by go-hb.
	PathStartX float32     // X component of the start of current path
	PathStartY float32     // Y component of the start of current path
	CurrentX   float32     // X component of current point
	CurrentY   float32     // Y component of current point
	reserved   [7 * 4]byte // [private] reserved by hb.
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-create
func DrawFuncsCreate() DrawFuncs {
	return C.hb_draw_funcs_create()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-get-empty
func DrawFuncsGetEmpty() DrawFuncs {
	return C.hb_draw_funcs_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-reference
func DrawFuncsReference(dfuncs DrawFuncs) DrawFuncs {
	return C.hb_draw_funcs_reference(dfuncs)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-destroy
func DrawFuncsDestroy(dfuncs DrawFuncs) {
	C.hb_draw_funcs_destroy(dfuncs)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-set-user-data
func DrawFuncsSetUserData(dfuncs DrawFuncs, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_draw_funcs_set_user_data(dfuncs, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-get-user-data
func DrawFuncsGetUserData(dfuncs DrawFuncs, key *UserDataKey) unsafe.Pointer {
	return C.hb_draw_funcs_get_user_data(dfuncs, (*C.hb_user_data_key_t)(key))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-make-immutable
func DrawFuncsMakeImmutable(dfuncs DrawFuncs) {
	C.hb_draw_funcs_make_immutable(dfuncs)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-is-immutable
func DrawFuncsIsImmutable(dfuncs DrawFuncs) bool {
	return C.hb_draw_funcs_is_immutable(dfuncs) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-move-to-func-t
type DrawMoveToFunc C.hb_draw_move_to_func_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-set-move-to-func
func DrawFuncsSetMoveToFunc(dfuncs DrawFuncs, fn DrawMoveToFunc, userData unsafe.Pointer, destroy DestroyFunc) {
	C.hb_draw_funcs_set_move_to_func(dfuncs, fn, userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-line-to-func-t
type DrawLineToFunc C.hb_draw_line_to_func_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-set-line-to-func
func DrawFuncsSetLineToFunc(dfuncs DrawFuncs, fn DrawLineToFunc, userData unsafe.Pointer, destroy DestroyFunc) {
	C.hb_draw_funcs_set_line_to_func(dfuncs, fn, userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-quadratic-to-func-t
type DrawQuadraticToFunc C.hb_draw_quadratic_to_func_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-set-quadratic-to-func
func DrawFuncsSetQuadraticToFunc(dfuncs DrawFuncs, fn DrawQuadraticToFunc, userData unsafe.Pointer, destroy DestroyFunc) {
	C.hb_draw_funcs_set_quadratic_to_func(dfuncs, fn, userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-cubic-to-func-t
type DrawCubicToFunc C.hb_draw_cubic_to_func_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-set-cubic-to-func
func DrawFuncsSetCubicToFunc(dfuncs DrawFuncs, fn DrawCubicToFunc, userData unsafe.Pointer, destroy DestroyFunc) {
	C.hb_draw_funcs_set_cubic_to_func(dfuncs, fn, userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-close-path-func-t
type DrawClosePathFunc C.hb_draw_close_path_func_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-funcs-set-close-path-func
func DrawFuncsSetClosePathFunc(dfuncs DrawFuncs, fn DrawClosePathFunc, userData unsafe.Pointer, destroy DestroyFunc) {
	C.hb_draw_funcs_set_close_path_func(dfuncs, fn, userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-move-to
func DrawMoveTo(dfuncs DrawFuncs, drawData unsafe.Pointer, st *DrawState, toX, toY float32) {
	C.hb_draw_move_to(dfuncs, drawData, (*C.hb_draw_state_t)(unsafe.Pointer(st)), C.float(toX), C.float(toY))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-line-to
func DrawLineTo(dfuncs DrawFuncs, drawData unsafe.Pointer, st *DrawState, toX, toY float32) {
	C.hb_draw_line_to(dfuncs, drawData, (*C.hb_draw_state_t)(unsafe.Pointer(st)), C.float(toX), C.float(toY))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-quadratic-to
func DrawQuadraticTo(dfuncs DrawFuncs, drawData unsafe.Pointer, st *DrawState, controlX, controlY, toX, toY float32) {
	C.hb_draw_quadratic_to(dfuncs, drawData, (*C.hb_draw_state_t)(unsafe.Pointer(st)), C.float(controlX), C.float(controlY), C.float(toX), C.float(toY))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-cubic-to
func DrawCubicTo(dfuncs DrawFuncs, drawData unsafe.Pointer, st *DrawState, control1X, control1Y, control2X, control2Y, toX, toY float32) {
	C.hb_draw_cubic_to(dfuncs, drawData, (*C.hb_draw_state_t)(unsafe.Pointer(st)), C.float(control1X), C.float(control1Y), C.float(control2X), C.float(control2Y), C.float(toX), C.float(toY))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-draw.html#hb-draw-close-path
func DrawClosePath(dfuncs DrawFuncs, drawData unsafe.Pointer, st *DrawState) {
	C.hb_draw_close_path(dfuncs, drawData, (*C.hb_draw_state_t)(unsafe.Pointer(st)))
}
