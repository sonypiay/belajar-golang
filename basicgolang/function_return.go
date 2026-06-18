package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func getFullname(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "Middle"
	lastName = "Doe"

	return firstName, middleName, lastName
}

func main() {
	// fmt.Println(getHello("John Doe"))

	// return multiple values
	// firstName, lastName := getFullname("John", "Doe")
	// fmt.Println(firstName, lastName)

	// // return only one values
	// result1, _ := getFullname("John", "Doe")
	// fmt.Println(result1)

	a, b, c := getCompleteName()

	fmt.Println(a, b, c)
}
