package main 
import "fmt"

func main() {
	// continue statement
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Counter", i)
	}
	fmt.Println("Selesai")
}