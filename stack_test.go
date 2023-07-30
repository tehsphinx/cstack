package cstack

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkStack_DefaultFormat_StringPlus(b *testing.B) {
	cs := CallStack(0)
	for i := 0; i < b.N; i++ {
		cs.Format(func(frame Frame) string {
			file, line := frame.FileLine()
			fName := frame.FuncName()

			return "  " + fName + "\n    " + file + ":" + strconv.Itoa(line) + "\n"
		})
	}
}

func BenchmarkStack_DefaultFormat_Sprintf(b *testing.B) {
	cs := CallStack(0)
	for i := 0; i < b.N; i++ {
		cs.Format(func(frame Frame) string {
			file, line := frame.FileLine()
			fName := frame.FuncName()

			return fmt.Sprintf("  %s\n    %s:%d\n", fName, file, line)
		})
	}
}

func BenchmarkStack_DefaultFormat_StringBuilder(b *testing.B) {
	cs := CallStack(0)
	for i := 0; i < b.N; i++ {
		cs.Format(func(frame Frame) string {
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
}
