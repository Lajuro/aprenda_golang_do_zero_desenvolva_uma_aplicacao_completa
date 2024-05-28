package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Altura    float64
}

type Estudante struct {
	Pessoa
	Curso     string
	Faculdade string
}

func main() {
	pessoa := Pessoa{
		Nome:      "Roberto",
		Sobrenome: "Silva",
		Idade:     22,
		Altura:    1.75,
	}
	estudante := Estudante{
		Pessoa:    pessoa,
		Curso:     "Engenharia de Computação",
		Faculdade: "USP",
	}

	fmt.Println(fmt.Sprintf("%v", estudante))

}
