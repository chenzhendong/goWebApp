package test
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"goWebApp/auth/models"
	"fmt"
)

func TestUserRepo(t *testing.T) {
	userRef := &models.UserLogin{Email: "abc@test.com"}
	uid, err := userRef.Create()

	Convey("Subject: Test Insert User Login\n", t, func() {
		Convey("User ID return Should Larger than 0", func() {
			So(uid, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})

	userRef = &models.UserLogin{Id: uid}
	err = userRef.Read()
	fmt.Println(userRef.Profile)

	Convey("Subject: Test Read User with Profile\n", t, func() {
		Convey("Profil ID return Should Larger than 0", func() {
			So(userRef.Profile.Id, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})
}

