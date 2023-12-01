package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	// read the file
	reader, err := os.Open("Customer.json")
	if err != nil {
		panic(err)
	}

	// decode the json
	decoder := json.NewDecoder(reader)

	// choose the struct we want to insert the json to
	customer := &Customer{}

	// initiate the decoder result into the chosen struct
	decoder.Decode(customer)

	fmt.Println(customer)
}
