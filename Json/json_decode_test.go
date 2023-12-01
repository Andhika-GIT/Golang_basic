package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeToObject(t *testing.T) {
	jsonRequest := `{"Firstname":"andhika","Lastname":"nugraha","Skills":["Excel","word"]}`

	// convert the json to slice of bytes
	jsonBytes := []byte(jsonRequest)

	// take the struct pointer to some variabel
	user := &User{}

	// convert json to struct (Json.Unmarshal(jsondata, ourvariabel))
	err := json.Unmarshal(jsonBytes, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(user))
	fmt.Println(user)

}

func TestDecodeToSlice(t *testing.T) {
	jsonRequest := `{"Firstname":"Hubla","Lastname":"Gobs","Skills":["Word","Excel"],"Addresses":[{"Street":"Tole iskandar","Country":"Indonesia"},{"Street":"Juanda","Country":"Indonesia"}]}`

	jsonBytes := []byte(jsonRequest)

	user := &User{}

	err := json.Unmarshal(jsonBytes, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
