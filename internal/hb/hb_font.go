package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-t
type Font *C.hb_font_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-extents-t
type FontExtents struct {
	Ascender  int32 // The height of typographic ascenders.
	Descender int32 // The depth of typographic descenders.
	LineGap   int32 // The suggested line-spacing gap.
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-glyph-extents-t
type GlyphExtents struct {
	XBearing int32 // Distance from the x-origin to the left extremum of the glyph.
	YBearing int32 //Distance from the top extremum of the glyph to the y-origin.
	Width    int32 //Distance from the left extremum of the glyph to the right extremum.
	Height   int32 //Distance from the top extremum of the glyph to the bottom extremum.
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-add-glyph-origin-for-direction
func FontAddGlyphOriginForDirection(font Font, glyph uint32, direction Direction, x, y *int) {
	C.hb_font_add_glyph_origin_for_direction(font, C.uint(glyph), C.hb_direction_t(direction), (*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-create
func FontCreate(face Face) Font {
	return C.hb_font_create(face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-create-sub-font
func FontCreateSubFont(parent Font) Font {
	return C.hb_font_create_sub_font(parent)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-empty
func FontGetEmpty() Font {
	return C.hb_font_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-reference
func FontReference(font Font) Font {
	return C.hb_font_reference(font)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-destroy
func FontDestroy(font Font) {
	C.hb_font_destroy(font)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-set-user-data
func FontSetUserData(font Font, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_font_set_user_data(font, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-user-data
func FontGetUserData(font Font, key *UserDataKey) unsafe.Pointer {
	return C.hb_font_get_user_data(font, (*C.hb_user_data_key_t)(key))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-make-immutable
func FontMakeImmutable(font Font) {
	C.hb_font_make_immutable(font)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-is-immutable
func FontIsImmutable(font Font) bool {
	return C.hb_font_is_immutable(font) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-set-face
func FontSetFace(font Font, face Face) {
	C.hb_font_set_face(font, face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-face
func FontGetFace(font Font) Face {
	return C.hb_font_get_face(font)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph
func FontGetGlyph(font Font, unicode, variationSelector Codepoint) (glyph Codepoint, ok bool) {
	ok = C.hb_font_get_glyph(font, C.uint(unicode), C.uint(variationSelector), (*C.uint)(&glyph)) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-advance-for-direction
func FontGetGlyphAdvanceForDirection(font Font, glyph Codepoint, direction Direction) (x, y int32) {
	C.hb_font_get_glyph_advance_for_direction(font, C.uint(glyph), C.hb_direction_t(direction), (*C.int)(&x), (*C.int)(&y))
	return
}

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-advances-for-direction
// func FontGetGlyphAdvancesForDirection(font Font, glyph Codepoint, direction Direction, firstGlyph Codepoint, glyphStride uint32) (firstAdvance int32, advanceStride uint32) {
// 	C.hb_font_get_glyph_advances_for_direction(font, C.hb_direction_t(direction), (*C.uint)(&firstGlyph), C.uint(glyphStride), (*C.int)(&firstAdvance), (*C.uint)(&advanceStride))
// 	return
// }

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-contour-point
func FontGetGlyphContourPoint(font Font, glyph Codepoint, pointIndex uint32) (x, y int32, ok bool) {
	ok = C.hb_font_get_glyph_contour_point(font, C.uint(glyph), C.uint(pointIndex), (*C.int)(&x), (*C.int)(&y)) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-contour-point-for-origin
func FontGetGlyphContourPointForOrigin(font Font, glyph Codepoint, pointIndex uint32, direction Direction) (x, y int32, ok bool) {
	ok = C.hb_font_get_glyph_contour_point_for_origin(font, C.uint(glyph), C.uint(pointIndex), C.hb_direction_t(direction), (*C.int)(&x), (*C.int)(&y)) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-extents
func FontGetGlyphExtents(font Font, glyph Codepoint) (extents GlyphExtents, ok bool) {
	ok = C.hb_font_get_glyph_extents(font, C.uint(glyph), (*C.hb_glyph_extents_t)(unsafe.Pointer(&extents))) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-extents-for-origin
func FontGetGlyphExtentsForOrigin(font Font, glyph Codepoint, direction Direction) (extents GlyphExtents, ok bool) {
	ok = C.hb_font_get_glyph_extents_for_origin(font, C.uint(glyph), C.hb_direction_t(direction), (*C.hb_glyph_extents_t)(unsafe.Pointer(&extents))) == 1
	return
}

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-from-name
// func FontGetGlyphFromName() {
// 	return C.hb_font_get_glyph_from_name()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-h-advance
// func FontGetGlyphHAdvance() {
// 	return C.hb_font_get_glyph_h_advance()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-v-advance
// func FontGetGlyphVAdvance() {
// 	return C.hb_font_get_glyph_v_advance()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-h-advances
// func FontGetGlyphHAdvances() {
// 	return C.hb_font_get_glyph_h_advances()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-v-advances
// func FontGetGlyphVAdvances() {
// 	return C.hb_font_get_glyph_v_advances()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-h-kerning
// func FontGetGlyphHKerning() {
// 	return C.hb_font_get_glyph_h_kerning()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-kerning-for-direction
// func FontGetGlyphKerningForDirection() {
// 	return C.hb_font_get_glyph_kerning_for_direction()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-h-origin
// func FontGetGlyphHOrigin() {
// 	return C.hb_font_get_glyph_h_origin()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-v-origin
// func FontGetGlyphVOrigin() {
// 	return C.hb_font_get_glyph_v_origin()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-origin-for-direction
// func FontGetGlyphOriginForDirection() {
// 	return C.hb_font_get_glyph_origin_for_direction()
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-glyph-name
// func FontGetGlyphName(font Font, glyph Codepoint) {
// 	return C.hb_font_get_glyph_name()
// }

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-draw-glyph
func FontDrawGlyph(font Font, glyph Codepoint, dfuncs DrawFuncs, drawData unsafe.Pointer) {
	C.hb_font_draw_glyph(font, C.uint(glyph), dfuncs, drawData)
}

// TODO:
// Learn more:
// func hb_font_paint_glyph(font Font) {
// 	return C.hb_font_paint_glyph(font)
// }

// TODO:
// Learn more:
// func hb_font_get_nominal_glyph(font Font) {
// 	return C.hb_font_get_nominal_glyph(font)
// }

// TODO:
// Learn more:
// func hb_font_get_nominal_glyphs(font Font) {
// 	return C.hb_font_get_nominal_glyphs(font)
// }

// TODO:
// Learn more:
// func hb_font_get_variation_glyph(font Font) {
// 	return C.hb_font_get_variation_glyph(font)
// }

// TODO:
// Learn more:
// func hb_font_set_parent(font Font) {
// 	return C.hb_font_set_parent(font)
// }

// TODO:
// Learn more:
// func hb_font_get_parent(font Font) {
// 	return C.hb_font_get_parent(font)
// }

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-set-ppem
func FontSetPpem(font Font, x, y uint32) {
	C.hb_font_set_ppem(font, C.uint(x), C.uint(y))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-ppem
func FontGetPpem(font Font) (x, y uint32) {
	C.hb_font_get_ppem(font, (*C.uint)(&x), (*C.uint)(&y))
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-set-ptem
func FontSetPtem(font Font, ptem float32) {
	C.hb_font_set_ptem(font, C.float(ptem))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-ptem
func FontGetPtem(font Font) float32 {
	return float32(C.hb_font_get_ptem(font))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-set-scale
func FontSetScale(font Font, x, y int32) {
	C.hb_font_set_scale(font, C.int(x), C.int(y))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-scale
func FontGetScale(font Font) (x, y int32) {
	C.hb_font_get_scale(font, (*C.int)(&x), (*C.int)(&y))
	return
}

// TODO: hb_font_get_synthetic_bold
// TODO: hb_font_set_synthetic_bold
// TODO: hb_font_set_synthetic_slant
// TODO: hb_font_get_synthetic_slant
// TODO: hb_font_set_variations
// TODO: hb_font_set_variation
// TODO: hb_font_set_var_named_instance
// TODO: hb_font_get_var_named_instance
// TODO: hb_font_set_var_coords_design
// TODO: hb_font_get_var_coords_design
// TODO: hb_font_set_var_coords_normalized
// TODO: hb_font_get_var_coords_normalized
// TODO: hb_font_glyph_from_string
// TODO: hb_font_glyph_to_string
// TODO: hb_font_get_serial
// TODO: hb_font_changed
// TODO: hb_font_set_funcs
// TODO: hb_font_set_funcs_data
// TODO: hb_font_subtract_glyph_origin_for_direction
// TODO: hb_font_funcs_create
// TODO: hb_font_funcs_get_empty
// TODO: hb_font_funcs_reference
// TODO: hb_font_funcs_destroy
// TODO: hb_font_funcs_set_user_data
// TODO: hb_font_funcs_get_user_data
// TODO: hb_font_funcs_make_immutable
// TODO: hb_font_funcs_is_immutable

// TODO: (*hb_font_get_glyph_contour_point_func_t)
// TODO: hb_font_funcs_set_glyph_contour_point_func
// TODO: (*hb_font_get_glyph_extents_func_t)
// TODO: hb_font_funcs_set_glyph_extents_func
// TODO: (*hb_font_get_glyph_from_name_func_t)
// TODO: hb_font_funcs_set_glyph_from_name_func
// TODO: (*hb_font_get_glyph_advance_func_t)
// TODO: hb_font_funcs_set_glyph_h_advance_func
// TODO: hb_font_funcs_set_glyph_v_advance_func
// TODO: (*hb_font_get_glyph_advances_func_t)
// TODO: hb_font_funcs_set_glyph_h_advances_func
// TODO: hb_font_funcs_set_glyph_v_advances_func
// TODO: (*hb_font_get_glyph_kerning_func_t)
// TODO: hb_font_funcs_set_glyph_h_kerning_func
// TODO: (*hb_font_get_glyph_origin_func_t)
// TODO: hb_font_funcs_set_glyph_h_origin_func
// TODO: hb_font_funcs_set_glyph_v_origin_func
// TODO: (*hb_font_get_glyph_name_func_t)
// TODO: hb_font_funcs_set_glyph_name_func
// TODO: (*hb_font_draw_glyph_func_t)
// TODO: hb_font_funcs_set_draw_glyph_func
// TODO: (*hb_font_paint_glyph_func_t)
// TODO: hb_font_funcs_set_paint_glyph_func
// TODO: (*hb_font_get_nominal_glyph_func_t)
// TODO: hb_font_funcs_set_nominal_glyph_func
// TODO: (*hb_font_get_nominal_glyphs_func_t)
// TODO: hb_font_funcs_set_nominal_glyphs_func
// TODO: (*hb_font_get_variation_glyph_func_t)
// TODO: hb_font_funcs_set_variation_glyph_func

// A callback function for FaceCreateForTables.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-reference-table-func-t
type ReferenceTableFunc C.hb_reference_table_func_t

// TODO: (*hb_font_get_font_extents_func_t)
// TODO: hb_font_funcs_set_font_h_extents_func
// TODO: hb_font_funcs_set_font_v_extents_func

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-h-extents
func FontGetHExtents(font Font) (extents FontExtents, ok bool) {
	ok = C.hb_font_get_h_extents(font, (*C.hb_font_extents_t)(unsafe.Pointer(&extents))) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-v-extents
func FontGetVExtents(font Font) (extents FontExtents, ok bool) {
	ok = C.hb_font_get_v_extents(font, (*C.hb_font_extents_t)(unsafe.Pointer(&extents))) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-extents-for-direction
func FontGetExtentsForDirection(font Font, direction Direction) (extents FontExtents) {
	C.hb_font_get_extents_for_direction(font, C.hb_direction_t(direction), (*C.hb_font_extents_t)(unsafe.Pointer(&extents)))
	return
}
