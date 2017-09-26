package debugo

import (
	"bytes"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogPanicInto(t *testing.T) {
	Convey("Given a panic", t, func() {
		buffer := bytes.NewBuffer(make([]byte, 0, 4096))
		func() {
			defer func() {
				recovered := recover()
				Convey("Then Recovered value should not be nil", func() {
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
		Convey("Then buffer's string value should contain substring 'AFunctionThatWillPanic'", func() {
			So(buffer.String(), ShouldContainSubstring, "AFunctionThatWillPanic")
		})
		Convey("Then buffer's string value should contain substring 'You were warned'", func() {
			So(buffer.String(), ShouldContainSubstring, "You were warned")
		})
	})
}

func AFunctionThatWillPanic() {
	panic(errors.New("You were warned"))
}
