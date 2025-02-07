package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	var numero int

	for {

		cabecalho("Ímpar ou Par")

		fmt.Print("Digite um número: ")
		fmt.Scan(&numero)

		if numero%2 == 0 {
			fmt.Println("O número", numero, "é par")
		} else {
			fmt.Println("O número", numero, "é ímpar")
		}

		if continuarOuSair() == "n" {
			break
		} else {
			limparTerminal()
		}
	}

}

func cabecalho(palavraCabecalho string) {
	tamanho := len(palavraCabecalho) + 6
	linha := strings.Repeat("-", tamanho)

	fmt.Println(linha)
	fmt.Println("| ", palavraCabecalho, " |")
	fmt.Println(linha)
}

func continuarOuSair() string {
	fmt.Println("\n\n Deseja continuar? (s/n)")
	var continuar string
	fmt.Scan(&continuar)

	if continuar != "s" && continuar != "n" {
		fmt.Println("Opção inválida")
		return continuarOuSair()
	}

	return continuar
}

func limparTerminal() {
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
