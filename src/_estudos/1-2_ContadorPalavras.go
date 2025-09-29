package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `
A lua, a lua, a lua vem.
A lua, a lua, a lua vai.
A lua vai e a lua vem.
A lua vem e a lua vai.`

	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ToLower(s)

	fmt.Println(s)

	palavras := strings.Fields(s)

	contador := make(map[string]int)

	for i := 0; i <= len(palavras)-1; i++ {
		contador[palavras[i]]++
	}

	for palavra, qtd := range contador {
		fmt.Printf("%s = %d\n", palavra, qtd)
	}

}
