package test

import (
	"log"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"../auth"
)

func TestAuthModel(t *testing.T) {

	userLogin := auth.UserLogin{ Email:"john@yahoo.com", Status: auth.NEW_STATUS}
	_,err := auth.X.Insert(&userLogin);


	Convey("Subject: Test create user by single Email \n", t, func() {
		Convey("Should Return Without Error", func() {
			So(err, ShouldBeNil)
		})
		Convey("The Id Should Greater than 0", func() {
			log.Printf("The return Login Obj: %v", userLogin)
			So(userLogin.Id, ShouldBeGreaterThan, 0)
		})
	})
}
