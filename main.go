package main

import "fmt"

func main() {
	var n, numero, suma int
	fmt.Scanf("%d", &n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numero)
		s = append(s, numero)
	}

	for _, v := range s {
		suma = suma + v
	}

	fmt.Println(suma)
}
