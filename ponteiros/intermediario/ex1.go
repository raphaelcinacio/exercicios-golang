package intermediario

func AtualizarNotas(notas []int, valorExtra *int) {

	for i := 0; i < len(notas); i++{
		notas[i] += *valorExtra
	}
}