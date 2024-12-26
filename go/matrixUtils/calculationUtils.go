package matrixutils

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type MatrixSize int

const (
	Small  MatrixSize = 1000
	Middle MatrixSize = 5000
	Big    MatrixSize = 10000
)

func CalculateMatrix(matrixA, matrixB [][]int) [][]int {
	// Check if matrix multiplication is possible
	if len(matrixA[0]) != len(matrixB) {
		fmt.Println("Matrix multiplication is not possible due to dimension mismatch")
		return nil
	}

	// Start time for matrix multiplication
	calcTimeStart := time.Now()

	// Create log file for calculation time
	calcTimeLog, errCreate := os.Create("go/log/calcTimeLog")
	if errCreate != nil {
		fmt.Println("Error creating calcTimeLog:", errCreate)
		return nil
	}
	writerLog := bufio.NewWriter(calcTimeLog)

	// Initialize result matrix
	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixB[0]))
	}

	// Use WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	numWorkers := 10
	chunkSize := len(matrixA) / numWorkers

	// Launch goroutines to perform matrix multiplication in parallel
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(startRow int) {
			defer wg.Done()
			endRow := startRow + chunkSize
			if endRow > len(matrixA) {
				endRow = len(matrixA)
			}
			for row := startRow; row < endRow; row++ {
				for col := 0; col < len(matrixB[0]); col++ {
					for k := 0; k < len(matrixB); k++ {
						result[row][col] += matrixA[row][k] * matrixB[k][col]
					}
				}
			}
		}(i * chunkSize)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// End time for matrix multiplication
	calcTimeEnde := time.Now()
	calcTimeTotal := calcTimeEnde.Sub(calcTimeStart)

	// Log the calculation time
	writerLog.WriteString("Matrix multiplication Start time: " + calcTimeStart.String() + "\n")
	writerLog.WriteString("Matrix multiplication End time: " + calcTimeEnde.String() + "\n")
	writerLog.WriteString("Matrix multiplication duration time: " + calcTimeTotal.String() + "\n")
	writerLog.Flush()

	fmt.Println("Matrix multiplication completed.")
	fmt.Println("Matrix multiplication time:", calcTimeTotal)

	return result
}
