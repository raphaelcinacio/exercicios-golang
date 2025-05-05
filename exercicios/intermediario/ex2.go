/*
Sequência de Fibonacci
Imprima os n primeiros termos da sequência de Fibonacci.
*/

package intermediario

import "fmt"

func CalcularSequenciaFibonacci(numero int) {
	sequenciaFibonacci := []int{0, 1} // [1, 1, 2, 3]

	for i := 0; i < numero - 2; i++ {
		numeroAnterior := sequenciaFibonacci[i]  // 1, 1
		proximoNumero := sequenciaFibonacci[i+1] // 1, 2
		sequenciaFibonacci = append(sequenciaFibonacci, (numeroAnterior + proximoNumero))
	}

	fmt.Println(sequenciaFibonacci)
}