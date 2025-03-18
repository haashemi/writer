package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

// Blob wraps a chunk of binary data and facilitates its lifecycle management
// between a client program and hb.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-t
type Blob *C.hb_blob_t

// MemoryMode holds the memory modes available to client programs.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-memory-mode-t
type MemoryMode C.hb_memory_mode_t

const (
	// HarfBuzz immediately makes a copy of the data.
	//
	// In no case shall the HarfBuzz client modify memory that is passed to
	// HarfBuzz in a blob. If there is any such possibility, this mode should be
	// used such that HarfBuzz makes a copy immediately,
	MemoryModeDuplicate MemoryMode = C.HB_MEMORY_MODE_DUPLICATE

	// HarfBuzz client will never modify the data, and HarfBuzz will never modify
	// the data.
	//
	// Use this if it's ok for Harfbuzz client to modify memory that is passed
	// too Harfbuzz in a blob, unless you really really really know what you are
	// doing.
	MemoryModeReadonly MemoryMode = C.HB_MEMORY_MODE_READONLY

	// HarfBuzz client made a copy f the data solely for HarfBuzz, so HarfBuzz
	// may modify the data.
	//
	// This mode is appropriate if you really made a copy of data solely for the
	// purpose of passing to HarfBuzz and doing that just once (no reuse!).
	MemoryModeWritable MemoryMode = C.HB_MEMORY_MODE_WRITABLE

	MemoryModeReadonlyMayMakeWritable MemoryMode = C.HB_MEMORY_MODE_READONLY_MAY_MAKE_WRITABLE
)

// BlobCreate creates a new Blob wrapping data. The mode parameter is used to
// negotiate ownership and lifecycle of data.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create
func BlobCreate(data []byte, mode MemoryMode, userData unsafe.Pointer, destroy DestroyFunc) Blob {
	return C.hb_blob_create((*C.char)(unsafe.Pointer(&data[0])), C.uint(len(data)), C.hb_memory_mode_t(mode), userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-or-fail
func BlobCreateOrFail(data []byte, mode MemoryMode, userData unsafe.Pointer, destroy DestroyFunc) Blob {
	return C.hb_blob_create_or_fail((*C.char)(unsafe.Pointer(&data[0])), C.uint(len(data)), C.hb_memory_mode_t(mode), userData, destroy)
}

// BlobCreateFromFile creates a new blob containing the data from the specified
// binary font file.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-from-file
func BlobCreateFromFile(filename string) Blob {
	file_name := C.CString(filename)
	defer C.free(unsafe.Pointer(file_name))

	return C.hb_blob_create_from_file(file_name)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-from-file-or-fail
func BlobCreateFromFileOrFail(filename string) Blob {
	file_name := C.CString(filename)
	defer C.free(unsafe.Pointer(file_name))

	return C.hb_blob_create_from_file_or_fail(file_name)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-sub-blob
func BlobCreateSubBlob(parent Blob, offset, length uint32) Blob {
	return C.hb_blob_create_sub_blob(parent, C.uint(offset), C.uint(length))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-copy-writable-or-fail
func BlobCopyWritableOrFail(blob Blob) Blob {
	return C.hb_blob_copy_writable_or_fail(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-empty
func BlobGetEmpty() Blob {
	return C.hb_blob_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-reference
func BlobReference(blob Blob) Blob {
	return C.hb_blob_reference(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-destroy
func BlobDestroy(blob Blob) {
	C.hb_blob_destroy(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-set-user-data
func BlobSetUserData(blob Blob, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_blob_set_user_data(blob, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-user-data
func BlobGetUserData(blob Blob, key *UserDataKey) unsafe.Pointer {
	return C.hb_blob_get_user_data(blob, (*C.hb_user_data_key_t)(key))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-make-immutable
func BlobMakeImmutable(blob Blob) {
	C.hb_blob_make_immutable(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-is-immutable
func BlobIsImmutable(blob Blob) bool {
	return C.hb_blob_is_immutable(blob) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-data
func BlobGetData(blob Blob) string {
	var length C.uint
	data := C.hb_blob_get_data(blob, &length)

	return C.GoStringN(data, C.int(length))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-length
func BlobGetLength(blob Blob) uint32 {
	return uint32(C.hb_blob_get_length(blob))
}
