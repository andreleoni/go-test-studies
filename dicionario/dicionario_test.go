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
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, ErrNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {
	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dictionary{}
		dicionario.Add("teste", "apenas um teste")

		esperado := "apenas um teste"
		resultado, err := dicionario.Busca("teste")

		if err != nil {
			t.Fatal("não foi possivel encontrar a palavra adicionada")
		}

		comparaStrings(t, esperado, resultado)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dictionary{palavra: definicao}
		err := dicionario.Add(palavra, "teste novo")

		comparaErro(t, err, ErrPalavraExistente)
	})
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "test")
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}
