package calculadora

import (
	"fmt"

	"estudos-brabos/utils"
)

/*
	MELHORIAS:
	1 - Considerar que o 5 entra em um fluxo de switchcase mas não emite as mensagens para inserir os valores ✅
	2 - Ao adicionar uma opção inválida, tem que sair do fluxo de execução. Exibir que a opção não existe e sair do fluxo de execução ✅
	3 - Considerar valores float, dado que existe a divisão ✅
	4 - Caso coloque letras, tem que lidar com o erro ✅
	5 - Pensar em cada uma das funcionalidades como funcionalidades. Se a minha calculadora tem a funcionalidade de somar, logo, terei um func somar; ✅
	6 - PLUS - Pensar em testes unitários para cada uma das funcionalidades
	7 - PLUS - Documentar cada uma das funcionalidades
*/

func Executar() {

	utils.Cabecalho("Calculadora em GO")

	operacao := utils.SolicitarOperacao()

	primeiroValor := utils.SolicitarNumeroDecimal("Digite o primeiro número: ")
	segundoValor := utils.SolicitarNumeroDecimal("Digite o segundo número: ")

	resultadoCalculo := utils.Calcular(primeiroValor, segundoValor, operacao)

	fmt.Println("Resultado: ", utils.FormatarNumero(resultadoCalculo))
}
