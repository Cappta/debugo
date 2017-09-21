package debugo

import (
	"fmt"
	"io"

	"github.com/go-errors/errors"
)

// LogPanicInto will recover from a panic, log with it's callstack into the writer and panic again
func LogPanicInto(w io.Writer) {
	if recovered := recover(); recovered != nil {
		fmt.Fprintln(w, errors.Wrap(recovered, 0).ErrorStack())
		panic(recovered)
	}
}
