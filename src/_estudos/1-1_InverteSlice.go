package main

import "fmt"

func inverter(slice []int) []int {
	r := []int{}
	for i := len(slice) - 1; i >= 0; i-- {
		r = append(r, slice[i])
	}
	return r
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array", a)

	s := a[3:8]
	fmt.Println("Fatia", s)

	r := inverter(s)
	fmt.Println("Fatia Invertida", r)
}
