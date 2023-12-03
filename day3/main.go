package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// import any other required packages
)

func main() {
	// Read input file
	data, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Process the data
	result := solvePart1(data)

	// Print the result
	fmt.Println("Answer:", result)
}

type Matrix struct {
	Height int
	Width  int
}

func (m *Matrix) createIsValidSymbolMatrix(matrix [][]byte) [][]bool {
	// Matrix telling if the character is a valid symbol
	isValidSymbolMatrix := make([][]bool, m.Height)
	for i, line := range matrix {
		isValidSymbolMatrix[i] = make([]bool, m.Width)
		for j, char := range line {
			// Check if . (byte 46)
			if char == 46 || isDigit(char) {
				isValidSymbolMatrix[i][j] = false
			} else {
				isValidSymbolMatrix[i][j] = true
			}
		}
	}
	return isValidSymbolMatrix
}

func (m *Matrix) createIsDigitMatrix(matrix [][]byte) [][]bool {
	// Matrix telling if the character is a digit
	isDigitMatrix := make([][]bool, m.Height)
	for i, line := range matrix {
		isDigitMatrix[i] = make([]bool, m.Width)
		for j, char := range line {
			if isDigit(char) {
				isDigitMatrix[i][j] = true
			} else {
				isDigitMatrix[i][j] = false
			}
		}
	}
	return isDigitMatrix
}

func (m *Matrix) createNumCloseToSymbol(isDigitMatrix [][]bool, isSymMatrix [][]bool) [][]bool {
	// Create matrix that shows if digit in position is close to symbol
	numCloseToSymbolMatrix := make([][]bool, m.Height)
	for i := 1; i < m.Height-1; i++ {
		numCloseToSymbolMatrix[i] = make([]bool, m.Width)
		for j := 1; j < m.Width-1; j++ {
			if isDigitMatrix[i][j] {
				if isSymMatrix[i][j] || // same place
					isSymMatrix[i][j+1] || // right
					isSymMatrix[i][j-1] || // left
					isSymMatrix[i+1][j] || // above
					isSymMatrix[i-1][j] || // below
					isSymMatrix[i-1][j+1] || // diagonal lower right
					isSymMatrix[i+1][j+1] || // diagonal upper right
					isSymMatrix[i+1][j-1] || // diagonal upper left
					isSymMatrix[i-1][j-1] { // diagonal lower left

					numCloseToSymbolMatrix[i][j] = true
				}
			} else {
				numCloseToSymbolMatrix[i][j] = false
			}
		}
	}
	return numCloseToSymbolMatrix
}

func solvePart1(input string) int {

	//f, _ := readFile("./day3/input.txt")
	data, err := readFile("./input.txt")
	if err != nil {
		log.Println("Error reading file:", err)
	}

	lines := strings.Split(data, "\n")
	width, height := getMatrixSize(lines)
	fmt.Println("Matrix width ", width, " - matrix height ", height)

	m := Matrix{height, width}

	// Create matrix of characters (bytes)
	matrix := make([][]byte, height)
	for i, line := range lines {
		matrix[i] = make([]byte, width)

		for j := 0; j < width; j++ {
			matrix[i][j] = line[j]

		}
	}

	isSymMatrix := m.createIsValidSymbolMatrix(matrix)
	isDigitMatrix := m.createIsDigitMatrix(matrix)
	numCloseToSymbol := m.createNumCloseToSymbol(isDigitMatrix, isSymMatrix)

	numberString := ""
	sum := 0
	// Loop over matrix
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			// Detect number
			if numCloseToSymbol[i][j] {
				numberString = string(matrix[i][j])
				found_first_digit := false
				jj := j
				// go left and add numbers to string
				for !found_first_digit {
					jj--
					if jj == 0 {
						break
					}
					if !isDigitMatrix[i][jj] {
						found_first_digit = true
						break
					}
					numberString = string(matrix[i][jj]) + numberString
				}
				jj = j
				// go right and add numbers to string
				found_last_digit := false
				for !found_last_digit {
					jj++
					if jj == width {
						break
					}
					if !isDigitMatrix[i][jj] {
						found_last_digit = true
						break
					}
					numberString = numberString + string(matrix[i][jj])

				}
				j = jj
				num, _ := strconv.Atoi(numberString)
				sum += num

			} else {
				numberString = ""
			}
		}
	}

	return sum
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func getMatrixSize(lines []string) (int, int) {
	height := len(lines)
	width := len(lines[0])
	return width, height
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
