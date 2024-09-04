package main

import "fmt"

func main() {
	// names := [...]string{"Sony", "Budi", "Supri", "Eko"}

	// slice1 := names[1:3]
	// fmt.Println(slice1)

	// slice2 := names[:3]
	// fmt.Println(slice2)

	// var slice3 []string = names[1:]
	// fmt.Println(slice3)

	// var slice4 []string = names[:]
	// fmt.Println(slice4)

	days := [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
}