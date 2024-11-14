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
	sourceMatrixA := "resources/bigMatrix1.txt"
	sourceMatrixB := "resources/bigMatrix2.txt"
	// Generate two 500x500 matrices and save them to files
	generateMatrixToFile(sourceMatrixA, 1000, 1000)
	generateMatrixToFile(sourceMatrixB, 1000, 1000)

	matrixA := readMatrixFromFile(sourceMatrixA)
	matrixB := readMatrixFromFile(sourceMatrixB)

	if len(matrixA[0]) != len(matrixB) {
		fmt.Println("Matrix multiplication is not possible due to dimension mismatch")
		return
	}

	// Multiply the two matrices
	calcTimeStart := time.Now()
	calcTimeLog, errCreate := os.Create("go/log/calcTimeLog")
	if errCreate != nil {
		fmt.Println("Error creating calcTimeLog:", errCreate)
		return
	}
	writerLog := bufio.NewWriter(calcTimeLog)

	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixB[0]))
	}

	var wg sync.WaitGroup
	numWorkers := 10
	chunkSize := len(matrixA) / numWorkers

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

	wg.Wait()
	calcTimeEnde := time.Now()
	calcTimeTotal := calcTimeEnde.Sub(calcTimeStart)

	writerLog.WriteString("Matrix multiplication Start time: " + calcTimeStart.String() + "\n")
	writerLog.WriteString("Matrix multiplication End time: " + calcTimeEnde.String() + "\n")
	writerLog.WriteString("Matrix multiplication time: " + calcTimeTotal.String() + "\n")
	writerLog.Flush()

	fmt.Println("Matrix multiplication completed. \n")
	fmt.Println("Matrix multiplication time:", calcTimeTotal)
}

func generateMatrixToFile(filename string, rows, cols int) {
	generateTimeStart := time.Now()
	generateMatrixFilesLog, err := os.Create("go/log/generateMatrixFilesLog")
	matrixFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating matrixFile:", err)
		return
	}
	defer generateMatrixFilesLog.Close()
	defer matrixFile.Close()

	rand.Seed(time.Now().UnixNano())
	writer := bufio.NewWriter(matrixFile)
	for i := 0; i < rows; i++ {
		var row []string
		for j := 0; j < cols; j++ {
			row = append(row, strconv.Itoa(rand.Intn(100)))
		}
		writer.WriteString(strings.Join(row, " ") + "\n")
	}
	generateTimeEnde := time.Now()
	generateTimeTotal := generateTimeEnde.Sub(generateTimeStart)
	fmt.Println("Matrix generation time:", generateTimeTotal)

	logWriter := bufio.NewWriter(generateMatrixFilesLog)

	logWriter.WriteString("Matrix generation Start time: " + generateTimeStart.String() + "\n")
	logWriter.WriteString("Matrix generation End time: " + generateTimeEnde.String() + "\n")
	logWriter.WriteString("Matrix generation time: " + generateTimeTotal.String() + "\n")

	writer.Flush()
	logWriter.Flush()
}

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
