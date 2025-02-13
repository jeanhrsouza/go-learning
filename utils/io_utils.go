package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ExecutarComLoop(titulo string, acao func()) {
	LimparEntrada()
	for {
		Cabecalho(titulo)
		acao()

		if !ContinuarOuSair() {
			break
		}
		LimparTerminal()
	}
}

// SolicitarNumeroInteiro é uma função que pede ao usuário para digitar um número inteiro, trata os erros e retorna o número digitado.
func SolicitarNumeroInteiro(mensagem string) int {

	for {
		fmt.Print(mensagem)
		var entrada string
		fmt.Scanln(&entrada)

		entrada = strings.TrimSpace(entrada)

		if entrada == "" {
			fmt.Println("Entrada inválida. Digite um número válido.")
			continue
		}

		numero, err := strconv.Atoi(entrada)
		if err == nil {
			return numero
		}

		fmt.Println("Entrada inválida. Digite um número válido.")

	}
}

func SolicitarNumeroDecimal(mensagem string) float64 {
	for {
		fmt.Print(mensagem)
		var entrada string
		fmt.Scanln(&entrada)

		entrada = strings.Replace(entrada, ",", ".", -1)

		numero, err := strconv.ParseFloat(entrada, 64)
		if err == nil {
			return numero
		}

		fmt.Println("Entrada inválida. Digite um número válido.")

	}
}

func ContinuarOuSair() bool {
	for {
		Cabecalho("Deseja continuar? (s/n)")

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

// LimparEntrada é uma função que remove qualquer caractere residual do buffer de entrada
// do teclado, evitando que a próxima leitura de dados seja afetada por esses caracteres.
func LimparEntrada() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

// LimparTerminal limpa a tela do terminal
func LimparTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
