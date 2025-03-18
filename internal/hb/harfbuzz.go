package hb

// #cgo pkg-config: harfbuzz
// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

func cBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func cFeatures(features []Feature) *C.hb_feature_t {
	if len(features) == 0 {
		return nil
	}

	return (*C.hb_feature_t)(unsafe.Pointer(&features[0]))
}

func cStringArray(items []string) []*C.char {
	if items == nil {
		return nil
	}

	arr := make([]*C.char, len(items)+1)

	for i, item := range items {
		arr[i] = C.CString(item)
	}
	arr[len(items)] = nil

	return arr
}

func freeStringArray(items []*C.char) {
	for _, item := range items {
		C.free(unsafe.Pointer(item))
	}
}
