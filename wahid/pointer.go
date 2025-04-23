package main
import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	var address3 *Address = &address1
	address4 := address1

	address2.City = "Jakarta"
	address3.City = "Jawa Tengah"
	address4.City = "Bandung"
	fmt.Println("address1 :", address1) // {Subang Jawa Barat Indonesia}
	fmt.Println("address2 :", address2) // {Jakarta Jawa Barat Indonesia}
	fmt.Println("address3 :",address3) // {Jawa Tengah Jawa Barat Indonesia}
	fmt.Println("address4 :", address4) // {Bandung Jawa Barat Indonesia}
}