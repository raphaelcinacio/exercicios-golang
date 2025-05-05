/*
Contador de Vogais
Peça uma string e conte quantas vogais ela tem.
*/

package iniciante

import (
	"fmt"
	"strings"
)

func ContadorVogais(){
	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	var quantidade int = 0
	i := 0
	for i < len(palavra) {
		if strings.ContainsAny(string(palavra[i]), "aeiou"){
			quantidade += 1
		}
		i++
	}

	fmt.Printf("Quantidade de vogais: %d", quantidade)



}