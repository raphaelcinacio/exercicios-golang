package iniciante

func AumentarPorcentagem(valor *float64, porcentagem float64) {
	*valor = *valor * (porcentagem / 100 + 1)
}