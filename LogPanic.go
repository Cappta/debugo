package debugo

import (
	"log"

	"github.com/go-errors/errors"
)

// LogPanic will recover from a panic, log with it's callstack and panic again
func LogPanic() {
	if recovered := recover(); recovered != nil {
		log.Println(errors.Wrap(recovered, 0).ErrorStack())
		panic(recovered)
	}
}
