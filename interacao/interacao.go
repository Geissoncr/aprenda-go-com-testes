package interacao

func Repetir(caractere string, vezes int) string {
	var repeticoes string
	for i := 0; i < vezes; i++ {
		repeticoes = repeticoes + caractere
	}
	return repeticoes
}
