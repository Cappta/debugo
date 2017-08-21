package debugo

import "runtime"

// StackCall represents a stack call unit
type StackCall struct {
	Func *runtime.Func
	File string
	Line int
}
