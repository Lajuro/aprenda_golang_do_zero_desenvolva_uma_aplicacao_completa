package main

import (
	"fmt"             // Importa o pacote fmt para formatação
	"modulo/auxiliar" // Importa o pacote auxiliar do módulo

	"github.com/badoux/checkmail" // Importa o pacote checkmail para validação de email
)

func main() {
	fmt.Println("Executando do arquivo main") // Imprime uma mensagem no console
	auxiliar.Escrever()                       // Chama a função Escrever do pacote auxiliar

	// Validar um email
	email := "example@example.com"
	err := checkmail.ValidateFormat(email)
	if err != nil {
		fmt.Println("Email inválido:", err) // Imprime uma mensagem se o email for inválido
	} else {
		fmt.Println("Email válido:", email) // Imprime uma mensagem se o email for válido
	}
}
