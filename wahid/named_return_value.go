package main
import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Royhan"
	middleName = "Fadhli"
	lastName = "Alfian"
	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println("First Name:", a)
	fmt.Println("Middle Name:", b)
	fmt.Println("Last Name:", c)
}