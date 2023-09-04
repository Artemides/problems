package leetcode

import (
	"fmt"
)

func FlippingMatrix(matrix [][]int32) int32 {
	matrix = flipRowCol(matrix, 0)

	fmt.Printf("%+v\n", matrix)
	return 0
}

func flipRowCol(matrix [][]int32, index int) [][]int32 {
	if index >= len(matrix[0]) {
		return matrix
	}

	fmt.Println("before reverse ", matrix)
	matrix = reverseIndex(matrix, index)
	fmt.Println("after reverse ", matrix)
	fmt.Printf("\n\n\n")
	return flipRowCol(matrix, index+1)
}
func reverseIndex(matrix [][]int32, index int) [][]int32 {
	if len(matrix[0]) <= 1 {
		return matrix
	}
	half := len(matrix[0]) / 2
	if index < half {
		if ableToReverseBasedOnSum(matrix, matrix[index]) {
			fmt.Println("REVERSE HALF ROW ", matrix)
			matrix[index] = reverseSlice(matrix[index])
		}
		if ableToReverseBasedOnSum(matrix, transposeMatrix(matrix)[index]) {
			fmt.Println("REVERSE HALF COLUMNS ", matrix)
			matrix = transposeMatrix(matrix)
			matrix[index] = reverseSlice(matrix[index])
			matrix = transposeMatrix(matrix)
		}
		return matrix

	} else {
		if !ableToReverse(matrix[index], transposeMatrix(matrix)[index]) && !ableToReverse(transposeMatrix(matrix)[index], matrix[index]) {
			return matrix
		}

		if ableToReverse(matrix[index], transposeMatrix(matrix)[index]) {
			matrix[index] = reverseSlice(matrix[index])
			fmt.Println("REVERSE  ROW ", matrix)
		}

		if ableToReverse(transposeMatrix(matrix)[index], matrix[index]) {
			matrix = transposeMatrix(matrix)
			matrix[index] = reverseSlice(matrix[index])
			matrix = transposeMatrix(matrix)
			fmt.Println("REVERSE COLUMN", matrix)

		}

		return reverseIndex(matrix, index)
	}

}

func ableToReverse(row []int32, col []int32) bool {
	length := len(row) / 2

	return sumSlice(row[:length]) < sumSlice(col[length:]) && sumSlice(row[:length]) < sumSlice(row[length:])

}

func ableToReverseBasedOnSum(matrix [][]int32, row []int32) bool {
	sum := sumHalfMatrix(matrix)
	half := len(row) / 2
	left := sumSlice(row[:half])
	right := sumSlice(row[half:])

	fmt.Println("reverse? ", sum, left, right, sum-left+right, sum < sum-left+right)
	return sum < sum-left+right
}

func sumSlice(slice []int32) int32 {
	var total int32
	for _, v := range slice {
		total += v
	}
	return total
}

func reverseSlice(slice []int32) []int32 {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func transposeMatrix(matrix [][]int32) [][]int32 {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	numRows := len(matrix)
	numCols := len(matrix[0])

	transposed := make([][]int32, numCols)
	for i := range transposed {
		transposed[i] = make([]int32, numRows)
		for j := range transposed[i] {
			transposed[i][j] = matrix[j][i]
		}
	}

	return transposed
}

func sumHalfMatrix(matrix [][]int32) int32 {
	half := len(matrix[0]) / 2
	var sum int32
	for i := 0; i < half; i++ {
		for j := 0; j < half; j++ {
			sum += matrix[i][j]
		}
	}
	return sum
}
