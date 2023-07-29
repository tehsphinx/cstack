# cstack

CStack implements a light-weight wrapper around the call stack returned by `runtime.Callers`.
The wrapper add convenience functions to get information on the stack.

## Usage

```go
package main

import (
	"fmt"

	"github.com/tehsphinx/cstack"
)

func main() {
	// get stack skipping 0 layers
	stack := cstack.CallStack(0)
	
	// Use Format to transform each frame to a string 
	stackTrace := stack.Format(func(frame cstack.Frame) string {
		fName := frame.FuncName()
		file, line := frame.FileLine()
		return fmt.Sprintf("  %s\n    %s:%d\n", fName, file, line)
	})
	
	fmt.Print(stackTrace)
}
```
