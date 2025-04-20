package main
import "fmt"

func sayHelloTo(firstName string, lastName string){
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func main() {
	sayHelloTo("John", "Doe")
	sayHelloTo("Jane", "Smith")
}