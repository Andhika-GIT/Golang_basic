package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	firstname string
	lastname  string
	id        int64
}

func TestEncodeObject(t *testing.T) {
	c1 := Customer{
		firstname: "andhika",
		lastname:  "nugraha",
		id:        1023910293,
	}

	bytes, _ := json.Marshal(c1)

	fmt.Println(string(bytes))
}
