package main

import "fmt"

func main() {
	lista := []int{}

	for i := 0; i < 10; i++ {
		lista = append(lista, i)
	}

	fmt.Println(filtraPares(lista))
}

func filtraPares(numeros []int) []int {
	var pares []int

	for _, x := range numeros {
		if x%2 == 0 {
			pares = append(pares, x)
		}
	}
	return pares
}
