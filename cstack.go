package cstack

import (
	"runtime"
)

func CallStack(skip int) Stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	return pcs[0:n]
}
