package main
import "fmt"

func factorialLoop (value int)int{
	result := 1
	for i:= value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 0 {
		return 1
	}
	return value * factorialRecursive(value - 1)
}

func main () {
	fmt.Println("Factorial of 5 using loop:", factorialLoop(5))
	fmt.Println("Factorial of 10 using loop:", factorialLoop(10))

	// Recursive function call
	fmt.Println("Factorial of 5 using recursion:", factorialRecursive(5))
	fmt.Println("Factorial of 10 using recursion:", factorialRecursive(10))
}