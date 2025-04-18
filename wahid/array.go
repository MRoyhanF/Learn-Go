package main

import "fmt"

func main() {
	var names [3]string // Array of strings with a length of 3
	names[0] = "Muahamd"
	names[1] = "Royhan"
	names[2] = "Fadhli"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{ // generate array directly with values
		1, 
		2, 
		// default 0 for int
	}
	fmt.Println(values)
} 