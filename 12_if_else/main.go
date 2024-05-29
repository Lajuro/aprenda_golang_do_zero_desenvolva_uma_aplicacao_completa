package main

import (
	"fmt"
)

func main() {
	numero := 10

	// Exemplo básico de if e else
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// If with short statement
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Outro número é maior que zero:", outroNumero)
	}

	// Outro exemplo de if with short statement com múltiplas condições
	novoNumero := -5

	if novoNumero < -10 {
		fmt.Println("Menor que -10")
	} else if novoNumero < 0 {
		fmt.Println("Entre 0 e -10")
	} else {
		fmt.Println("Maior ou igual a 0")
	}

	// Tentativa de acessar a variável fora do if (vai causar erro)
	// fmt.Println(outroNumero) // Isso causaria um erro, pois outroNumero não existe fora do if
}
