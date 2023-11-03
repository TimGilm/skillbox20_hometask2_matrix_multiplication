// Задание 2. Умножение матриц
// Что нужно сделать - Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.
package main

import "fmt"

const rowsA = 3
const colsA = 5
const rowsB = 5
const colsB = 4

func prodMatrix(A [rowsA][colsA]int, B [rowsB][colsB]int) (C [rowsA][colsB]int) {
	for i := 0; i < rowsA; i++ {
		l := 0
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				C[i][j] += A[i][k] * B[k][l]
			}
			l++
		}
	}
	return C
}

func main() {
	A := [rowsA][colsA]int{
		{1, 2, 3, 2, 1},
		{1, 2, 3, 2, 1},
		{1, 2, 3, 2, 1},
	}
	B := [rowsB][colsB]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
		{2, 2, 2, 2},
		{1, 1, 1, 1},
	}
	fmt.Println(prodMatrix(A, B))
}
