package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/user","description":"Operations about Users\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/user":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/user","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"createUser","type":"","summary":"create users","parameters":[{"paramType":"body","name":"body","description":"\"body for user content\"","dataType":"User","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.User","responseModel":"User"},{"code":403,"message":"body is empty","responseModel":""}]}]}],"models":{"Address":{"id":"Address","properties":{"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"DeletedAt":{"type":"\u0026{time Time}","description":"","format":""},"ID":{"type":"uint64","description":"","format":""},"UpdatedAt":{"type":"\u0026{time Time}","description":"","format":""},"addressLine1":{"type":"string","description":"","format":""},"addressLine2":{"type":"string","description":"","format":""},"addressType":{"type":"uint8","description":"","format":""},"attn":{"type":"string","description":"","format":""},"city":{"type":"string","description":"","format":""},"country":{"type":"string","description":"","format":""},"phone":{"type":"string","description":"","format":""},"postalCode":{"type":"string","description":"","format":""},"profileId":{"type":"int64","description":"","format":""},"stateProvince":{"type":"string","description":"","format":""}}},"Profile":{"id":"Profile","properties":{"AddressID":{"type":"int64","description":"","format":""},"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"DeletedAt":{"type":"\u0026{time Time}","description":"","format":""},"ID":{"type":"uint64","description":"","format":""},"UpdatedAt":{"type":"\u0026{time Time}","description":"","format":""},"UserID":{"type":"int64","description":"","format":""},"addresses":{"type":"array","description":"","items":{"$ref":"Address"},"format":""},"birthDate":{"type":"\u0026{time Time}","description":"","format":""},"firstName":{"type":"string","description":"","format":""},"lastName":{"type":"string","description":"","format":""},"middleName":{"type":"string","description":"","format":""},"phone":{"type":"string","description":"","format":""}}},"User":{"id":"User","properties":{"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"DeletedAt":{"type":"\u0026{time Time}","description":"","format":""},"ID":{"type":"uint64","description":"","format":""},"Profile":{"type":"Profile","description":"","format":""},"UpdatedAt":{"type":"\u0026{time Time}","description":"","format":""},"email":{"type":"string","description":"","format":""},"mobile":{"type":"string","description":"","format":""},"password":{"type":"string","description":"","format":""},"status;default(0)":{"type":"uint8","description":"","format":""},"userName":{"type":"string","description":"","format":""}}}}}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
	if beego.BConfig.WebConfig.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocAPI["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.APIs {
				a.Path = urlReplace(k + a.Path)
				v.APIs[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocAPI[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
