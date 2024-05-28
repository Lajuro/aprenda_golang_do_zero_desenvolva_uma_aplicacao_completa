package main

import (
	"errors"
	"fmt"
)

func main() {
	// Números inteiros
	var numInt8 int8 = 127
	var numInt16 int16 = 32767
	var numInt32 int32 = 2147483647
	var numInt64 int64 = 9223372036854775807

	fmt.Println("int8:", numInt8)
	fmt.Println("int16:", numInt16)
	fmt.Println("int32:", numInt32)
	fmt.Println("int64:", numInt64)

	// Números sem sinal
	var numUint8 uint8 = 255
	var numUint16 uint16 = 65535
	var numUint32 uint32 = 4294967295
	var numUint64 uint64 = 18446744073709551615

	fmt.Println("uint8:", numUint8)
	fmt.Println("uint16:", numUint16)
	fmt.Println("uint32:", numUint32)
	fmt.Println("uint64:", numUint64)

	// Números reais
	var numFloat32 float32 = 123.45
	var numFloat64 float64 = 123456789.123456789

	fmt.Println("float32:", numFloat32)
	fmt.Println("float64:", numFloat64)

	// Strings
	var texto string = "Olá, Go!"
	fmt.Println("string:", texto)

	// Booleanos
	var verdadeiro bool = true
	var falso bool = false
	fmt.Println("boolean verdadeiro:", verdadeiro)
	fmt.Println("boolean falso:", falso)

	// Erros
	var erro error = errors.New("Este é um erro")
	fmt.Println("error:", erro)
}
