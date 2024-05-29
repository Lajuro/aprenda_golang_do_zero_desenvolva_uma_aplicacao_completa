package main

import "fmt"

func main() {
	s := make([]float32, 10, 15)

	fmt.Println(s)
	fmt.Println(len(s)) // Output: 10
	fmt.Println(cap(s)) // Output: 15

	s = append(s, 1)    // Excede a capacidade inicial
	fmt.Println(len(s)) // Output: 11
	fmt.Println(cap(s)) // Output: 30 (Capacidade dobrada)

	s2 := make([]int, 5)
	fmt.Println(s2) // Output: [0 0 0 0 0]

	fmt.Println(len(s2)) // Output: 5
	fmt.Println(cap(s2)) // Output: 5
}
