package iniciante

import "fmt"

func ValidarParOuImpar(numero int) {

	if numero % 2 == 0 {
		fmt.Println("Número par")
	} else {
		fmt.Println("Número ímpar")
	}

}