package main

import "fmt"

func main() {
	fmt.Println("Brincadeira Pin Pan")
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("Pin Pan (%d)\n", i)
		} else if i%3 == 0 {
			fmt.Printf("Pin (%d)\n", i)
		} else if i%5 == 0 {
			fmt.Printf("Pan (%d)\n", i)
		}

	}
}
