package main
import "fmt"

type Customer struct {
	Name, Address string
	Age          int
}

func (customer Customer) SayHello(name string) {
	fmt.Println("Hello,", name, "my name is", customer.Name + " and I am", customer.Age, "years old.")
}

func main() {
	var customer Customer
	customer.Name = "John Doe"
	customer.Address = "123 Main St"
	customer.Age = 30
	fmt.Println(customer)

	c := Customer{
		Name:    "John Doe",
		Address: "123 Main St",
		Age:     30,
	}
	fmt.Println(c)
	fmt.Printf("Name: %s, Address: %s, Age: %d\n", c.Name, c.Address, c.Age)
	fmt.Printf("%+v\n", c) // Print struct with field names
	fmt.Printf("%#v\n", c) // Print struct with field names and types
	fmt.Printf("%T\n", c)  // Print type of struct

	joko := Customer{
		Name:    "Joko",
		Address: "Jl. Jendral Sudirman",
		Age:     25,
	}
	fmt.Println(joko)
	budi := Customer{"Budi", "Jl. Merdeka", 28}
	fmt.Println(budi)



	// Call method
	joko.SayHello("Budi")
	budi.SayHello("Joko")
}