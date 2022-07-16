package soma

import "testing"

func TestSoma(t *testing.T) {
	numeros := [5]int{1, 2, 3, 4, 5}

	resultado := Soma(numeros)
	esperado := 15

	if esperado != resultado {
		t.Errorf("Resultado esperado '%d', veio '%d' entrada '%v'", esperado, resultado, numeros)
	}
}
