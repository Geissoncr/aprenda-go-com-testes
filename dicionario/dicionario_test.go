package main

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}
	t.Run("Palavra Conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"

		compararStrings(t, resultado, esperado)
	})
	t.Run("Palavra Desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("Desconhecida")

		comparaErro(t, err, ErrNaoEncontrado)

	})
}
func compararStrings(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado '%s', esperado'%s', dado '%s'!", resultado, esperado, "teste")
	}

}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("Resultado '%s', esperado'%s'", resultado, esperado)
	}

}
