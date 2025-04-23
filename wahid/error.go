package main
import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("error: pembagi tidak boleh nol")
	}
	return nilai / pembagi, nil
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil pembagian:", hasil)
	} else {
		fmt.Println(err)
	}

}