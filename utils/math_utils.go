package utils

import (
	"fmt"
	"strings"
)

func SolicitarOperacao() string {
	operacoesValidas := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	for {
		LimparEntrada()
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

func Calcular(a, b float64, operacao string) float64 {
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

func FormatarNumero(numero float64) string {
	str := fmt.Sprintf("%.2f", numero)
	str = strings.Replace(str, ".", ",", -1)

	if strings.HasSuffix(str, ",00") {
		str = strings.TrimSuffix(str, ",00")
	}

	return str
}
