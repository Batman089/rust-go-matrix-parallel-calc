package testing

import (
	"../matrixutils"
	"os"
	"testing"
)

func TestSmallMatrixMultiplyBigMatrix(t *testing.T) {
	matrixSizeA := matrixutils.Small
	matrixSizeB := matrixutils.Big
	numWorkers := 4

	sourceMatrixA := "./go/generated/resources/matrixA.txt"
	sourceMatrixB := "./go/generated/resources/matrixB.txt"

	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixSizeA))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixSizeB))

	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if result == nil {
		t.Error("Expected non-nil result for valid matrix multiplication")
	}
}

func TestValidMatrixInvalidWorkerNumbers(t *testing.T) {
	matrixSizeA := matrixutils.Small
	matrixSizeB := matrixutils.Small
	numWorkers := 0

	sourceMatrixA := "./go/generated/resources/matrixA.txt"
	sourceMatrixB := "./go/generated/resources/matrixB.txt"

	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixSizeA))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixSizeB))

	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if result != nil {
		t.Error("Expected nil result for invalid worker number")
	}
}

func TestUnavailableMatrixNames(t *testing.T) {
	sourceMatrixA := "./go/generated/resources/nonexistentA.txt"
	sourceMatrixB := "./go/generated/resources/nonexistentB.txt"

	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	if matrixA != nil || matrixB != nil {
		t.Error("Expected nil matrices for unavailable matrix files")
	}
}

func TestDimensionMismatch(t *testing.T) {
	matrixSizeA := matrixutils.Small
	matrixSizeB := matrixutils.Middle
	numWorkers := 4

	sourceMatrixA := "./go/generated/resources/matrixA.txt"
	sourceMatrixB := "./go/generated/resources/matrixB.txt"

	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixSizeA))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixSizeB))

	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if result != nil {
		t.Error("Expected nil result for dimension mismatch")
	}
}

func TestValidMatrixValidWorkerNumbers(t *testing.T) {
	matrixSizeA := matrixutils.Middle
	matrixSizeB := matrixutils.Middle
	numWorkers := 4

	sourceMatrixA := "./go/generated/resources/matrixA.txt"
	sourceMatrixB := "./go/generated/resources/matrixB.txt"

	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixSizeA))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixSizeB))

	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if result == nil {
		t.Error("Expected non-nil result for valid matrix multiplication")
	}
}

func TestEmptyMatrices(t *testing.T) {
	matrixA := [][]int{}
	matrixB := [][]int{}
	numWorkers := 4

	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if result != nil {
		t.Error("Expected nil result for empty matrices")
	}
}

func TestSingleElementMatrices(t *testing.T) {
	matrixA := [][]int{{1}}
	matrixB := [][]int{{2}}
	numWorkers := 4

	expected := [][]int{{2}}
	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if !matrixutils.CompareMatrices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestLargeMatrices(t *testing.T) {
	matrixSizeA := matrixutils.Big
	matrixSizeB := matrixutils.Big
	numWorkers := 4

	sourceMatrixA := "./go/generated/resources/largeMatrixA.txt"
	sourceMatrixB := "./go/generated/resources/largeMatrixB.txt"

	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixSizeA))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixSizeB))

	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if result == nil {
		t.Error("Expected non-nil result for large matrix multiplication")
	}
}

func TestNegativeWorkerNumbers(t *testing.T) {
	matrixSizeA := matrixutils.Small
	matrixSizeB := matrixutils.Small
	numWorkers := -1

	sourceMatrixA := "./go/generated/resources/matrixA.txt"
	sourceMatrixB := "./go/generated/resources/matrixB.txt"

	matrixutils.GenerateMatrixToFile(sourceMatrixA, int(matrixSizeA))
	matrixutils.GenerateMatrixToFile(sourceMatrixB, int(matrixSizeB))

	matrixA := matrixutils.ReadMatrixFromFile(sourceMatrixA)
	matrixB := matrixutils.ReadMatrixFromFile(sourceMatrixB)

	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if result != nil {
		t.Error("Expected nil result for negative worker number")
	}
}

func TestNonSquareMatrices(t *testing.T) {
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrixB := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}
	numWorkers := 4

	expected := [][]int{
		{58, 64},
		{139, 154},
	}
	result := matrixutils.CalculateMatrix(matrixA, matrixB, numWorkers)
	if !matrixutils.CompareMatrices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMain(m *testing.M) {
	os.MkdirAll("./go/generated/log", os.ModePerm)
	os.MkdirAll("./go/generated/resources", os.ModePerm)
	code := m.Run()
	os.RemoveAll("./go/generated")
	os.Exit(code)
}
