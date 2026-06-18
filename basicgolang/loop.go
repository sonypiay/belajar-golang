package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Perulangan ke", i)
	// }

	// names := []string{"Eko", "Kurniawan", "Khannedy"}

	// for index, name := range names {
	// 	fmt.Println("Index", index, "=", name)
	// }

	// for _, name := range names {
	// 	fmt.Println(name)
	// }

	// break
	fmt.Println("Break")
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			break
		}

		fmt.Println("Perulangan ke", i)
	}

	// continue
	fmt.Println("Continue")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", i)
	}
}
