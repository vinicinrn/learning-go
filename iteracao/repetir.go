package iteracao

const quantidadeRepeticoes = 5

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere //OPERADOR += -> adicionar e atribuir
	}
	return repeticoes
}
