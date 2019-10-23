package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Eko"
	names[1] = "Rudi"
	names[2] = "Joko"
	names[3] = "Joko"
	names[4] = "Joko"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, value := range names {
		// fmt.Println("index", index, "=", value)
		fmt.Println(value)
	}
}
