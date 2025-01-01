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
	if err != nil {
		fmt.Println("Error creating generateMatrixFilesLog:", err)
		return
	}
	defer generateMatrixFilesLog.Close()

	matrixFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating matrixFile:", err)
		return
	}
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

	fileTimeLogger(generateMatrixFilesLog, generateTimeStart, generateTimeEnde, generateTimeTotal)
}

func ReadMatrixFromFile(filename string) [][]int {
	handleError := func(err error, message string) [][]int {
		if err != nil {
			fmt.Println(message, err)
			return nil
		}
		return nil
	}

	file, err := os.Open(filename)
	if handleError(err, "Error opening file:") != nil {
		return nil
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

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

	if handleError(scanner.Err(), "Error reading file:") != nil {
		return nil
	}

	return matrix
}

func fileTimeLogger(generateMatrixFilesLog *os.File, generateTimeStart time.Time, generateTimeEnde time.Time, generateTimeTotal time.Duration) {
	logWriter := bufio.NewWriter(generateMatrixFilesLog)
	_, err := logWriter.WriteString("Matrix generation Start time: " + generateTimeStart.String() + "\n")
	if err != nil {
		return
	}
	_, err = logWriter.WriteString("Matrix generation End time: " + generateTimeEnde.String() + "\n")
	if err != nil {
		return
	}
	_, err = logWriter.WriteString("Matrix generation time: " + generateTimeTotal.String() + "\n")
	if err != nil {
		return
	}

	err = logWriter.Flush()
	if err != nil {
		return
	}
}
