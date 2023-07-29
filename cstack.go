package cstack

import (
	"runtime"
	"strings"
)

func CallStack(skip int) Stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	return pcs[0:n]
}

type Stack []uintptr

func (s Stack) Format(f func(Frame) string) string {
	var b strings.Builder
	for _, ptr := range []uintptr(s) {
		b.WriteString(f(Frame(ptr)))
	}
	return b.String()
}

type Frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f Frame) pc() uintptr { return uintptr(f) - 1 }

// FileLine returns the full path to the file that contains the
// function for this Frame's pc as well as the line number.
func (f Frame) FileLine() (string, int) {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown", 0
	}
	file, line := fn.FileLine(f.pc())
	return file, line
}

// FuncName returns the name of the function, if known.
func (f Frame) FuncName() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}
