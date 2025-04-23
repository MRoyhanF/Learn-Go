package main
import (
	// "errors"
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type nonFoundError struct {
	Message string
}

func (n *nonFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation Error: ID tidak boleh kosong"}
	}
	if id != "Royahn" {
		return &nonFoundError{Message: "Not Found Error: Data tidak ditemukan"}
	}

	// oke
	return nil
}

func main() {
	err:= SaveData("", nil)
	if err != nil {
		// terjadi error
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationError.Error())
		// } else if nonFoundError, ok := err.(*nonFoundError); ok {
		// 	fmt.Println("not found error:", nonFoundError.Error())
		// } else {
		// 	fmt.Println("unknow error:", err.Error())
		// }

		switch finalError := err.(type){
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *nonFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknow error:", finalError.Error())
		}
	} else {
		// sukses
		fmt.Println("Data berhasil disimpan")
	}


}