package main

import "fmt"

func main() {
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Eko"
	mahasiswa["1002"] = "Kurniawan"
	mahasiswa["1003"] = "Khannedy"

	fmt.Println(mahasiswa)

	fmt.Println(mahasiswa["1001"])
	fmt.Println(mahasiswa["1002"])
	fmt.Println(mahasiswa["1003"])

	for nim, name := range mahasiswa {
		fmt.Println("Nim", nim, "Bernama", name)
	}
}
