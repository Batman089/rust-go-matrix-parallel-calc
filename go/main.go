package main

import (
	"./matrixutils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getMatrixSizeFromUser(matrixName string) matrixutils.MatrixSize {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter size for %s matrix (small, middle, big): ", matrixName)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "small":
			return matrixutils.Small
		case "middle":
			return matrixutils.Middle
		case "big":
			return matrixutils.Big
		default:
			fmt.Println("Invalid input. Please enter 'small', 'middle', or 'big'.")
		}
	}
}

func main() {
	// Get matrix size from user for each matrix
	matrixSizeA := getMatrixSizeFromUser("Matrix A")
	matrixSizeB := getMatrixSizeFromUser("Matrix B")

	// File paths for the source matrices
	sourceMatrixA := "resources/matrixA.txt"
	sourceMatrixB := "resources/matrixB.txt"

	// Generate two matrices and save them to files using the user input
	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixSizeA))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixSizeB))

	// Read matrices from files
	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	// Calculate the result matrix
	result := matrixutils.CalculateMatrix(matrixA, matrixB)
	_ = result
}
