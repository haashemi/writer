package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape.html#hb-shape
func Shape(font Font, buffer Buffer, features []Feature) {
	C.hb_shape(font, buffer, cFeatures(features), C.uint(len(features)))
}
