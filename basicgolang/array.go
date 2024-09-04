package main

import "fmt"

func main() {
	var data[3]string

	data[0] = "Sony"
	data[1] = "John"
	data[2] = "Doe"

	fmt.Println("Data: ", data)
	fmt.Println("Panjang: ", len(data))
}