package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := map[string]string{"test": "isso é apenas um teste"}

	resultado := Busca(dicionario, "test")
	esperado := "isso é apenas um teste"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "test")
	}
}
