package main

import (
	"fmt"
)

type Produto struct {
	Nome       string
	Quantidade int
	Preco      float64
}

func main() {
	estoque := make(map[string]Produto)

	fmt.Println("Inserindo Produtos...")

	addProduto(estoque, "1", "Computador", 10, 1250.75)
	addProduto(estoque, "2", "Monitor", 20, 550.99)
	addProduto(estoque, "3", "Mouse", 10, 31.89)
	addProduto(estoque, "4", "Teclado", 10, 250.0)
	addProduto(estoque, "5", "Microfone", 5, 129.90)

	imprimeTodosProdutos(estoque)

	fmt.Println("\nRemovendo Produtos")
	delProduto(estoque, "2")
	delProduto(estoque, "4")

	imprimeTodosProdutos(estoque)

	fmt.Println("\nBuscando Produto código 15 ...")

	if prod, ok := buscaProduto(estoque, "15"); ok {
		imprimeProduto(prod)
	} else {
		fmt.Println("Produto não encontrado.")
	}

}

func imprimeTodosProdutos(estoque map[string]Produto) {
	for _, p := range estoque {
		imprimeProduto(p)
	}
}

func imprimeProduto(prod Produto) {
	fmt.Printf("\nNome: %s\nQuantidade: %d\nPreço: %.2f\n", prod.Nome, prod.Quantidade, prod.Preco)
}

func addProduto(estoque map[string]Produto,
	codigo string,
	produto string,
	quantidade int,
	preco float64) {

	estoque[codigo] = Produto{Nome: produto, Quantidade: quantidade, Preco: preco}
}

func delProduto(estoque map[string]Produto, chave string) {
	delete(estoque, chave)
}

func buscaProduto(estoque map[string]Produto, chave string) (Produto, bool) {
	produto, ok := estoque[chave]
	return produto, ok
}
