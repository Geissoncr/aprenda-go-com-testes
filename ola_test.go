package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Geisson")
	esperado := "Olaaa, Geisson"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
