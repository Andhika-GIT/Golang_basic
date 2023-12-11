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

func TestStructValidation(t *testing.T) {
	type User struct {
		Username string `validate:"required,min=3"`
		Email    string `validate:"required,email,min=3"`
	}

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

func TestValidationErrors(t *testing.T) {
	type User struct {
		Username string `validate:"required,min=3"`
		Email    string `validate:"required,email,min=3"`
	}
	validate := validator.New()

	user := User{
		Username: "A",
		Email:    "E",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println("err")
		// insert all validation erorrs info into validationErrors variabel ( return []fielderror)
		validationErrors := err.(validator.ValidationErrors)

		// iterate through all errors
		for _, fieldError := range validationErrors {
			fmt.Println("error field : ", fieldError.Field(), "\ton tag", fieldError.Tag(), "\twith error : ", fieldError.Error())
		}
	}
}

func TestCrossField(t *testing.T) {
	type User struct {
		Username        string `validate:"required,min=3"`
		Password        string `validate:"required,min=3"`
		ConfirmPassword string `validate:"required,min=3,eqfield=Password"`
	}

	validate := validator.New()

	user := User{
		Username:        "A",
		Password:        "halo",
		ConfirmPassword: "helo",
	}

	err := validate.Struct(user)
	fmt.Println(err)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error field : ", fieldError.Field(), "\ton tag", fieldError.Tag(), "\twith error : ", fieldError.Error())
		}
	}
}

// ----------- NESTED STRUCT AND COLLECTION ----------_

type Employee struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

type Company struct {
	Name    string `validate:"required"`
	Address string `validate:"required"`
	// we use dive when it comes to slice collection, or nested struct
	Employees []Employee `validate:"dive,required"`
	Branch    []string   `validate:"dive,min=3"`
}

func TestNestedStructCollection(t *testing.T) {
	company := Company{
		Name:    "Hubla world",
		Address: "Depok country",
		Employees: []Employee{
			{
				Name:  "",
				Email: "",
			},
			{
				Name:  "",
				Email: "",
			},
		},
		Branch: []string{
			"h",
			"a",
		},
	}

	validate := validator.New()

	err := validate.Struct(company)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error field : ", fieldError.Field(), "\ton tag", fieldError.Tag(), "\twith error : ", fieldError.Error())
		}
	}
}
