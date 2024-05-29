package main

import (
	"fmt"
)

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaSemana string

	switch {
	case numero == 1:
		diaSemana = "Domingo"
	case numero == 2:
		diaSemana = "Segunda-feira"
	case numero == 3:
		diaSemana = "Terça-feira"
	case numero == 4:
		diaSemana = "Quarta-feira"
	case numero == 5:
		diaSemana = "Quinta-feira"
	case numero == 6:
		diaSemana = "Sexta-feira"
	case numero == 7:
		diaSemana = "Sábado"
	default:
		diaSemana = "Número inválido"
	}

	return diaSemana
}

func main() {
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(10))

	fmt.Println(diaDaSemana2(3))
	fmt.Println(diaDaSemana2(5))
	fmt.Println(diaDaSemana2(10))

	// Exemplo de fallthrough
	numero := 1
	switch numero {
	case 1:
		fmt.Println("Domingo")
		fallthrough
	case 2:
		fmt.Println("Segunda-feira")
	default:
		fmt.Println("Número inválido")
	}
}
