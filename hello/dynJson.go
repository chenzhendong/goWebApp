package main
import (
	"fmt"
	"encoding/json"
    "github.com/Jeffail/gabs"
)

func main() {
	input := []byte(`{
    "firstName": "John",
    "lastName": "Smith",
    "isAlive": true,
    "age": 25,
    "heightCm": "167.64",
    "addresses": [
    {
        "streetAddress": "21 2nd Street",
        "city": "New York",
        "state": "NY",
        "postalCode": "10021-3100",
        "phone": null
    },
    {
        "streetAddress": "22 2nd Street",
        "city": "Washington",
        "state": "DC",
        "postalCode": "20001",
        "phone": null
    }

    ]
}`)
	var p map[string]interface{}

	json.Unmarshal(input, &p)

	fmt.Println(p["addresses"].([]interface{})[1].(map[string]interface{})["postalCode"])

    jsonParsed, _ := gabs.ParseJSON(input)

    v, _ := jsonParsed.Path("addresses").Index(1).Path("city").Data().(string)
    fmt.Println(v)
}