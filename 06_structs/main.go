package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
}

type Usuario struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func main() {
	endereco := Endereco{"Rua dos Bobos", 0}
	usuario := Usuario{"Davi", 21, endereco}

	fmt.Println(usuario)
}
