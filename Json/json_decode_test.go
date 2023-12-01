package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{"Firstname":"andhika","Lastname":"nugraha","Id":12323}`

	// convert the json to slice of bytes
	jsonBytes := []byte(jsonRequest)

	// take the struct pointer to some variabel
	customer := &Customer{}

	// convert json to struct (Json.Unmarshal(jsondata, ourvariabel))
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(customer))
	fmt.Println(customer)

}
