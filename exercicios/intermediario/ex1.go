/*
Números primos
Verifique se um número é primo

Explicação:

A partir do momento que ambos os fatores de um produto são maiores que a raiz quadrada do número, o resultado excederá o número que está sendo avaliado.

Por esse motivo, ao verificar se um número é primo, só é necessário validar se ele é divisível por algum valor entre 2 e a raiz quadrada desse número (inclusive).

Caso não haja nenhum divisor nesse intervalo, o número é considerado primo, pois qualquer possível divisor acima da raiz quadrada teria que ser multiplicado por um número menor que ela — o que já teria sido verificado.

*/

package intermediario

import (
	"fmt"
	"math"
)

func VerificarNumeroPrimo(numero int) {

	divisores := []int{1, numero}
	for i := 2; i < int(math.Sqrt(float64(numero)))+1; i++ {
		if numero%i == 0 {
			divisores = append(divisores, i)
		}
	}

	if len(divisores) > 2 {
		fmt.Println("Número não é primo. Divisores:", divisores)
	} else {
		fmt.Println("Número primo")
	}

}
