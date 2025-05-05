/*
Gerador de senhas aleatórias
Crie uma função que gere senhas aleatórias com letras, números e símbolos.
*/

package avancado

import (
	"math/rand"
)

func GerarSenha(tamanho int) string {
	// Defina o conjunto de caracteres possíveis
	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*"

	// Gerar a senha
	var senha string
	for i := 0; i < tamanho; i++ {
		senha += string(caracteres[rand.Intn(len(caracteres))])
	}

	return senha
}
