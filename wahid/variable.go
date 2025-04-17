package main
import "fmt"

func main()  {
	// var name string;
	var name string = "Royhan" // using var keyword
	fmt.Println(name)

	name = "Fadhli"
	fmt.Println(name)

	lastName := "lastName" // using short variable declaration without var keyword
	fmt.Println(lastName)

	lastName = "lastName2"
	fmt.Println(lastName)

	var ( // multiple variable declaration
		namaDepan string = "Royhan"
		namaBelakang string = "Fadhli"
	)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println("Hello", namaDepan, namaBelakang)

}