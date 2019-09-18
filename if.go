package main

import "fmt"

func main() {

	// if i % 3 = Foo
	// if i % 5 = Bar
	// if i % 3 && i % 5 = FooBar

	for i := 1; i <= 30; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FooBar")

		} else if i%3 == 0 {
			fmt.Println("Foo")

		} else if i%5 == 0 {
			fmt.Println("Bar")

		} else {
			fmt.Println(i)
		}

	}

}
