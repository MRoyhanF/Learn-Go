package main
import "fmt"

func getFullName() (string, string) {
	return "Royhan" , "Fadhli"
}

func main() {
	fistName, lastName := getFullName()
	fmt.Println("First Name:", fistName)
	fmt.Println("Last Name:", lastName)

	pertama, _ := getFullName() // ignore second value
	// _ is a blank identifier, it means we don't care about the second value
	fmt.Println("Pertama:", pertama)
}