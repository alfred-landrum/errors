// +build !go1.9

package errors

import (
	"runtime"
)

// Frame represents a program counter inside a stack frame.
type Frame struct {
	pcp uintptr
}

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f Frame) pc() uintptr { return f.pcp - 1 }

func (f Frame) runtimeFunc() *runtime.Func {
	return runtime.FuncForPC(f.pc())
}

func (f Frame) function() string {
	return f.runtimeFunc().Name()
}

// file returns the full path to the file that contains the
// function for this Frame's pc.
func (f Frame) file() string {
	fn := f.runtimeFunc()
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

// line returns the line number of source code of the
// function for this Frame's pc.
func (f Frame) line() int {
	fn := f.runtimeFunc()
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

func (s *stack) StackTrace() StackTrace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame{(*s)[i]}
	}
	return f
}
