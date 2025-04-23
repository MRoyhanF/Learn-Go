package main
import "fmt"

func random() any {
	return "OK"
}

func main(){
	var result any = random()
	// fmt.Println(result) // Output: OK
	// resultString := result.(string)
	// fmt.Println(resultString) // Output: OK

	// resultInt := result.(int)
	// fmt.Println(resultInt) // Output: panic: interface conversion: interface {} is string, not int
	
	switch value := result.(type) {
		case string:
			fmt.Println("String:", value) // Output: String: OK
		case int:
			fmt.Println("Int:", value) // Output: Int: 0
		default:
			fmt.Println("Unknown type") // Output: Unknown type
	}

}