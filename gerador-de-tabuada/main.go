package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

	var numero int

	for {

		cabecalho("Gerador de Tabuada")

		fmt.Print("Digite um número: ")
		fmt.Scan(&numero)

		for i := 1; i <= 10; i++ {
			fmt.Println(numero, " x ", i, " = ", numero*i)
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
	cabecalho("Deseja continuar? (s/n)")
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
