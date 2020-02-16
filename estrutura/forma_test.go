package estrutura

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %2.f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	testesArea := []struct {
		forma    Forma
		esperado float64
	}{
		{Retangulo{Largura: 12, Altura: 6}, 72.0},
		{Circulo{Raio: 10}, 314.1592653589793},
		{Triangulo{Base: 12, Altura: 6}, 36.0},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
		}
	}
}
