package test
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"goWebApp/auth/models"
)

func TestUserRepo(t *testing.T) {
	userRepo := new(models.UserRepo)
	userRef := userRepo.New()
	userRef.Email = "abc@mail.com"
	userRef.Profile.BillingAddress.Attn = "Billing"
	userRef.Profile.MailingAddress.Attn = "Mailing"
	err := userRepo.Save()

	Convey("Subject: Test Insert User Login\n", t, func() {
		Convey("User ID return Should Larger than 0", func() {
			So(userRef.Id, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})


	userRef.Profile.FirstName = "John"
	userRef.Profile.LastName = "Smith"
	userRef.IsChanged = true
	userRef.Profile.BillingAddress.AddressLine1 = "123 main st"
	userRef.Profile.BillingAddress.IsChanged = true

	userRepo.Save()

	Convey("Subject: Test Read User with Profile\n", t, func() {
		Convey("Profil ID return Should Larger than 0", func() {
			So(userRef.Profile.BillingAddress.AddressLine1, ShouldEqual, "123 main st")
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})
}

