package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	message := recover()
	fmt.Println("Error :", message)
}

func runApp(error bool) {
	defer endApp()

	fmt.Println("App running")

	if error {
		panic("Application Error")
	}

}

func main() {
	runApp(true)
}
