package hb

// #include <hb.h>
import "C"
import "unsafe"

// SetValueInvalid is an unset Set value.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#HB-SET-VALUE-INVALID:CAPS
const SetValueInvalid = C.HB_SET_VALUE_INVALID

// Set holds a set of integers. Set's are used to gather and contain glyph IDs,
// Unicode code points, and various other collections of discrete values.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-t
type Set *C.hb_set_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-create
func SetCreate() Set {
	return C.hb_set_create()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-allocation-successful
func SetAllocationSuccessful(set Set) bool {
	return C.hb_set_allocation_successful(set) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-copy
func SetCopy(set Set) Set {
	return C.hb_set_copy(set)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-get-empty
func SetGetEmpty() Set {
	return C.hb_set_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-reference
func SetReference(set Set) Set {
	return C.hb_set_reference(set)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-destroy
func SetDestroy(set Set) {
	C.hb_set_destroy(set)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-set-user-data
func SetSetUserData(set Set, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_set_set_user_data(set, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-get-user-data
func SetGetUserData(set Set, key *UserDataKey) unsafe.Pointer {
	return C.hb_set_get_user_data(set, (*C.hb_user_data_key_t)(key))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-clear
func SetClear(set Set) {
	C.hb_set_clear(set)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-set
func SetSet(set, other Set) {
	C.hb_set_set(set, other)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-has
func SetHas(set Set, codePoint Codepoint) bool {
	return C.hb_set_has(set, C.uint(codePoint)) == 1
}

// SetAdd adds codePoint to set.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-add
func SetAdd(set Set, codePoint Codepoint) {
	C.hb_set_add(set, C.uint(codePoint))
}

// SetAddRange adds all of the elements from first to last (inclusive) to set.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-add-range
func SetAddRange(set Set, first, last Codepoint) {
	C.hb_set_add_range(set, C.uint(first), C.uint(last))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-add-sorted-array
func SetAddSortedArray(set Set, sortedCodepoints []Codepoint) {
	C.hb_set_add_sorted_array(set, (*C.uint)(unsafe.Pointer(&sortedCodepoints[0])), C.uint(len(sortedCodepoints)))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-del
func SetDel(set Set, codepoint Codepoint) {
	C.hb_set_del(set, C.uint(codepoint))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-del-range
func SetDelRange(set Set, first, last Codepoint) {
	C.hb_set_del_range(set, C.uint(first), C.uint(last))
}

// SetGetMax returns the largest element in the set.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-get-max
func SetGetMax(set Set) uint32 {
	return uint32(C.hb_set_get_max(set))
}

// SetGetMin returns the smallest element in the set.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-get-min
func SetGetMin(set Set) uint32 {
	return uint32(C.hb_set_get_min(set))
}

// SetGetPopulation returns the number of elements in the set.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-get-population
func SetGetPopulation(set Set) uint32 {
	return uint32(C.hb_set_get_population(set))
}

// SetIsEmpty tests whether a set is empty (contains no elements).
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-is-empty
func SetIsEmpty(set Set) bool {
	return C.hb_set_is_empty(set) == 1
}

// SetHash Creates and returns a hash representing set .
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-hash
func SetHash(set Set) uint32 {
	return uint32(C.hb_set_hash(set))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-subtract
func SetSubtract(set, other Set) {
	C.hb_set_subtract(set, other)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-intersect
func SetIntersect(set, other Set) {
	C.hb_set_intersect(set, other)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-union
func SetUnion(set, other Set) {
	C.hb_set_union(set, other)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-symmetric-difference
func SetSymmetricDifference(set, other Set) {
	C.hb_set_symmetric_difference(set, other)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-invert
func SetInvert(set Set) {
	C.hb_set_invert(set)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-is-inverted
func SetIsInverted(set Set) bool {
	return C.hb_set_is_inverted(set) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-is-equal
func SetIsEqual(set, other Set) bool {
	return C.hb_set_is_equal(set, other) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-is-subset
func SetIsSubset(set, other Set) bool {
	return C.hb_set_is_subset(set, other) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-next
func SetNext(set Set, codepoint Codepoint) (Codepoint, bool) {
	ok := C.hb_set_next(set, (*C.uint)(&codepoint)) == 1
	return codepoint, ok
}

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-next-range
// func SetNextRange(set Set) bool {
// 	return C.hb_set_next_range(set) == 1
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-next-many
// func SetNextMany(set Set) uint32 {
// 	return uint32(C.hb_set_next_many(set))
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-previous
// func SetPrevious(set Set) bool {
// 	return C.hb_set_previous(set) == 1
// }

// TODO:
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-set.html#hb-set-previous-range
// func SetPreviousRange(set Set) bool {
// 	return C.hb_set_previous_range(set) == 1
// }
