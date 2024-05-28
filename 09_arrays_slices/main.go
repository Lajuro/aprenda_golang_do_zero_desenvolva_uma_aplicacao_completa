package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr[0] = 10
	fmt.Println(arr[0])

	s := []int{1, 2, 3, 4, 5}
	s = append(s, 6)
	fmt.Println(s)

	arr2 := [5]int{1, 2, 3, 4, 5}
	s2 := arr2[1:4] // s = [2, 3, 4]
	fmt.Println(s2)

	fmt.Println(reflect.TypeOf(s2))
}
