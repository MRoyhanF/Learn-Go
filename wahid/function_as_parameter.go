package main
import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	filteredName := filter(name)
	fmt.Println("Hello:", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "****"
	} else if name == "Babi" {
		return "****"
	} else {
		return name
	}
}

func main () {
	sayHelloWithFilter("Royhan", spamFilter) // Output: Hello: Royhan

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter) // Output: Hello: ****
}