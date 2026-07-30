package golangvaliation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationField(t *testing.T) {
	validate := validator.New()

	var user string = ""

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateionTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	var user string = "1234"

	err := validate.Var(user, "required,numeric")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestParameter(t *testing.T) {
	validate := validator.New()

	user := "9966"

	err := validate.Var(user, "required,numeric,min=5,max=10")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginRequest := LoginRequest{
		Username: "booshir@gmail.com",
		Password: "ra",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()

	registerUser := RegisterUser{
		Username:        "boshir@gmail.com",
		Password:        "123456",
		ConfirmPassword: "1234erer",
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)

		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

func TestNested(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}
	validate := validator.New()

	registerUser := User{
		Id:   "223",
		Name: "boshir",
		Address: Address{
			City:    "Batu Caves",
			Country: "MY",
		},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)

		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
	}
	validate := validator.New()

	registerUser := User{
		Id:   "223",
		Name: "boshir",
		Address: []Address{
			{
				City:    "Batu Caves",
				Country: "MY",
			},
			{
				City:    "Batu Pahat",
				Country: "MY",
			},
		},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)

		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobbies []string  `validate:"dive,required,min=1"`
	}
	validate := validator.New()

	registerUser := User{
		Id:   "223",
		Name: "boshir",
		Address: []Address{
			{
				City:    "Batu Caves",
				Country: "MY",
			},
			{
				City:    "Batu Pahat",
				Country: "MY",
			},
		},
		Hobbies: []string{"hai"},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)

		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

func TestMap(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=1"`
		School    map[string]School `validate:"dive,keys,required,min=2,endkeys"`
	}

	validate := validator.New()

	registerUser := User{
		Id:   "223",
		Name: "boshir",
		Addresses: []Address{
			{
				City:    "Batu Caves",
				Country: "MY",
			},
			{
				City:    "Batu Pahat",
				Country: "MY",
			},
		},
		Hobbies: []string{"hai", "gaming"},
		School: map[string]School{
			"SD": {
				Name: "Hello",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)

		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}

}

func TestBasicMap(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=1"`
		School    map[string]School `validate:"dive,keys,required,min=2,endkeys"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=0"`
	}

	validate := validator.New()

	registerUser := User{
		Id:   "223",
		Name: "boshir",
		Addresses: []Address{
			{
				City:    "Batu Caves",
				Country: "MY",
			},
			{
				City:    "Batu Pahat",
				Country: "MY",
			},
		},
		Hobbies: []string{"hai", "gaming"},
		School: map[string]School{
			"SD": {
				Name: "Hello",
			},
			"SMP": {
				Name: "asdasd",
			},
		},
		Wallet: map[string]int{"2": 3, "asdas": 12},
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)

		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}

}

func TestTagAlias(t *testing.T) {

	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}
	seller := Seller{
		Id:     "asd123",
		Name:   "asd",
		Owner:  "asd",
		Slogan: "jaya jaya jaya ",
	}

	err := validate.Struct(seller)

	if err != nil {
		validateErr := err.(validator.ValidationErrors)

		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

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

func TestCustomValidation(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	login := LoginRequest{
		Username: "USERNAMETEST",
		Password: "123456",
	}
	err := validate.Struct(login)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)
		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())

	if err != nil {
		panic(err)
	}

	value := field.Field().String()

	if !regexNumber.MatchString(value) {
		return false
	}
	return len(value) == length
}

func TestCustomValidationParam(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}
	login := Login{
		Phone: "213123",
		Pin:   "222222",
	}
	err := validate.Struct(login)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)
		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}

}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	validate := validator.New()

	login := Login{
		Username: "234234324",
		Password: "1234dsad",
	}
	validate.Struct(login)

	err := validate.Struct(login)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)
		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("Field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCustomCrossField(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "test@example.com",
		Email:    "test@example.com",
		Phone:    "34324234",
		Name:     "Hello",
	}
	err := validate.Struct(user)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)
		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		//success
	} else {
		// failed
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}
func TestStructLevel(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "12345",
		Email:    "test@example.com",
		Phone:    "12345",
		Password: "rahasia",
	}
	err := validate.Struct(request)
	if err != nil {
		validateErr := err.(validator.ValidationErrors)
		for _, fieldErr := range validateErr {
			fmt.Println("error", fieldErr.Field(), "on tag ", fieldErr.Tag(), "with error ", fieldErr.Error())
		}
	}

}
