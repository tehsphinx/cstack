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
		lineStr := strconv.Itoa(line)

		var builder strings.Builder
		builder.Grow(9 + len(fName) + len(file) + len(lineStr))
		builder.WriteString("  ")
		builder.WriteString(fName)
		builder.WriteString("\n    ")
		builder.WriteString(file)
		builder.WriteByte(':')
		builder.WriteString(lineStr)

		return builder.String()
	})
}

func (s Stack) StackInfo() []FrameInfo {
	frames := make([]FrameInfo, 0, len(s))
	for _, f := range s {
		frames = append(frames, Frame(f).FrameInfo())
	}
	return frames
}
