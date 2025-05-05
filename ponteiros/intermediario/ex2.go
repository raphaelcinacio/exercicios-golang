package intermediario

type Estoque struct {
	NomeEstoque string
	Quantidade  int
}

type Pedido struct {
	NomePedido string
	Quantidade int
	Efetuado   bool
}

func AtualizarEstoque(estoque *Estoque, pedido *Pedido) {

	if pedido.Quantidade > estoque.Quantidade {
		pedido.Efetuado = false
	} else {
		estoque.Quantidade -= pedido.Quantidade
	}

}
