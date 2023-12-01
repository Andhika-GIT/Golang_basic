package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname string
	Lastname  string
	Id        int64
}

func TestEncodeObject(t *testing.T) {
	c1 := Customer{
		Firstname: "andhika",
		Lastname:  "nugraha",
		Id:        1023910293,
	}

	bytes, _ := json.Marshal(c1)

	fmt.Println(string(bytes))
}
