package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/*
	Bugs:
	1. Realizar tratamentos para letras
	2. Realizar tratamentos para números decimais. Deve aceitar números decimais e inteiros

	Melhorias:
	1. Fazer todas as operações matemáticas (mais, menos, vezes e dividido)


*/

func main() {
	for {
		cabecalho("Gerador de Tabuada")

		numero := solicitarNumero()
		operacao := solicitarOperacao()

		exibirTabuada(numero, operacao)

		if !continuarOuSair() {
			break
		}

		limparTerminal()
	}
}

func cabecalho(palavraCabecalho string) {
	tamanho := len(palavraCabecalho) + 6
	linha := strings.Repeat("-", tamanho)

	fmt.Println(linha)
	fmt.Printf("| %s |\n", palavraCabecalho)
	fmt.Println(linha)
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

func solicitarOperacao() string {
	operacoesValidas := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	for {
		fmt.Print("Digite a operação (+, -, * ou /): ")
		var operacao string
		fmt.Scanln(&operacao)

		operacao = strings.TrimSpace(operacao)

		if operacoesValidas[operacao] {
			return operacao
		}

		fmt.Println("Operação inválida. Digite uma operação válida (+, -, * ou /).")

	}
}

func exibirTabuada(numero float64, operacao string) {
	for i := 1; i <= 10; i++ {
		resultado := calcular(numero, float64(i), operacao)
		fmt.Printf("%s %s %2d = %s\n", formatarNumero(numero), operacao, i, formatarNumero(resultado))
	}
}

func calcular(a, b float64, operacao string) float64 {
	switch operacao {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Não é possível dividir por zero.")
			return 0
		}
		return a / b
	default:
		return 0
	}
}

func formatarNumero(numero float64) string {
	str := fmt.Sprintf("%.2f", numero)
	str = strings.Replace(str, ".", ",", -1)

	if strings.HasSuffix(str, ",00") {
		str = strings.TrimSuffix(str, ",00")
	}

	return str
}

func continuarOuSair() bool {
	for {
		cabecalho("Deseja continuar? (s/n)")

		var continuar string
		fmt.Print("> ")
		fmt.Scanln(&continuar)
		continuar = strings.TrimSpace(strings.ToLower(continuar))

		if continuar == "s" {
			return true
		}
		if continuar == "n" {
			return false
		}

		fmt.Println("Opção inválida. Digite 's' para sim ou 'n' para não.")
	}
}

func limparTerminal() {
	switch runtime.GOOS {
	case "windows":
		exec.Command("cmd", "/c", "cls").Run()
	default:
		exec.Command("clear").Run()
	}
}
