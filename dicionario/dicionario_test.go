package main

import "testing"

func TestAdiciona(t *testing.T) {
	t.Run("Palavra Nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "isso é apenas um teste"

		err := dicionario.Adiciona(palavra, definicao)
		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)

	})
	t.Run("Palavra Existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "tete novo")
		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("palavra existente", func(t *testing.T) {

		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		novaDefinicao := "Nova Definição"

		err := dicionario.Atualiza(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})
	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)

		comparaErro(t, err, ErrPalavraInexistente)
	})

}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definição de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != ErrNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado: ", palavra)
	}
}

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

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("deveria ter encontrado palavra adicionada:", err)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s',  esperado '%s'", resultado, definicao)
	}
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
