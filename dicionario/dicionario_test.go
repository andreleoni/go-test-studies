package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dictionary{"test": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("test")
		esperado := "isso é apenas um teste"

		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecida")

		if err == nil {
			t.Fatal("é esperado que um erro seja obtido.")
		}
	})
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "test")
	}
}
