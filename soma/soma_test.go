package soma

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("coleção de 5 numeros", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("Resultado esperado '%d', veio '%d' entrada '%v'", esperado, resultado, numeros)
		}
	})
	t.Run("Coleção com qq tamanho de numero", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("Resultado esperado '%d', veio '%d' entrada '%v'", esperado, resultado, numeros)
		}
	})
}
func TestSomaTudo(t *testing.T) {
	t.Run("Coleção de slices", func(t *testing.T) {

		resultado := SomaTudo([]int{1, 2, 3}, []int{1, 2})
		esperado := []int{6, 3}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Resultado esperado '%d', veio '%d' entrada", esperado, resultado)
		}
	})
}

func TestSomaTodoOResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("Soma o resto dos slices tirando a primeira posicao", func(t *testing.T) {

		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSomas(t, resultado, esperado)

	})
	t.Run("Soma o resto dos slices tirando a primeira posicao - 1 slice vazio", func(t *testing.T) {

		resultado := SomaTodoOResto([]int{}, []int{0, 9})
		esperado := []int{0, 9}

		verificaSomas(t, resultado, esperado)
	})
}
