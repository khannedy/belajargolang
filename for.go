package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "Hello World")
	}

	fmt.Println("Selesai")

	for i := 10; i >= 1; i-- {
		fmt.Println(i, "Hello World")
	}
}
