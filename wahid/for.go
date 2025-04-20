package main
import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Counter", counter)
	// 	counter++
	// }
	
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Counter", counter)
	}
	fmt.Println("Selesai")

	names := []string{"Roy", "Joko", "Budi", "Siti", "Ani"}
	// for loop with index
	for i := 0; i < len(names); i++ {
		fmt.Println("Nama", i, "=", names[i])
	}
	// for range
	for index, name := range names {
		fmt.Println("Nama", index, "=", name)
	}
	// for range with underscore without using key
	for _, name := range names {
		fmt.Println(name)
	}
}