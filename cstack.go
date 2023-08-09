// Package cstack implements some helper functionality around the call stack.
package cstack

import (
	"runtime"
)

// CallStack gets the Stack by taking the current call stack and skipping given amount of frames.
func CallStack(skip int) Stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	return pcs[0:n]
}
