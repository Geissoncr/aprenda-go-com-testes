package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("Diz olá para qualquer Pessoa.", func(t *testing.T) {
		resultado := Ola("Geisson")
		esperado := "Olá, Geisson"
		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("Diz Olá, mundo quando não manda parametro nenhum", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

}
