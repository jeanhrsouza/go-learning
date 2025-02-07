package main

import (
	"fmt"
)

// Fórmula do Juros Simples: M = P x (1 + i x t)

// Onde M é o montante final (Valor total acumulado)
// P é o principal (capital inicial)
// i é a taxa de juros (em decimal, por exemplo, 0.05 para 5%)
// t é o tempo em anos

var (
	M float32
	P float32
	i float32
	t uint16
)

func main() {
	// fmt.Scan(&P)
	P = 1000
	i = 0.10
	t = 2

	M = P * (1 + i*float32(t))
	fmt.Printf("O valor do montate é: %2.f", M)

}
