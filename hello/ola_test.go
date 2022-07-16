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
		resultado := Ola("Geisson", "")
		esperado := "Olá, Geisson"
		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("Diz Olá, mundo quando não manda parametro nenhum", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Teste passando outro idioma - Espanhol", func(t *testing.T) {
		resultado := Ola("Jack", "espanhol")
		esperado := "Hola, Jack"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Teste passando outro idioma - Frances", func(t *testing.T) {
		resultado := Ola("Jack", "frances")
		esperado := "Bonjour, Jack"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Teste passando outro idioma - Ingles", func(t *testing.T) {
		resultado := Ola("Jack", "ingles")
		esperado := "Hi, Jack"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
