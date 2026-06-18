package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}

func main() {
	sayHello("Sony", "Darmawan")
}
