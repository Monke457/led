package main

import "fmt"

func main() {
	height := 32
	width := 32
	matrix := make([][]int, height)
	for i := 0; i < height; i++ {
		matrix[i] = make([]int, width)
		for j := 0; j < width; j++ {
			matrix[i][j] = 0
		}
	}

	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
