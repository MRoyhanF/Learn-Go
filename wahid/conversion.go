package main
import "fmt"

// int8   -128 to 127
// int16 -32768 to 32767
// int32 -2147483648 to 2147483647
// int64 -9223372036854775808 to 9223372036854775807
// uint8  0 to 255
// uint16 0 to 65535
// uint32 0 to 4294967295
// uint64 0 to 18446744073709551615
// float32 1.401298464324817e-45 to 3.4028234663852886e+38
// float64 5.0000000000000000e-324 to 1.7976931348623157e+308

func main()  {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32) // 32768
	fmt.Println(nilai64) // 32768
	fmt.Println(nilai16) // -32768

	var name = "Royhan Fadhli"
	var e = name[0] // 82
	var eString = string(e) // R

	fmt.Println(name) // Royhan Fadhli
	fmt.Println(e) // 82
	fmt.Println(eString) // R
}