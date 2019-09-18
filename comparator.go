package main

import "fmt"

func main() {
	a := 1
	b := 2

	// result := a == b // sama dengan
	// result := a != b // tidak sama dengan
	// result := a < b // kurang dari
	// result := a > b // lebih dari
	// result := a <= b // kurang dari atau sama dengan
	result := a >= b // lebih dari atau sama dengan

	fmt.Println(result)
}
