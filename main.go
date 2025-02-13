package main

import (
	"fmt"
	"os"

	"estudos-brabos/basicao"
	"estudos-brabos/calculadora"
	"estudos-brabos/fundamentos"
	tabuada "estudos-brabos/gerador-de-tabuada"
	imparoupar "estudos-brabos/impar-ou-par"
	"estudos-brabos/utils"
)

func main() {

	utils.CabecalhoASCII("Estudos dos Brabos")

	menu := map[string]func(){
		"1": basicao.Executar,
		"2": calculadora.Executar,
		"3": fundamentos.Executar,
		"4": tabuada.Executar,
		"5": imparoupar.Executar,
	}

	fmt.Printf("\nEscolha um projeto para executar:\n")
	fmt.Println("1 - Básico")
	fmt.Println("2 - Calculadora")
	fmt.Println("3 - Fundamentos")
	fmt.Println("4 - Gerador de Tabuada")
	fmt.Println("5 - Ímpar ou Par")
	fmt.Println("0 - Sair")

	var escolha string
	fmt.Print(">")
	fmt.Scan(&escolha)

	if escolha == "0" {
		fmt.Println("Saindo...")
		os.Exit(0)
	}

	if executar, ok := menu[escolha]; ok {
		executar()
	} else {
		fmt.Println("Opção inválida")
	}

}
