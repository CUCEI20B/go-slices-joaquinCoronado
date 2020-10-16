package main

import "fmt"

func main() {
	var n, numero, suma int
	fmt.Scanln(&n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&numero)
		s = append(s, numero)
	}

	for _, v := range s {
		suma = suma + v
	}

	fmt.Println(suma)
}
