package calculadora

import (
	"fmt"
	"strings"
)

/*
	MELHORIAS:
	1 - Considerar que o 5 entra em um fluxo de switchcase mas não emite as mensagens para inserir os valores
	2 - Ao adicionar uma opção inválida, tem que sair do fluxo de execução. Exibir que a opção não existe e sair do fluxo
	3 - Considerar valores float, dado que existe a divisão
	4 - Caso coloque letras, tem que lidar com o erro
	5 - Pensar em cada uma das funcionalidades como funcionalidades. Se a minha calculadora tem a funcionalidade de somar, logo, terei um func somar;
	6 - PLUS - Pensar em testes unitários para cada uma das funcionalidades
	7 - PLUS - Documentar cada uma das funcionalidades
*/

func Executar() {

	var opcao, i, j int

	cabecalho("Calculadora em GO")

	menuEscolhas()
	fmt.Scan(&opcao)

	if opcao < 1 || opcao >= 5 {
		cabecalho("Opção inválida")
		return
	}

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&i)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&j)

	switch opcao {
	case 1:
		cabecalho("Resultado da Soma")
		fmt.Println(i + j)

	case 2:
		cabecalho("Resultado da Subtração")
		fmt.Println(i - j)
	case 3:
		cabecalho("Resultado da Multiplicação")
		fmt.Println(i * j)
	case 4:
		cabecalho("Resultado da Divisão")
		fmt.Println(i / j)
	case 5:
		cabecalho("Sair")
		return
	default:
		cabecalho("Opção inválida")
		return
	}

}

func cabecalho(palavraCabecalho string) {
	tamanho := len(palavraCabecalho) + 6
	linha := strings.Repeat("-", tamanho)

	fmt.Println(linha)
	fmt.Println("| ", palavraCabecalho, " |")
	fmt.Println(linha)
}

func menuEscolhas() {
	fmt.Println("Escolha a operação: ")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("5 - Sair")
}
