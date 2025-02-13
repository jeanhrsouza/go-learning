package imparoupar

import (
	"fmt"

	"estudos-brabos/utils"
)

/*
	Bugzillas:
	1. Ao adicionar letras nos números, deve ter tratamento ✅
	2. Apertar enter sem digitar nada, deve ter tratamento ✅
	3. Ao digitar "s" ou "n" para continuar, deve ser case insensitive ✅

	Melhorias:
	1. Adicionar tratamento para números negativos ✅
	2. Adicionar tratamento para números decimais ✅
	3. Diminuir a quantidade de if e else. Evitar ao máximo o uso de else ✅
	4. Fazer a lógica de continuar ou sair totalmente na função continuarOuSair ✅
*/

func Executar() {

	utils.ExecutarComLoop("Ímpar ou Par", func() {
		numero := utils.SolicitarNumeroInteiro("Digite um número inteiro: ")
		fmt.Printf("O número %d é %s.\n", numero, identificarImparOuPar(numero))
	})
}

func identificarImparOuPar(numero int) string {
	if numero%2 == 0 {
		return "par"
	}
	return "ímpar"

}
