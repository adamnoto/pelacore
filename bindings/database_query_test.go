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
		dbq.Connection.Database = "database"
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

			Convey("It fails without database", func() {
				dbq.Connection.Database = ""
				So(dbq.Validate().Error(), ShouldEqual, ErrBlankDatabase.Error())
			})

			Convey("It fails without query", func() {
				dbq.Query = ""
				So(dbq.Validate().Error(), ShouldEqual, ErrBlankQuery.Error())
			})
		})

		Convey("It can convert the data into DatabaseConnection object", func() {
			dbq.Connection.Password = "secret"
			dbq.Connection.Port = "1234"
			conn := dbq.ToDatabaseConnection()
			So(conn.Database, ShouldEqual, dbq.Connection.Database)
			So(conn.Driver, ShouldEqual, dbq.Connection.Driver)
			So(conn.Host, ShouldEqual, dbq.Connection.Host)
			So(conn.Username, ShouldEqual, dbq.Connection.Username)
			So(conn.Password, ShouldEqual, dbq.Connection.Password)
			So(conn.Port, ShouldEqual, dbq.Connection.Port)
			So(conn.SSLMode, ShouldEqual, dbq.Connection.SSLMode)
		})
	})
}
