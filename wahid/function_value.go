package main
import "fmt"

func getGoodBye (name string) string {
	return "Goodbye " + name
}

func main() {
	contoh := getGoodBye
	misal := contoh

	fmt.Println(contoh("Royhan")) // Output: Goodbye John
	fmt.Println(misal("Fadhli")) // Output: Goodbye John
}