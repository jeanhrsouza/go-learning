package tabuada

import (
	"estudos-brabos/utils"
	"testing"
)

func TestCalcular(t *testing.T) {
	testes := []struct {
		a, b     float64
		operacao string
		esperado float64
	}{
		{5, 3, "+", 8},
		{5, 3, "-", 2},
		{5, 3, "*", 15},
		{6, 3, "/", 2},
		{10, 0, "/", 0},
		{0, 0, "/", 0},
	}

	for _, teste := range testes {
		resultado := utils.Calcular(teste.a, teste.b, teste.operacao)
		if resultado != teste.esperado {
			t.Errorf("Erro: calcular (%v, %v, %v) retornou %v, esperado %v", teste.a, teste.b, teste.operacao, resultado, teste.esperado)
		}
	}

}
