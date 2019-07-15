package database

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConnection(t *testing.T) {
	Convey("Connection string generation", t, func() {
		Convey("When driver is unknown", func() {
			con := Connection{
				Driver: "unknown",
			}
			Convey("An error is returned", func() {
				_, err := con.String()
				So(err, ShouldEqual, ErrUnknownDriver)
			})
			Convey("A blank connection string is returned", func() {
				conStr, _ := con.String()
				So(conStr, ShouldBeBlank)
			})
		})

		Convey("When driver is Postgres", func() {
			con := Connection{
				Driver:   Postgres,
				Host:     "localhost",
				Port:     "1234",
				Username: "root",
				Database: "somedb",
			}
			Convey("It can build a connection string", func() {
				conStr, err := con.String()
				So(conStr, ShouldEqual, "host=localhost port=1234 user=root dbname=somedb")
				So(err, ShouldBeNil)
			})
			Convey("It can build a connection string with password", func() {
				con.Password = "secret"
				conStr, err := con.String()
				So(conStr, ShouldEqual, "host=localhost port=1234 user=root password=secret dbname=somedb")
				So(err, ShouldBeNil)
			})
			Convey("It can build a connection string with SSLMode", func() {
				con.SSLMode = true
				conStr, err := con.String()
				So(conStr, ShouldEqual, "host=localhost port=1234 user=root dbname=somedb sslmode=require")
				So(err, ShouldBeNil)
			})
		})
	})
}
