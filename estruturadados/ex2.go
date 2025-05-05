package estruturadados

import "fmt"

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func CriarProduto(nome string, preco float64) *Produto {

	produto := Produto{
		Nome:  nome,
		Preco: preco,
	}

	return &produto
}

func (produto *Produto) AtualizarPreco(preco float64) {
	produto.Preco = preco
}

func (produto *Produto) RemoverEstoque(quantidade int) {

	if produto.Quantidade < quantidade {
		fmt.Println("Não é possível remover")
	} else {
		produto.Quantidade -= quantidade
	}
}

func (produto *Produto) AdicionarEstoque(quantidade int) {
	produto.Quantidade += quantidade
}

func ListarProdutos(produtos map[int]*Produto) {

	for k, v := range produtos {
		fmt.Printf("%d - Nome: %s - Preço: %f - Quantidade: %d\n", k, v.Nome, v.Preco, v.Quantidade)
	}

}
