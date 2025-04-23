package main
import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main(){
	royhan := Man{"Royhan"}
	royhan.Married()	

	fmt.Println("royhan :", royhan.Name) // {Royhan}
}