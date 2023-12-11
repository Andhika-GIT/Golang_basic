package validation

import (
	"fmt"
	"strings"
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

// ----------- NESTED STRUCT AND COLLECTION -----------

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

// ----------- NESTED STRUCT AND MAPS -----------
type School struct {
	Name string `validate:"required"`
}

type Student struct {
	Name    string   `validate:"required,min=3"`
	Hobbies []string `validate:"dive,required"`

	// Schools -> required
	// map[string] -> dive,keys,required,min=2
	// school -> endkeys,dive
	Schools map[string]School `validate:"required,dive,keys,required,min=2,endkeys"`

	// Wallet ->
	// map[string] -> dive, keys, required
	// int -> endkeys,required
	Wallet map[string]int `validate:"dive,keys,required,endkeys,required"`
}

func TestNestedStructMaps(t *testing.T) {

	school := map[string]School{
		"SD": {
			Name: "Santo aloysius",
		},
		"SMP": {
			Name: "",
		},
	}

	wallet := map[string]int{
		"ovo":   2,
		"gopay": 0,
	}

	student := Student{
		Name:    "Hubla",
		Hobbies: []string{"berack"},
		Schools: school,
		Wallet:  wallet,
	}

	validate := validator.New()

	err := validate.Struct(student)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error field : ", fieldError.Field(), "\ton tag", fieldError.Tag(), "\twith error : ", fieldError.Error())
		}
	}
}

// ----------- ALIAS TAGS -----------

func TestAliasTags(t *testing.T) {
	validate := validator.New()

	// create our own tag validation
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Name  string `validate:"varchar"`
		Owner string `validate:"varchar"`
	}
}

// ----------- CUSTOM VALIDATION -----------
func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)

	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}

	return true
}

func MustValidEmail(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)

	if ok {
		if strings.Contains(value, "@") == false {
			return false
		}

		if strings.Contains(value, ".com") == false {
			return false
		}

	}

	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)
	validate.RegisterValidation("email", MustValidEmail)

	type User struct {
		Username string `validate:"username"`
		Email    string `validate:"email"`
	}

	user := User{
		Username: "HUBLA",
		Email:    "hubla",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err)
	}
}

// ----------- OR RULE -----------

func TestOrRule(t *testing.T) {
	type User[T any] struct {
		Username T      `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	user := User[int]{
		Username: 10239,
		Password: "testpass",
	}

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err)
	}
}

// ----------- GET THE DEFAULT VALUE AND SECOND VALUE -----------

func MustEqual(field validator.FieldLevel) bool {
	// GetStructFieldOK2 returns 4 things
	// the value itself, the value type, is the value is nullable, is the value nil or empty
	// we only need first thing (the value) and four thing (is the value nil or emtpy)
	value, _, _, ok := field.GetStructFieldOK2()

	if !ok {
		fmt.Println("field not ok")
	}

	//	ex :  Username `validate:"mustEqual = Email"`
	// Username -> first value
	// Email 	-> second value
	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	// return true or false whenver if the first value is the same as second value
	return firstValue == secondValue
}

func TestValidationCrossField(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("must_equal", MustEqual)

	type User struct {
		Username string `validate:"required"`
		Password string `validate:"required,min=3"`

		// ConfirmPassword string `validate:"must_equal=Password"`
		// ConfirmPassword -> first value
		// Password 	   -> second value
		ConfirmPassword string `validate:"required,must_equal=Password"`
	}

	user := User{
		Username:        "Hubla",
		Password:        "halo",
		ConfirmPassword: "helo",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err)
	}
}

// ----------- STRUCT LEVEL VALIDATION -----------
type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Phone    string `validate:"required"`
}

func MustValidRegisterField(level validator.StructLevel) {
	// get the RegisterRequest into variabel
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Phone || registerRequest.Username == registerRequest.Email {
		// success
	} else {
		// reportError return 4 value
		// variabel of our choice, Field name, Field name of the struct, Tag name, parameter
		level.ReportError(registerRequest.Username, "Username", "username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()

	validate.RegisterStructValidation(MustValidRegisterField, RegisterRequest{})

	request := RegisterRequest{
		Username: "halo",
		Email:    "halo",
		Phone:    "helo",
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}

}
