package test
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"goWebApp/auth/models"
	"fmt"
)

func TestUserRepo(t *testing.T) {
	fmt.Println("Start testing ....")
	userRepo := models.NewUserRepo()
	userRef := userRepo.NewEntry()
	userRef.Email = "abc@mail.com"
	mailingAddress := models.Address{AddressType: models.MAILING_ADDRESS}
	billingAddress := models.Address{AddressType: models.BILLING_ADDRESS}
	userRef.Profile.Addresses = append(userRef.Profile.Addresses, &mailingAddress)
	userRef.Profile.Addresses = append(userRef.Profile.Addresses, &billingAddress)
	err := userRepo.Save()

	Convey("Subject: Test Insert User Login\n", t, func() {
		Convey("User ID return Should Larger than 0", func() {
			So(userRef.Id, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})


	userRepo = models.NewUserRepo()
	qb := userRepo.QueryBuilder()
	userRepo.Get(qb)
	/*userRepo.Save()

	Convey("Subject: Test Read User with Profile\n", t, func() {
		Convey("Profil ID return Should Larger than 0", func() {
			So(userRef.Profile.Addresses[0].AddressLine1, ShouldEqual, "123 main st")
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})*/
}

