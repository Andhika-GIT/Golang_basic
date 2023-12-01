package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeArray(t *testing.T) {
	u1 := User{
		Firstname: "Hubla",
		Lastname:  "Gobs",
		Skills:    []string{"Word", "Excel"},
		Addresses: []Address{
			{
				Street:  "Tole iskandar",
				Country: "Indonesia",
			},
			{
				Street:  "Juanda",
				Country: "Indonesia",
			},
		},
	}

	bytes, _ := json.Marshal(u1)

	fmt.Println(string(bytes))
}
