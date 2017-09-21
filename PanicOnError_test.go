package debugo

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPanicOnError(t *testing.T) {
	Convey("Given an error", t, func() {
		err := errors.New("An error")
		Convey("Then PanicOnError should panic", func() {
			So(func() { PanicOnError(err) }, ShouldPanic)
		})
	})
	Convey("Given a nil error", t, func() {
		var err error
		Convey("Then PanicOnError should not panic", func() {
			So(func() { PanicOnError(err) }, ShouldNotPanic)
		})
	})
}
