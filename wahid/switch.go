package main
import "fmt"

func main() {
	name:= "roy"

	switch name {
	case "roy":
		fmt.Println("Hello Roy")
	case "joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hello, Boleh Kenalan?")
	}

	// switch statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch statement with condition
	leghth := len(name)
	switch {
	case leghth > 10:
		fmt.Println("Nama terlalu panjang")
	case leghth < 5:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Nama tidak valid")
	}
}