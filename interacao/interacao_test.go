package interacao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("Era esperado '%s' mas veio '%s'", esperado, repeticoes)
	}
}

func TestRepetirComValorDeRepeticao(t *testing.T) {
	repeticoes := Repetir("a", 3)
	esperado := "aaa"

	if repeticoes != esperado {
		t.Errorf("Era esperado '%s' mas veio '%s'", esperado, repeticoes)
	}
}
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 0)
	}
}
