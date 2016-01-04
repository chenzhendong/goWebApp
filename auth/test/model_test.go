package test
import (
	"testing"
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
	"goWebApp/auth/models"
)

func TestOrm(t *testing.T) {
	o := orm.NewOrm()
	userLogin := models.UserLogin{Email: "abc@test.com"}
	id, err := o.Insert(&userLogin)

	Convey("Subject: Test Insert User Login\n", t, func() {
		Convey("ID return Should Larger than 0", func() {
			So(id, ShouldBeGreaterThan, 0)
		})
		Convey("The Error Should be Nil", func() {
			So(err, ShouldBeNil)
		})
	})
}

