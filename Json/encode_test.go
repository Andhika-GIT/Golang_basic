package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Dhika")
	logJson(1)
	logJson(true)
	logJson([]string{"Eko", "kurniawan", "Khannedy"})
}

func TestEncodeProduct(t *testing.T) {
	p := Product{
		Id:    "uisuadfiu29",
		Name:  "greatproduct",
		Price: 13293,
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
