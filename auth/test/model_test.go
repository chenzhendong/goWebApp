package test
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"goWebApp/auth/models"
	"fmt"
)

func TestUserRepo(t *testing.T) {
	userRepo := new(models.UserRepo)
	userRef := userRepo.New()
	userRef.Email = "abc@mail.com"
	err := userRepo.Save()

	fmt.Println(userRepo.Collection[0])

	Convey("Subject: Test Insert User Login\n", t, func() {
		Convey("User ID return Should Larger than 0", func() {
			So(userRef.Id, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})

	err = userRepo.LoadChildren()



	Convey("Subject: Test Read User with Profile\n", t, func() {
		Convey("Profil ID return Should Larger than 0", func() {
			So(userRef.Profile.Id, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})
}

