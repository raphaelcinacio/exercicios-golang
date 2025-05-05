// O Go, internamente, cria um ponteiro para a instância de aluno e o passa para o método, como se fosse:
// (&aluno).AdicionarNota(8.0)

package avancado

import "fmt"

type Aluno struct {
	Nome  string
	notas []float64
}

func (aluno *Aluno) AdicionarNota(nota float64) {
	aluno.notas = append(aluno.notas, nota)
}

func (aluno *Aluno) CalcularMedia() float64 {
	if len(aluno.notas) != 0 {

		soma := 0.0
		for _, nota := range aluno.notas {
			soma += nota
		}
		return soma / float64(len(aluno.notas))
	}

	return 0.0
}

func (aluno *Aluno) MostrarNotas() {
	fmt.Println("\n=====NOTAS=====")
	for _, nota := range aluno.notas {
		fmt.Println(nota)
	}
}

func (aluno *Aluno) Aprovar() bool {
	return aluno.CalcularMedia() >= 7
}
