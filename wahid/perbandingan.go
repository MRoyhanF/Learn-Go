package main

import "fmt"

func main() {
	var name1 = "Royhan"
	var name2 = "Royhan"
	var result bool = name1 == name2
	fmt.Println("name1 == name2", result)

	var a = "a"
	var b = "b"
	var result2 bool = a < b
	fmt.Println("a < b", result2)
}