package main

import "fmt"

func main() {
	// buat map baru dengan inisial
	person := map[string]string{
		"name":  "John Doe",
		"email": "johndoe@example.com",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["email"])

	// buat map baru tanpa inisial
	book := make(map[string]string)
	book["title"] = "Learn Go Programming"
	book["author"] = "John Doe"
	book["wrong"] = "this will be deleted"

	fmt.Println(book)

	// hapus map
	delete(book, "wrong")
	fmt.Println(book)
}
