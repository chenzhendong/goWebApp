package test
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"goWebApp/auth/models"
	"fmt"
)

func TestUserRepo(t *testing.T) {

	db := models.DB

	writeUser := models.UserLogin{Email:"abc@mail.com", Profile: models.UserProfile{FirstName: "John", Addresses: []models.Address{{AddressLine1: "123 main st"},{AddressLine1: "456 wall st"}}}}
	db.Create(&writeUser)
	fmt.Println(writeUser)
	readUser := models.UserLogin{ID: writeUser.ID}
	db.First(&readUser)
	fmt.Println("Read User After Insert: ", readUser)

	Convey("Subject: Test Insert User\n", t, func() {
		Convey("User ID return Should Larger than 0", func() {
			So(readUser.ID, ShouldBeGreaterThan, 0)
		})
	})

	writeUser.Profile.Addresses[0].Attn = "mailling"

	db.Updates(writeUser)

	readUser = models.UserLogin{ID: writeUser.ID}
	db.First(&readUser)

	fmt.Println("Read User After Update:", readUser)

	Convey("Subject: Test Update User\n", t, func() {
		Convey("Address 0 Attn should return 'mailing'", func() {
			//So(readUser.Profile.Addresses[0].Attn, ShouldEqual, "mailing")
		})
	})

	/*fmt.Println("Start testing ....")
	userRepo := models.NewUserRepo()
	userRef := userRepo.NewEntry()
	userRef.Email = "abc@mail.com"
	mailingAddress := models.Address{AddressType: models.MAILING_ADDRESS}
	billingAddress := models.Address{AddressType: models.BILLING_ADDRESS}
	userRef.Profile.Addresses = append(userRef.Profile.Addresses, &mailingAddress)
	userRef.Profile.Addresses = append(userRef.Profile.Addresses, &billingAddress)
	userRef.Profile.FirstName = "John"
	userRef.Profile.Addresses[0].Attn = "mailing"
	userRef.Profile.Addresses[1].Attn = "billing"

	err := userRepo.Save()

	Convey("Subject: Test Insert User\n", t, func() {
		Convey("User ID return Should Larger than 0", func() {
			So(userRef.Id, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})


	userRepo = models.NewUserRepo()
	qb := userRepo.QueryBuilder()
	err = userRepo.Get(qb)
	userRef = userRepo.QueryEntries[1]

	Convey("Subject: Test Read User\n", t, func() {
		Convey("User Email should display", func() {
			So(userRef.Email, ShouldEqual, "abc@mail.com")
		})
		Convey("User Profile first name should match", func() {
			So(userRef.Profile.FirstName, ShouldEqual, "John")
		})
		Convey("Address 1 Attn should match", func() {
			So(userRef.Profile.Addresses[0].Attn, ShouldEqual, "mailing")
		})
		Convey("Address 2 Attn should match", func() {
			So(userRef.Profile.Addresses[1].Attn, ShouldEqual, "billing")
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})


	userRef.Email = "abc1@mail.com"
	userRef.Profile.FirstName = "John1"
	userRef.Profile.IsChanged = true
	userRef.Profile.Addresses[0].Attn = "mailing1"
	userRef.Profile.Addresses[0].IsChanged = true
	userRef.Profile.Addresses[1].Attn = "billing1"
	userRef.Profile.Addresses[1].IsChanged = true
	err = userRepo.Save()

	Convey("Subject: Test Update User \n", t, func() {
		Convey("User Email should display", func() {
			So(userRef.Email, ShouldEqual, "abc1@mail.com")
		})
		Convey("User Profile first name should match", func() {
			So(userRef.Profile.FirstName, ShouldEqual, "John1")
		})
		Convey("Address 1 Attn should match", func() {
			So(userRef.Profile.Addresses[0].Attn, ShouldEqual, "mailing1")
		})
		Convey("Address 2 Attn should match", func() {
			So(userRef.Profile.Addresses[1].Attn, ShouldEqual, "billing1")
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})

	userRepo = models.NewUserRepo()
	qb = userRepo.QueryBuilder()
	err = userRepo.Get(qb)
	userRef = userRepo.QueryEntries[1]

	Convey("Subject: Test Update User and Read Again\n", t, func() {
		Convey("User Email should display", func() {
			So(userRef.Email, ShouldEqual, "abc1@mail.com")
		})
		Convey("User Profile first name should match", func() {
			So(userRef.Profile.FirstName, ShouldEqual, "John1")
		})
		Convey("Address 1 Attn should match", func() {
			So(userRef.Profile.Addresses[0].Attn, ShouldEqual, "mailing1")
		})
		Convey("Address 2 Attn should match", func() {
			So(userRef.Profile.Addresses[1].Attn, ShouldEqual, "billing1")
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})*/
}

