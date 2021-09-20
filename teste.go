package main

import (
	"fmt"
	arvore "hello_world/tree"
)

func Add(value1 int, value2 int) int {

	return value1 + value2
}
func Sub(value1 int, value2 int) int {
	return value1 * value2
}

func main() {
	matrix := [6][6]int{
		{5, 5, 7, 6, 9, 6},
		{4, 5, 3, 0, 5, 9},
		{43, 52, 36, 10, 51, 96},
		{45, 33, 34, 21, 35, 87},
		{5, 2, 15, 55, 61, 9},
		{1, 7, 6, 2, 1, 9},
	}
	for i := 0; i < len(matrix[i]); i++ {
		for j := 0; j < len(matrix[j]); j++ {
			fmt.Printf("Matriz[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}
	// var resultado int = adicionar.Soma(5, 9)

	// fmt.Printf("Resultado da soma: %d", resultado)

	var result = arvore.Insert(arvore.MyTree())
	fmt.Println("Resultado da chave:", result)
}
