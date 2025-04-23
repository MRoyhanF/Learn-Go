package main
import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blacklisted", name)
	} else {
		fmt.Println("Welcome", name)
	}	
}

func main() {
	blacklist := func(name string) bool {
		return name == " Anjing"
	}

	registerUser("Royhan", blacklist) // Output: Welcome Royhan

	registerUser("Anjing", func(name string)bool{
		return name == "Anjing"
	}) // Output: You are blacklisted Anjing

}