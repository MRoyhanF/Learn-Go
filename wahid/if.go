package main
import "fmt"

func main() {
	name := "roy"

	if name == "roy" {
		fmt.Println("Hello Roy")
	} else if name == "joko"{
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hello, Boleh Kenalan?")
	}

	// if statement
	if leghth := len(name); leghth > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}