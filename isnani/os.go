package main
import(
	"os"
	"fmt"
)

func main(){
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
	} else {
		fmt.Println("Hostname:", hostname)
	}
}