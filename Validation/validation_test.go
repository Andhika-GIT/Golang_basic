package validation

import (
	"fmt"
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

func TestMultipleTags(t *testing.T) {
	validate := validator.New()

	// var1 := "eko123"
	var2 := 12313

	err := validate.Var(var2, "required,number")

	if err != nil {
		fmt.Println(err.Error())
	}
}

// -------------- STRUCT VALIDATION ---------------

type User struct {
	Username string `validate:"required,min=3"`
	Email    string `validate:"required,email,min=3"`
}

func TestStructValidation(t *testing.T) {
	validate := validator.New()

	user1 := User{
		Username: "Andhika",
		Email:    "asdfasfsdf",
	}

	err := validate.Struct(user1)

	if err != nil {
		fmt.Println(err)
	}
}
