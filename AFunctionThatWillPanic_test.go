package debugo

import "errors"

func AFunctionThatWillPanic() {
	panic(errors.New("You were warned"))
}
