package debugo

import (
	"errors"
	"runtime"
)

// GetCaller returns the caller Function's reference and the call's file and line number after the provided depth
func GetCaller(depth int) (stackCall *StackCall, err error) {
	pc, file, line, ok := runtime.Caller(depth + 2)
	if ok == false {
		err = errors.New("Failed to get function's Caller")
		return
	}
	stackCall = &StackCall{
		Func: runtime.FuncForPC(pc),
		File: file,
		Line: line,
	}
	return
}
