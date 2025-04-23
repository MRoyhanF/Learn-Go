package main
import "fmt"

func sumAll (numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println(total) // Output: 15

	fmt.Println(sumAll(1, 2, 3)) // Output: 6
	fmt.Println(sumAll(1, 2, 3, 4, 5)) // Output: 15
	fmt.Println(sumAll()) // Output: 0

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...)) // Output: 40
}
