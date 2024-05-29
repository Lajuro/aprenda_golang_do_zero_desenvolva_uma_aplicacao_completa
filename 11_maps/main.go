package main

import "fmt"

func main() {
	// Criando um mapa simples
	usuario := make(map[string]string)
	usuario["nome"] = "Pedro"
	usuario["sobrenome"] = "Silva"
	fmt.Println("Mapa simples:", usuario)

	// Acessando valores do mapa
	nome := usuario["nome"]
	fmt.Println("Nome:", nome) // Output: Pedro

	// Modificando valores do mapa
	usuario["nome"] = "Carlos"
	fmt.Println("Nome modificado:", usuario["nome"]) // Output: Carlos

	// Criando um mapa aninhado
	usuario2 := make(map[string]map[string]string)
	usuario2["nome"] = map[string]string{"primeiro": "João", "último": "Pedro"}
	fmt.Println("Mapa aninhado:", usuario2)

	// Deletando uma chave do mapa
	delete(usuario, "sobrenome")
	fmt.Println("Após deleção:", usuario) // Output: map[nome:Carlos]

	// Adicionando novas chaves ao mapa
	usuario["idade"] = "30"
	fmt.Println("Após adição:", usuario) // Output: map[nome:Carlos idade:30]
}
