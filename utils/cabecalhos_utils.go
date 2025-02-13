package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// Cabecalho é uma função que imprime um cabeçalho com o título dentro de uma borda
func Cabecalho(titulo string) {

	linha := strings.Repeat("-", len(titulo)+6)
	fmt.Println(Cyan + linha + Reset)
	fmt.Printf(Cyan+"| %s |\n"+Reset, titulo)
	fmt.Println(Cyan + linha + Reset)

}

func digitarTexto(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

func CabecalhoAnimado(titulo string) {
	linha := strings.Repeat("-", len(titulo)+6)
	fmt.Println(Cyan + linha + Reset)
	digitarTexto(fmt.Sprintf("\033[36m| %s |\033[0m", titulo), 50*time.Millisecond)
	fmt.Println(Cyan + linha + Reset)
}

func CabecalhoASCII(titulo string) {
	fig := figure.NewColorFigure(titulo, "slant", "green", true)
	fig.Print()
}
