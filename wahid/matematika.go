package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b

	fmt.Println("Hasil Penjumlahan:", c) // Hasil Penjumlahan: 30
	fmt.Println("Hasil Penjumlahan:", a+b) // Hasil Penjumlahan: 30
	fmt.Println("Hasil Penjumlahan:", a-b) // Hasil Penjumlahan: -10
	fmt.Println("Hasil Penjumlahan:", a*b) // Hasil Penjumlahan: 200
	fmt.Println("Hasil Penjumlahan:", a/b) // Hasil Penjumlahan: 0
	fmt.Println("Hasil Penjumlahan:", a%b) // Hasil Penjumlahan: 10
	fmt.Println("Hasil Penjumlahan:", a&b) // Hasil Penjumlahan: 0
	fmt.Println("Hasil Penjumlahan:", a|b) // Hasil Penjumlahan: 30
	fmt.Println("Hasil Penjumlahan:", a^b) // Hasil Penjumlahan: 30
	fmt.Println("Hasil Penjumlahan:", a<<b) // Hasil Penjumlahan: 10485760


	// Augmented Assignment Operators
	// +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i) // Hasil Penjumlahan: 20

	i+= 5
	fmt.Println(i) // Hasil Penjumlahan: 25

	// Unary Operators
	// +, -, ++, --, !
	var j = 10
	j++ // j = j + 1
	fmt.Println(j) // Hasil Penjumlahan: 11

	j-- // j = j - 1
	fmt.Println(j) // Hasil Penjumlahan: 10

	j = -j // j = -10
	fmt.Println(j) // Hasil Penjumlahan: -10

}