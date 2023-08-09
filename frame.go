package cstack

import "runtime"

// Frame defines a stack frame.
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

// FrameInfo defines a stack frames information struct.
type FrameInfo struct {
	Func string `json:"func"`
	File string `json:"file"`
	Line int    `json:"line"`
}

// FrameInfo extracts the FrameInfo from a Frame.
func (f Frame) FrameInfo() FrameInfo {
	file, line := f.FileLine()

	return FrameInfo{
		File: file,
		Line: line,
		Func: f.FuncName(),
	}
}
