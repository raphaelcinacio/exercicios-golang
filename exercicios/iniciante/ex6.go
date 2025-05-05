/*
Crie um struct Funcionario com nome e salário. Implemente um método que calcula o salário com bônus (10% a mais).
*/

package iniciante

var MensagensErro = map[string]int{
	"Assistente": 1,
	"Analista": 2,
	"Lead": 3,
	"Gerente": 4,
}

type Funcionario struct {
	Nome string
	Salario float64
	Cargo int
}

func (f *Funcionario) SalarioComBonus() float64 {

	if f.Cargo <= 2 {
		return f.Salario * 1.10
	} else if f.Cargo == 3{
		return f.Salario * 1.70
	}
	return f.Salario * 2.0
}