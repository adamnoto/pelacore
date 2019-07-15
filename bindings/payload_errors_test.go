package bindings

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPayloadErrors(t *testing.T) {
	Convey("Payload errors", t, func() {
		Convey("When there are more than 1 error", func() {
			Convey("Errors are concatenated", func() {
				err := PayloadErrors{}
				err.Append(errors.New("A is missing"))
				err.Append(errors.New("B is invalid"))
				So(err.Error(), ShouldEqual, "A is missing. B is invalid")
			})
		})
	})
}
