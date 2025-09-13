package main

import "fmt"

var tempEbulicaoK int = 373

func main() {
	tempEbulicaoC := tempEbulicaoK - 273
	fmt.Printf("A temperatura de ebulição da água em °C é de %d", tempEbulicaoC)
}
