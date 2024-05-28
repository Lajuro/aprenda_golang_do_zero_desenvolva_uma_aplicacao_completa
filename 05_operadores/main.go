package main

func main() {
	// Operadores
	soma := 10 + 5          // Adição
	subtracao := 10 - 5     // Subtração
	multiplicacao := 10 * 5 // Multiplicação
	divisao := 10 / 5       // Divisão
	resto := 10 % 3         // Módulo (resto da divisão)

	println(soma, subtracao, multiplicacao, divisao, resto)

	// Conversão de tipos
	var num1 int16 = 10
	var num2 int32 = 20
	// soma := num1 + num2 // Isso dará erro

	soma2 := num1 + int16(num2)
	println(soma2)

	// Operadores aritméticos
	var a int
	a = 10  // Atribuição simples
	b := 20 // Atribuição com inferência de tipo
	a += 5  // Atribuição composta (a = a + 5)
	a -= 3  // Atribuição composta (a = a - 3)
	a *= 2  // Atribuição composta (a = a * 2)
	a /= 2  // Atribuição composta (a = a / 2)
	a %= 3  // Atribuição composta (a = a % 3)

	println(a, b)

	// Operadores relacionais
	igual := (10 == 10)     // Igualdade
	diferente := (10 != 5)  // Diferença
	maior := (10 > 5)       // Maior que
	menor := (10 < 5)       // Menor que
	maiorIgual := (10 >= 5) // Maior ou igual a
	menorIgual := (10 <= 5) // Menor ou igual a

	println(igual, diferente, maior, menor, maiorIgual, menorIgual)

	// Operadores lógicos
	and := (true && false) // E lógico
	or := (true || false)  // Ou lógico
	not := !true           // Negação

	println(and, or, not)

	// Incremento e decremento
	var num int = 10
	num++    // Incrementa 1 (equivale a num = num + 1)
	num--    // Decrementa 1 (equivale a num = num - 1)
	num += 5 // Incrementa 5 (equivale a num = num + 5)
	num -= 3 // Decrementa 3 (equivale a num = num - 3)

	println(num)

	// Condicional
	var texto string
	if num > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor ou igual a 5"
	}

	println(texto)
}
