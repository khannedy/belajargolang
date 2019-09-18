package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		//if i % 2 == 0 {
		//	continue
		//}

		fmt.Println(i)

		if i == 50 {
			break
		}
	}

}
