package soma

func Soma(numeros []int) int {
	soma := 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTudo(slices ...[]int) (somas []int) {

	for _, numeros := range slices {
		somas = append(somas, Soma(numeros))
	}
	return somas
}

func SomaTodoOResto(slices ...[]int) (somas []int) {

	for _, numeros := range slices {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			resto := numeros[1:]
			somas = append(somas, Soma(resto))
		}
	}
	return somas
}
