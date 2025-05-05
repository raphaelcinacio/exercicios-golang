/*
Validador de CPF (sem usar bibliotecas)
Crie uma função que valide se um CPF é válido.
*/

package avancado

import (
	"fmt"
	"strings"
)

func ValidarCPF(documento string) bool {

	fmt.Println(documento)

	documentoFormatado := strings.ReplaceAll(strings.ReplaceAll(documento, "-", ""), ".", "")
	
	if primeiroDigitoVerificador(documentoFormatado) != int(documentoFormatado[9] - '0') || segundoDigitoVerificador(documentoFormatado) != int(documentoFormatado[10] - '0') {
		return false
	}
	return true
}

func primeiroDigitoVerificador(documentoFormatado string) int {

	soma := 0
	var digitoVerificador int = 0
	for i, j := 0, 10; i <= 8; i, j = i+1, j-1 {
		
		soma += int(documentoFormatado[i] - '0') * j
	}

	calculo := 11 - (soma % 11)

	if calculo < 10{
		digitoVerificador = calculo
	}
	
	return digitoVerificador
}

func segundoDigitoVerificador(documentoFormatado string) int {

	soma := 0
	var digitoVerificador int = 0
	for i, j := 0, 11; i <= 9; i, j = i+1, j-1 {
		
		soma += int(documentoFormatado[i] - '0') * j
	} 

	calculo := 11 - (soma % 11)

	if calculo < 10{
		digitoVerificador = calculo
	}

	return digitoVerificador
}