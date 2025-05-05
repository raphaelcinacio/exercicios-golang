/*
Verificador de Palíndromo
Verifique se uma palavra é um palíndromo (ex: "arara", "radar").
*/

package iniciante

import (
	"fmt"
	"strings"
)

func ValidarPalindromo() {
	var palavra string = "anilina"
	var novaPalavra string = ""

	i := 1
	for i <= len(palavra){
		novaPalavra += (string(palavra[len(palavra) - i]))
		i++
	}

	if strings.EqualFold(novaPalavra, palavra) {
		fmt.Println("É um palíndromo")
	} else {
		fmt.Println("Não é um palíndromo")
	}

}