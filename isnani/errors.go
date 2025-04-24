package main
import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetBtId(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "royhan" {
		return NotFoundError
	}
	return nil
}

func main(){
	err := GetBtId("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error occurred")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error occurred")
		} else {
			fmt.Println("Unknown error occurred")
		}
	} else {
		fmt.Println("No error occurred")
	}
}