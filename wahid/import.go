package main
import (
	"wahid/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Wahid")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // cannot access version because it's unexported
	// fmt.Println(helper.sayGoodbye("Wahid")) // cannot access sayGoodbye because it's unexported
}