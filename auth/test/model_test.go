package test
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"goWebApp/auth/models"
	"fmt"
)

func TestUserRepo(t *testing.T) {

	repo := models.RepoInstance

	// test single insert
	writeUser := models.User{Email:"abc@mail.com", Profile: models.Profile{FirstName: "John", Addresses: []models.Address{{AddressLine1: "123 main st"},{AddressLine1: "456 wall st"}}}}
	fmt.Println("Origin Record Before Insert: ", writeUser)
	readUser, _ := repo.SaveUser(writeUser)
	fmt.Println("User After Insert: ", readUser)

	Convey("Subject: Test Insert User\n", t, func() {
		Convey("User ID return Should Larger than 0", func() {
			So(readUser.ID, ShouldBeGreaterThan, 0)
		})
		Convey("Profile ID return Should Larger than 0", func() {
			So(readUser.Profile.ID, ShouldBeGreaterThan, 0)
		})
		Convey("Two Addresses ID return Should Larger than 0", func() {
			So(readUser.Profile.Addresses[0].ID, ShouldBeGreaterThan, 0)
			So(readUser.Profile.Addresses[1].ID, ShouldBeGreaterThan, 0)
		})
	})

	// test cache clone
	writeUser = readUser
	writeUser.Profile.FirstName = "James"
	writeUser.Profile.Addresses[0].Attn = "mailing"
	writeUser.Profile.Addresses[1].Attn = "billing"
	readUser,_ = repo.Get(readUser.ID)

	Convey("Subject: Make sure return objects from cache are cloned copies, not the same one in the cache\n", t, func() {
		Convey("Addresses Attn should return empty string or nil", func() {
			So(readUser.Profile.FirstName, ShouldEqual, "John")
			So(readUser.Profile.Addresses[0].Attn, ShouldEqual, "")
			So(readUser.Profile.Addresses[1].Attn, ShouldEqual, "")
		})
	})

	// test single update
	repo.SaveUser(writeUser)
	readUser,_ = repo.Get(writeUser.ID)

	fmt.Println("Read User After Update:", readUser)

	Convey("Subject: Test Update User\n", t, func() {
		Convey("Address 0 Attn should return 'mailing' and 'billing'", func() {
			So(readUser.Profile.Addresses[0].Attn, ShouldEqual, "mailing")
			So(readUser.Profile.Addresses[1].Attn, ShouldEqual, "billing")
		})
	})

	// test multiple inserts
	userSlice := make([]models.User, 0)
	writeUser = models.User{Email:"def@mail.com", Profile: models.Profile{FirstName: "Jane", Addresses: []models.Address{{AddressLine1: "123 main st"},{AddressLine1: "456 wall st"}}}}
	userSlice = append(userSlice, writeUser)
	writeUser = models.User{Email:"fgh@mail.com", Profile: models.Profile{FirstName: "Sam", Addresses: []models.Address{{AddressLine1: "133 main st"},{AddressLine1: "436 wall st"}}}}
	userSlice = append(userSlice, writeUser)
	repo.SaveUsers(userSlice)

	builder := repo.GetQueryBuilder()
	userSlice = repo.FindAll(builder)

	Convey("Subject: Test Batch Insert User\n", t, func() {
		Convey("All User ID in Slice Should Larger than 0", func() {
			So(userSlice[0].ID, ShouldBeGreaterThan, 0)
			So(userSlice[1].ID, ShouldBeGreaterThan, 0)
			So(userSlice[2].ID, ShouldBeGreaterThan, 0)
		})
	})

	//insert duplicate email
	var err error
	writeUser = models.User{Email:"abc@mail.com", Profile: models.Profile{FirstName: "John", Addresses: []models.Address{{AddressLine1: "123 main st"},{AddressLine1: "456 wall st"}}}}
	_, err = repo.SaveUser(writeUser)
	Convey("Subject: Test Save Error\n", t, func() {
		Convey("Insert duplicate email should return an error", func() {
			So(err, ShouldNotBeNil)
			fmt.Println("\n\tError Message is:", err)
		})
	})


}

