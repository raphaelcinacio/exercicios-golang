/*
Dado um slice de inteiros, crie uma função que retorne a soma apenas dos números ímpares.
*/

package iniciante

func SomarImpares(numeros []int) int {

	soma := 0
	for i := 0; i < len(numeros); i++{
		if numeros[i] % 2 != 0{
			soma += numeros[i]
		}
	}
	return soma
}