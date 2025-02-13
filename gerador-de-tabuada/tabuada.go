package tabuada

import (
	"fmt"

	"estudos-brabos/utils"
)

/*
	Bugs:
	1. Realizar tratamentos para letras ✅
	2. Realizar tratamentos para números decimais. Deve aceitar números decimais e inteiros ✅

	Melhorias:
	1. Fazer todas as operações matemáticas (mais, menos, vezes e dividido) ✅


*/

func Executar() {

	utils.ExecutarComLoop("Gerado de Tabuada", func() {
		numero := utils.SolicitarNumeroDecimal("Digite um número (inteiro ou decimal): ")
		operacao := utils.SolicitarOperacao()

		exibirTabuada(numero, operacao)
	})
}

func exibirTabuada(numero float64, operacao string) {
	for i := 1; i <= 10; i++ {
		resultado := utils.Calcular(numero, float64(i), operacao)
		fmt.Printf("%s %s %2d = %s\n", utils.FormatarNumero(numero), operacao, i, utils.FormatarNumero(resultado))
	}
}
