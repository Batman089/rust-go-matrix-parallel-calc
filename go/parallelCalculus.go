package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	// File paths for the source matrices
	sourceMatrixA := "resources/bigMatrix1.txt"
	sourceMatrixB := "resources/bigMatrix2.txt"

	// Generate two 1000x1000 matrices and save them to files
	generateMatrixToFile(sourceMatrixA, 1000, 1000)
	generateMatrixToFile(sourceMatrixB, 1000, 1000)

	// Read matrices from files
	matrixA := readMatrixFromFile(sourceMatrixA)
	matrixB := readMatrixFromFile(sourceMatrixB)

	// Check if matrix multiplication is possible
	if len(matrixA[0]) != len(matrixB) {
		fmt.Println("Matrix multiplication is not possible due to dimension mismatch")
		return
	}

	// Start time for matrix multiplication
	calcTimeStart := time.Now()

	// Create log file for calculation time
	calcTimeLog, errCreate := os.Create("go/log/calcTimeLog")
	if errCreate != nil {
		fmt.Println("Error creating calcTimeLog:", errCreate)
		return
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
}

// generateMatrixToFile generates a matrix with random values and saves it to a file
func generateMatrixToFile(filename string, rows, cols int) {
	// Start time for matrix generation
	generateTimeStart := time.Now()

	// Create log file for matrix generation
	generateMatrixFilesLog, err := os.Create("go/log/generateMatrixFilesLog")
	matrixFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating matrixFile:", err)
		return
	}
	defer generateMatrixFilesLog.Close()
	defer matrixFile.Close()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	writer := bufio.NewWriter(matrixFile)
	for i := 0; i < rows; i++ {
		var row []string
		for j := 0; j < cols; j++ {
			row = append(row, strconv.Itoa(rand.Intn(100)))
		}
		writer.WriteString(strings.Join(row, " ") + "\n")
	}

	// End time for matrix generation
	generateTimeEnde := time.Now()
	generateTimeTotal := generateTimeEnde.Sub(generateTimeStart)
	fmt.Println("Matrix generation time:", generateTimeTotal)

	// Log the generation time
	logWriter := bufio.NewWriter(generateMatrixFilesLog)
	logWriter.WriteString("Matrix generation Start time: " + generateTimeStart.String() + "\n")
	logWriter.WriteString("Matrix generation End time: " + generateTimeEnde.String() + "\n")
	logWriter.WriteString("Matrix generation time: " + generateTimeTotal.String() + "\n")

	writer.Flush()
	logWriter.Flush()
}

// readMatrixFromFile reads a matrix from a file
func readMatrixFromFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		row := make([]int, len(values))
		for i, value := range values {
			row[i], _ = strconv.Atoi(value)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return matrix
}
