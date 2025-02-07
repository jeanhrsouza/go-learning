package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

/*
	Bugzillas:
	1. Ao adicionar letras nos números, deve ter tratamento
	2. Apertar enter sem digitar nada, deve ter tratamento
	3. Ao digitar "s" ou "n" para continuar, deve ser case insensitive

	Melhorias:
	1. Adicionar tratamento para números negativos
	2. Adicionar tratamento para números decimais
	3. Diminuir a quantidade de if e else. Evitar ao máximo o uso de else
	4. Fazer a lógica de continuar ou sair totalmente na função continuarOuSair
*/

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
