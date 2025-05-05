/*
Crie uma função chamada Somar que receba dois ponteiros para int, some os valores apontados por eles e retorne o resultado.
*/

package iniciante


func Somar(a, b *int) int {
	valorA := *a
	valorB := *b
	return valorA + valorB
}