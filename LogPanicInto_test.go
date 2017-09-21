package debugo

import (
	"bytes"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogPanicInto(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 4096))
	func() {
		defer func() {
			recovered := recover()
			Convey("Then Recovered value should not be nil", t, func() {
				So(recovered, ShouldNotBeNil)
			})
		}()
		func() {
			defer LogPanicInto(buffer)
			func() {
				AFunctionThatWillPanic()
			}()
		}()
	}()
	Convey("Then buffer's string value should contain substring 'AFunctionThatWillPanic'", t, func() {
		So(buffer.String(), ShouldContainSubstring, "AFunctionThatWillPanic")
	})
	Convey("Then buffer's string value should contain substring 'You were warned'", t, func() {
		So(buffer.String(), ShouldContainSubstring, "You were warned")
	})
}

func AFunctionThatWillPanic() {
	panic(errors.New("You were warned"))
}
