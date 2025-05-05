/*
Tabuada
Gere a tabuada de um número de 1 a 10.
*/

package iniciante

import "fmt"

func Tabuada() {

	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, i*numero)
	}
}
