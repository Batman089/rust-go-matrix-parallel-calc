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

	generateMatrixFilesLog, matrixFile, done := createFiles(filename)
	if done {
		return
	}

	defer generateMatrixFilesLog.Close()
	defer matrixFile.Close()

	writeInFile(matrixFile, size)

	generateTimeEnd := time.Now()
	fileTimeLogger(generateMatrixFilesLog, generateTimeStart, generateTimeEnd, generateTimeEnd.Sub(generateTimeStart))
}

func writeInFile(matrixFile *os.File, size int) {
	rand.Seed(time.Now().UnixNano())
	writer := bufio.NewWriter(matrixFile)
	for i := 0; i < size; i++ {
		var row []string
		for j := 0; j < size; j++ {
			row = append(row, strconv.Itoa(rand.Intn(100)))
		}
		writer.WriteString(strings.Join(row, " ") + "\n")
	}

	writer.Flush()
}

func createFiles(filename string) (*os.File, *os.File, bool) {
	err := os.MkdirAll("../go/generated/resources", os.ModePerm)
	if err != nil {
		return nil, nil, true
	}

	// Create log file for matrix generation
	generateMatrixFilesLog, err := os.Create("./go/generated/log/generateMatrixFilesLog.txt")
	matrixFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating matrixFile:", err)
		return nil, nil, true
	}
	return generateMatrixFilesLog, matrixFile, false
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

func fileTimeLogger(generateMatrixFilesLog *os.File, generateTimeStart time.Time, generateTimeEnd time.Time, generateTimeTotal time.Duration) {
	logWriter := bufio.NewWriter(generateMatrixFilesLog)
	fmt.Println("Matrix generation time:", generateTimeTotal)
	_, err := logWriter.WriteString("Matrix generation Start time: " + generateTimeStart.String() + "\n")
	if err != nil {
		return
	}
	_, err = logWriter.WriteString("Matrix generation End time: " + generateTimeEnd.String() + "\n")
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
