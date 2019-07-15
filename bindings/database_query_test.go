package bindings

import (
	"testing"

	"github.com/saveav/pelacore/app/database"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDatabaseQuery(t *testing.T) {
	Convey("Database query payload", t, func() {
		dbq := DatabaseQuery{}
		dbq.Connection.Driver = database.Postgres
		dbq.Connection.Host = "localhost"
		dbq.Connection.Username = "host"
		dbq.Query = "SELECT * FROM table"

		Convey("When required fields are blank", func() {
			Convey("It fails without driver", func() {
				dbq.Connection.Driver = ""
				So(dbq.Validate().Error(), ShouldEqual, ErrBlankDriver.Error())
			})

			Convey("It fails without host", func() {
				dbq.Connection.Host = ""
				So(dbq.Validate().Error(), ShouldEqual, ErrBlankHost.Error())
			})

			Convey("It fails without username", func() {
				dbq.Connection.Username = ""
				So(dbq.Validate().Error(), ShouldEqual, ErrBlankUsername.Error())
			})

			Convey("It fails without query", func() {
				dbq.Query = ""
				So(dbq.Validate().Error(), ShouldEqual, ErrBlankQuery.Error())
			})
		})
	})
}
