package main
import "fmt"

func endApp() {
	fmt.Println("Ending application...")
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application error!")
	}
}

func main() {
	runApp(true)
	fmt.Println("Main function...")
}