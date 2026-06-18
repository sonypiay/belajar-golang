package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Eko"
	customer.Address = "Khannedy"
	customer.Age = 24

	fmt.Println(customer)

	budi := Customer{
		Name:    "Budi",
		Address: "Jakarta",
		Age:     30,
	}
	fmt.Println(budi)

	joko := Customer{"Joko", "Surabaya", 35}
	fmt.Println(joko)
}
