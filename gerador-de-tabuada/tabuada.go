package tabuada

import (
	"fmt"
	"strconv"
	"strings"

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
	for {
		utils.Cabecalho("Gerador de Tabuada")

		numero := solicitarNumero()
		operacao := utils.SolicitarOperacao()

		exibirTabuada(numero, operacao)

		if !utils.ContinuarOuSair() {
			break
		}

		utils.LimparTerminal()
	}
}

func solicitarNumero() float64 {
	for {
		fmt.Print("Digite um número (inteiro ou decimal): ")
		var entrada string
		fmt.Scanln(&entrada)

		entrada = strings.Replace(entrada, ",", ".", -1)

		numero, err := strconv.ParseFloat(entrada, 64)
		if err == nil {
			return numero
		}

		fmt.Println("Entrada inválida. Digite um número válido (ex: 5 ou 3.14).")
	}
}

func exibirTabuada(numero float64, operacao string) {
	for i := 1; i <= 10; i++ {
		resultado := utils.Calcular(numero, float64(i), operacao)
		fmt.Printf("%s %s %2d = %s\n", utils.FormatarNumero(numero), operacao, i, utils.FormatarNumero(resultado))
	}
}
