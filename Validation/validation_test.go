package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationVariabel(t *testing.T) {
	validate := validator.New()

	var user string = ""

	err := validate.Var(user, "required")

	if err != nil {
		panic(err)
	}
}

func TestValidationTwoVariabels(t *testing.T) {
	validate := validator.New()

	var1 := "hello"
	var2 := "hello"

	err := validate.VarWithValue(var1, var2, "eqfield")

	if err != nil {
		panic(err)
	}
}
