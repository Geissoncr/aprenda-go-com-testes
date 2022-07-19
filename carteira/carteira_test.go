package carteira

import (
	"fmt"
	"testing"
)

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)

	resultado := carteira.Saldo()
	esperado := 10

	fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)

	if resultado != esperado {
		t.Errorf("resultado %d esperado %d", resultado, esperado)
	}
}
