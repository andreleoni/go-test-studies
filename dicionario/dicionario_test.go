package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := map[string]string{"test": "isso é apenas um teste"}

	resultado := Busca(dicionario, "test")
	esperado := "isso é apenas um teste"

	comparaStrings(t, resultado, esperado)
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "test")
	}
}
