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

func CalculateMatrix(matrixA, matrixB [][]int, numWorkers int) [][]int {

	// Pre-checks
	i, done := preCheck(matrixA, matrixB, numWorkers)
	if done {
		return i
	}

	// Create log file for calculation time
	calcTimeLog, errCreate := os.Create("./go/generated/log/calcTimeLog.txt")
	if errCreate != nil {
		fmt.Println("Error creating calcTimeLog:", errCreate)
		return nil
	}

	calcTimeStart := time.Now()
	writerLog := bufio.NewWriter(calcTimeLog)

	// Initialize result matrix
	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixB[0]))
	}

	// Use WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	rowChannel := make(chan int, len(matrixA))

	// Launch goroutines to perform matrix multiplication in parallel
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rowChannel {
				for col := 0; col < len(matrixB[0]); col++ {
					for k := 0; k < len(matrixB); k++ {
						result[row][col] += matrixA[row][k] * matrixB[k][col]
					}
				}
			}
		}()
	}

	// Send row indices to the channel
	for row := 0; row < len(matrixA); row++ {
		rowChannel <- row
	}
	close(rowChannel)

	// Wait for all goroutines to finish
	wg.Wait()

	// End time for matrix multiplication
	calcTimeEnde := time.Now()
	calcTimeTotal := calcTimeEnde.Sub(calcTimeStart)

	timeLogger(writerLog, calcTimeStart, calcTimeEnde, calcTimeTotal)

	return result
}

func preCheck(matrixA [][]int, matrixB [][]int, numWorkers int) ([][]int, bool) {
	if numWorkers <= 0 {
		fmt.Println("Number of workers must be greater than zero")
		return nil, true
	}

	// Check if matrices are nil
	if matrixA == nil || matrixB == nil {
		fmt.Println("Matrix is nil")
		return nil, true
	}

	// Check if matrices are empty
	if len(matrixA) == 0 || len(matrixB) == 0 {
		fmt.Println("Matrix is empty")
		return nil, true
	}

	// Check if matrix multiplication is possible
	if len(matrixA[0]) != len(matrixB) {
		fmt.Println("Matrix multiplication is not possible due to dimension mismatch")
		return nil, true
	}

	return nil, false
}

func CompareMatrices(matrixA, matrixB [][]int) bool {
	if len(matrixA) != len(matrixB) {
		return false
	}
	for i := range matrixA {
		if len(matrixA[i]) != len(matrixB[i]) {
			return false
		}
		for j := range matrixA[i] {
			if matrixA[i][j] != matrixB[i][j] {
				return false
			}
		}
	}
	return true
}

func timeLogger(writerLog *bufio.Writer, calcTimeStart time.Time, calcTimeEnde time.Time, calcTimeTotal time.Duration) {
	// Log the calculation time
	writerLog.WriteString("Matrix multiplication Start time: " + calcTimeStart.String() + "\n")
	writerLog.WriteString("Matrix multiplication End time: " + calcTimeEnde.String() + "\n")
	writerLog.WriteString("Matrix multiplication duration time: " + calcTimeTotal.String() + "\n")
	writerLog.Flush()

	fmt.Println("Matrix multiplication completed.")
	fmt.Println("Matrix multiplication time:", calcTimeTotal)
}
