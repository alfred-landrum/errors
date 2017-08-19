// +build go1.9

package errors

import (
	"runtime"
)

// Frame holds call frame information from a stack trace.
type Frame struct {
	rf runtime.Frame
}

func (f Frame) runtimeFunc() *runtime.Func {
	return f.rf.Func
}

func (f Frame) function() string {
	return f.rf.Function
}

func (f Frame) file() string {
	return f.rf.File
}

func (f Frame) line() int {
	return f.rf.Line
}

func (s *stack) StackTrace() StackTrace {
	cframes := runtime.CallersFrames(*s)

	frames := make([]Frame, 0, len(*s))
	for {
		f, more := cframes.Next()
		frames = append(frames, Frame{f})
		if !more {
			break
		}
	}

	return frames
}
