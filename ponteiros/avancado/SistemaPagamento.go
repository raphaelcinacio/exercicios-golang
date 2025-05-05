package avancado

import "fmt"

type Pagamento interface {
	Pagar(valor float64)
}

type CartaoCredito struct {
	Titular string
}

type Boleto struct {
	CodigoBarras string
}

func (cartao *CartaoCredito) Pagar(valor float64) {
	fmt.Printf("\nPagando %f com cartão de crédito do titular %s", valor, cartao.Titular)
}

func (boleto *Boleto) Pagar(valor float64) {
	fmt.Printf("\nPagando %f com boleto. Código de barras: %s", valor, boleto.CodigoBarras)
}

func ProcessarPagamento(p Pagamento, valor float64) {
	p.Pagar(valor)
}
