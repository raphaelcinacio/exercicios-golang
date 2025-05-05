/*
Contar ocorrências de palavras
Crie um map[string]int que conta quantas vezes cada palavra aparece em um slice de strings.
*/

package estruturadados

func RetornaMap(palavras []string) map[string]int {
	
	palavrasMap := make(map[string]int)

	for _, valor := range palavras {
		palavrasMap[valor]++
	}

	return palavrasMap
}