package cstack

import (
	"strconv"
	"strings"
)

type Stack []uintptr

func (s Stack) Format(f func(Frame) string) string {
	var b strings.Builder
	for _, ptr := range []uintptr(s) {
		b.WriteString(f(Frame(ptr)))
	}
	return b.String()
}

func (s Stack) DefaultFormat() string {
	return s.Format(func(frame Frame) string {
		file, line := frame.FileLine()
		fName := frame.FuncName()
		// TODO: performance test
		return "  " + fName + "\n    " + file + ":" + strconv.Itoa(line) + "\n"
	})
}

func (s Stack) StackInfo() []FrameInfo {
	frames := make([]FrameInfo, 0, len(s))
	for _, f := range s {
		frame := Frame(f)
		file, line := frame.FileLine()

		frames = append(frames, FrameInfo{
			File: file,
			Line: line,
			Func: frame.FuncName(),
		})
	}
	return frames
}
