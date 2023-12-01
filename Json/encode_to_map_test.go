package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeToMaps(t *testing.T) {
	jsonRequest := `{"id":"uasdfiu123","name":"greatproduct","price":125000}`
	jsonBytes := []byte(jsonRequest)

	var response map[string]interface{}

	err := json.Unmarshal(jsonBytes, &response)
	if err != nil {
		panic(err)
	}

	for key, value := range response {
		fmt.Printf("%v : %v\n", key, value)
	}
}
