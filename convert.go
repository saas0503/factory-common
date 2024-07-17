package common

import "unsafe"

// UnsafeString returns a string pointer without allocation
func UnsafeString(b []byte) string {
	// the new way is slower `return unsafe.String(unsafe.SliceData(b), len(b))`
	// unsafe.Pointer variant: 0.3538 ns/op vs unsafe.String variant: 0.5410 ns/op
	// #nosec G103
	return *(*string)(unsafe.Pointer(&b))
}
