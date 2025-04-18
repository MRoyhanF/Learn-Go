package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRoyhan NoKTP = "11111111"

	var contoh string = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRoyhan) // 11111111
	fmt.Println(contohKtp) // 2222222222

}