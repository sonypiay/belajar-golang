package main

import "fmt"

func main() {

	name := "John"

	switch name {
	case "John":
		println("Hello John")

	default:
		println("Hello stranger")
	}

	// short statement
	switch age := 15; age < 17 {
	case true:
		fmt.Println("Umur dibawah 17 tahun")
	case false:
		fmt.Println("Umur diatas 17 tahun")
	}
}
