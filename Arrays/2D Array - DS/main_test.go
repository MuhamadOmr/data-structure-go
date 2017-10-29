package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var matrix [][]int

func main() {
	// read the matrix
	lineScanner := bufio.NewScanner(os.Stdin)
	lineScanner.Split(bufio.ScanLines)
	matrix = [][]int{}


	for lineScanner.Scan() {
		line := lineScanner.Text()
		numberScanner := bufio.NewScanner(strings.NewReader(line))
		numberScanner.Split(bufio.ScanWords)
		row := []int{}
		for numberScanner.Scan() {
			n, err := strconv.Atoi(numberScanner.Text())
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		if len(row) != 6 {
			log.Fatalf("each row must be exact 6 columns")
		}
		matrix = append(matrix, row)
	}
	if len(matrix) != 6 {
		log.Fatalf("matrix must have exact 6 rows")
	}

	largest := -63
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := matrix[i][j] + matrix[i][j+1] + matrix[i][j+2] + matrix[i+1][j+1] + matrix[i+2][j] + matrix[i+2][j+1] + matrix[i+2][j+2]
			if sum > largest {
				largest = sum
			}
		}
	}
	fmt.Println(largest)
}