package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncode(t *testing.T) {
	// create new file
	writer, err := os.Create("sample_output.json")
	if err != nil {
		panic(err)
	}

	// encode the file writer (preparing the new file)
	encoder := json.NewEncoder(writer)

	// preparing the struct variabel that later convert to json and save to new file
	customer := Customer{
		Firstname: "Hubla",
		Lastname:  "Hey",
		Id:        12239304039,
	}

	// encode the struct variabel into json, and save it into new file
	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)

}
