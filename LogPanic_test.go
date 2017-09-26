package debugo

import (
	"bytes"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogPanic(t *testing.T) {
	Convey("Given a panic", t, func() {
		buffer := bytes.NewBuffer(make([]byte, 0, 4096))
		log.SetOutput(buffer)
		func() {
			defer func() {
				recovered := recover()
				Convey("Then Recovered value should not be nil", func() {
					So(recovered, ShouldNotBeNil)
				})
			}()
			func() {
				defer LogPanic()
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
