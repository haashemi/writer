package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

// Face holds font faces.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-t
type Face *C.hb_face_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-count
func FaceCount(blob Blob) uint32 {
	return uint32(C.hb_face_count(blob))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-create
func FaceCreate(blob Blob, index uint32) Face {
	return C.hb_face_create(blob, C.uint(index))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-create-for-tables
func FaceCreateForTables(referenceTableFunc ReferenceTableFunc, userData unsafe.Pointer, destroy DestroyFunc) Face {
	return C.hb_face_create_for_tables(referenceTableFunc, userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-get-empty
func FaceGetEmpty() Face {
	return C.hb_face_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-reference
func FaceReference(face Face) Face {
	return C.hb_face_reference(face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-destroy
func FaceDestroy(face Face) {
	C.hb_face_destroy(face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-set-user-data
func FaceSetUserData(face Face, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_face_set_user_data(face, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-get-user-data
func FaceGetUserData(face Face, key *UserDataKey) unsafe.Pointer {
	return C.hb_face_get_user_data(face, (*C.hb_user_data_key_t)(key))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-make-immutable
func FaceMakeImmutable(face Face) {
	C.hb_face_make_immutable(face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-is-immutable
func FaceIsImmutable(face Face) bool {
	return C.hb_face_is_immutable(face) == 1
}

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-get-table-tags
// func FaceGetTableTags(face Face, startOffset uint32) (uint32) {
// 	return uint32(C.hb_face_get_table_tags(face, C.uint(startOffset)))
// }

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-set-glyph-count
func FaceSetGlyphCount(face Face, glyphCount uint32) {
	C.hb_face_set_glyph_count(face, C.uint(glyphCount))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-get-glyph-count
func FaceGetGlyphCount(face Face) uint32 {
	return uint32(C.hb_face_get_glyph_count(face))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-set-index
func FaceSetIndex(face Face, index uint32) {
	C.hb_face_set_index(face, C.uint(index))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-get-index
func FaceGetIndex(face Face) uint32 {
	return uint32(C.hb_face_get_index(face))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-set-upem
func FaceSetUpem(face Face, upem uint32) {
	C.hb_face_set_upem(face, C.uint(upem))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-get-upem
func FaceGetUpem(face Face) uint32 {
	return uint32(C.hb_face_get_upem(face))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-reference-blob
func FaceReferenceBlob(face Face) Blob {
	return C.hb_face_reference_blob(face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-reference-table
func FaceReferenceTable(face Face, tag Tag) Blob {
	return C.hb_face_reference_table(face, *(*C.hb_tag_t)(unsafe.Pointer(&tag[0])))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-collect-unicodes
func FaceCollectUnicodes(face Face, out Set) {
	C.hb_face_collect_unicodes(face, out)
}

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-collect-nominal-glyph-mapping
// func FaceCollectNominalGlyphMapping(face Face) {
// 	C.hb_face_collect_nominal_glyph_mapping()
// }

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-collect-variation-selectors
func FaceCollectVariationSelectors(face Face, out Set) {
	C.hb_face_collect_variation_selectors(face, out)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-collect-variation-unicodes
func FaceCollectVariationUnicodes(face Face, variationSelector Codepoint, out Set) {
	C.hb_face_collect_variation_unicodes(face, C.uint(variationSelector), out)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-builder-create
func FaceBuilderCreate() Face {
	return C.hb_face_builder_create()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-builder-add-table
func FaceBuilderAddTable(face Face, tag Tag, blob Blob) bool {
	return C.hb_face_builder_add_table(face, *(*C.hb_tag_t)(unsafe.Pointer(&tag[0])), blob) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-builder-sort-tables
func FaceBuilderSortTables(face Face, tags []Tag) {
	var cTags []*C.hb_tag_t
	for _, tag := range tags {
		cTags = append(cTags, (*C.hb_tag_t)(unsafe.Pointer(&tag[0])))
	}
	cTags = append(cTags, nil)

	C.hb_face_builder_sort_tables(face, cTags[0])
}
