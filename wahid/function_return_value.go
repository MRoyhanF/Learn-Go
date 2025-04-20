package main
import "fmt"

func getHello(name string) string {
	hello := "Hello, " + name + "!"
	return hello
}

func main() {
	result := getHello("John")
	fmt.Println(result)
	fmt.Println(getHello("Jane"))
	fmt.Println(getHello("Smith"))
}