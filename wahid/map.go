package main

import "fmt"

func main() {
	person:= map[string]string{
		"name": "John",
		"address": "123 Main St",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"]) // This will return an empty string since "age" is not a key in the map

	book := make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "John Doe"
	book["publisher"] = "Tech Books Publishing"

	fmt.Println(book)
	delete(book, "publisher") // Remove the publisher key from the map
	fmt.Println(book)
}