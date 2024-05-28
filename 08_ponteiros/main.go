package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int = 100
	var variavel4 *int = &variavel3
	fmt.Println(variavel3, *variavel4)
	variavel3++
	fmt.Println(variavel3, *variavel4)
	*variavel4++
	fmt.Println(variavel3, *variavel4)
}
