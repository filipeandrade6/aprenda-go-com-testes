package iteracao

const quatidadeRepeticoes = 5

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < quatidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
