package helper

import "fmt"

var version = "1.0.0"
var Application = "Belajar Golang"

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	fmt.Print(sayGoodbye("Wahid"))
	fmt.Println(version)
}