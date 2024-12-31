package matrixutils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateMatrixToFile(filename string, size int) {
	// Start time for matrix generation
	generateTimeStart := time.Now()

	// Create log file for matrix generation
	generateMatrixFilesLog, err := os.Create("./go/generated/log/generateMatrixFilesLog.txt")
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
	for i := 0; i < size; i++ {
		var row []string
		for j := 0; j < size; j++ {
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

func ReadMatrixFromFile(filename string) [][]int {
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
