package main

import (
	"./matrixutils"
)

func main() {
	// File paths for the source matrices
	sourceMatrixA := "resources/middleMatrix1.txt"
	sourceMatrixB := "resources/middleMatrix2.txt"

	// extract the variables as enum from the matrixutils package
	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixutils.Middle))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixutils.Middle))

	// Read matrices from files
	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	// Calculate the calculateMatrix matrix
	calculateMatrix := matrixutils.CalculateMatrix(matrixA, matrixB)

	// You can now use the calculateMatrix matrix as needed
	_ = calculateMatrix
}
