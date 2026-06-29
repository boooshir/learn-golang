package helper

import "fmt"

type validationError struct {
	Message string
}

func (validation *validationError) Error() string {
	return validation.Message
}

type notfoundError struct {
	Message string
}

func (n *notfoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}
	if id != "arif" {
		return &notfoundError{"not found error"}
	}
	return nil
}
func CustomError() {
	err := SaveData("", nil)
	if err != nil {
		//   if validationErr, ok := err.(*validationError); ok {
		//     fmt.Println("validation error:", validationErr.Error())
		//   } else if notfoundErr, ok := err.(*notfoundError); ok {
		//     fmt.Println("not found error", notfoundErr.Error())
		//   } else {
		//     fmt.Println("unknown error", err.Error())
		//   }
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Error())
		case *notfoundError:
			fmt.Println("not found error", finalError.Error())
		default:
			fmt.Println("Unknown error", finalError.Error())
		}
	} else {
		fmt.Println("sukses")
	}
}
