package main
import "fmt"

func main () {
	names := [...]string{
		"nol",
		"satu", 
		"dua", 
		"tiga", 
		"empat", 
		"lima", 
		// "Enam",
	}
	slice1 := names[4:6]
	fmt.Println("Slice1:", slice1) // Output: Slice1: [Alice Bob]

	slice2 := names[:3]
	fmt.Println("Slice2:", slice2) // Output: Slice2: [Alice Bob Charlie]

	slice3 := names[3:]
	fmt.Println("Slice3:", slice3) // Output: Slice3: [David]


	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println("Day Slice 1:", daySlice1) // Output: Day Slice 1: [sabtu minggu]

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println("Day Slice 1:", daySlice1) // Output: Day Slice 1: [sabtu baru minggu baru]
	fmt.Println("Days:", days) // Output: Days: [senin selasa rabu kamis jumat sabtu baru minggu baru]

	fmt.Println("\n===========================\n") 
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "sabtu Lama"
	fmt.Println("Day Slice 1:", daySlice1) // Output: Day Slice 1: [sabtu baru minggu baru]
	fmt.Println("Day Slice 2:", daySlice2) // Output: Day Slice 2: [sabtu baru minggu baru Libur Baru]
	fmt.Println("Days:", days) // Output: Days: [senin selasa rabu kamis jumat sabtu baru minggu baru]
	
	fmt.Println("\n===========================\n") 

	// var newSlice []string = make([]string, 2, 5)
	newSlice := make([]string, 2, 5) // Create a new slice with length 2 and capacity 5
	newSlice[0] = "nol"
	newSlice[1] = "satu"
	// newSlice[2] = "dua" // This will cause a runtime panic: index out of range

	fmt.Println("New Slice:", newSlice) // Output: New Slice: [nol satu]
	fmt.Println("New Slice Length:", len(newSlice)) // Output: New Slice Length: 2
	fmt.Println("New Slice Capacity:", cap(newSlice)) // Output: New Slice Capacity: 5

	newSlice2 := append(newSlice, "dua", "tiga", "empat", "lima")
	fmt.Println("\nNew Slice 2:", newSlice2) // Output: New Slice 2: [nol satu dua tiga empat lima]

	newSlice2[0] = "nol baru"
	fmt.Println("New Slice 2:", newSlice2) // Output: New Slice 2: [nol baru dua tiga empat lima]

	fmt.Println("\n===========================\n") 
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println("From Slice:", fromSlice) // Output: From Slice: [senin selasa rabu kamis jumat sabtu baru minggu baru]
	fmt.Println("To Slice:", toSlice) // Output: To Slice: [senin selasa rabu kamis jumat sabtu baru minggu baru]

	fmt.Println("\n===========================\n") 
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Ini Array:", iniArray) // Output: Ini Array: [1 2 3 4 5]
	fmt.Println("Ini Slice:", iniSlice) // Output: Ini Slice: [1 2 3 4 5]

}