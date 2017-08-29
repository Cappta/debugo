package debugo

import (
	"testing"

	. "github.com/Cappta/gohelpconvey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCaller(t *testing.T) {
	CalledFunc(t)
}

func CalledFunc(t *testing.T) {
	stackCall, err := GetCaller(0)
	Convey("When Getting Caller from CalledFunc", t, func() {
		Convey("Given the callerFunc reference", func() {
			Convey("Then Caller's Name should equal TestCaller", func() {
				So(stackCall.Func.Name(), ShouldEqual, "github.com/Cappta/debugo.TestCaller")
			})
		})
		Convey("Then File should match current file", func() {
			So(stackCall.File, ShouldMatch, ".*/github.com/Cappta/debugo/caller_test.go")
		})
		Convey("Then Line should equal CalledFunc's call line", func() {
			So(stackCall.Line, ShouldEqual, 11)
		})
		Convey("Then error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
	stackCall, err = GetCaller(1)
	Convey("When Getting Caller After Caller Method from CalledFunc", t, func() {
		Convey("Given the callerFunc reference", func() {
			Convey("Then Caller's Name should equal testing.tRunner", func() {
				So(stackCall.Func.Name(), ShouldEqual, "testing.tRunner")
			})
		})
		Convey("Then File should match current file", func() {
			So(stackCall.File, ShouldMatch, ".*/testing/testing.go")
		})
		Convey("Then error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
	stackCall, err = GetCaller(2)
	Convey("When Getting Caller After Caller Method*2 from CalledFunc", t, func() {
		Convey("Given the callerFunc reference", func() {
			Convey("Then Caller's Name should equal runtime.goexit", func() {
				So(stackCall.Func.Name(), ShouldEqual, "runtime.goexit")
			})
		})
		Convey("Then File should match current file", func() {
			So(stackCall.File, ShouldMatch, ".*/runtime/.+")
		})
		Convey("Then error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
	stackCall, err = GetCaller(3)
	Convey("When Getting Caller After Caller Method*3 from CalledFunc", t, func() {
		Convey("Then error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
}
