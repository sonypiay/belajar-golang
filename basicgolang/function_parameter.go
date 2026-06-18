package main

import "fmt"

type Filter func(string) string

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func sayHelloWithFilter(name string, filter Filter) {
	filtered := filter(name)
	fmt.Println("Hello ", filtered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "****"
	}

	return name
}

func main() {
	sayHello("Eko", "Kurniawan")
	sayHelloWithFilter("Anjing", spamFilter)
}
