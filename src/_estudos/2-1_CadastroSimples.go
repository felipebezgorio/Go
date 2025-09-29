package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
	Email string
}

func main() {
	pessoas := []Pessoa{
		{"Mariani", 36, "mariani@email.com"},
		{"Felipe", 34, "felipe@email.com"},
		{"João", 30, "j@jog.com"},
		{"Ana", 25, "a@na.com"},
		{"Pedro", 40, "pedro@ped.com"}}

	for _, p := range pessoas {
		imprimePessoa(p)
	}

	sPessoas := pessoas[0:2]

	fmt.Println("\nCópia\n")

	for _, p := range sPessoas {
		imprimePessoa(p)
	}

}

func imprimePessoa(p Pessoa) {
	fmt.Printf("Nome: %s; Idade: %d; Email: %s\n", p.Nome, p.Idade, p.Email)
}
