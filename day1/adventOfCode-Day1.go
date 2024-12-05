package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(checkFile("input.txt"))
}

func checkFile(filename string) int {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count++
	}
	left := make([]int, count)
	right := make([]int, count)

	file, err = os.Open(filename) //reopen file
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	count = 0
	for scanner.Scan() {
		// split each line into number array
		row := strings.Split(scanner.Text(), "   ")
		intRow := make([]int, len(row))
		for i := 0; i < len(row); i++ {
			intRow[i], err = strconv.Atoi(row[i])
		}
		left[count] = intRow[0]
		right[count] = intRow[1]
		count++
	}

	slices.Sort(right)
	slices.Sort(left)
	sum := 0
	for i := 0; i < len(right); i++ {
		sum += findDifference(right[i], left[i])
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return sum
}

func splitLines(array []int) ([]int, []int) {
	var left []int = make([]int, len(array)/2)
	var right []int = make([]int, len(array)/2)
	for i := 0; i < len(array); i++ {
		if i%2 == 0 {
			left[i/2] = array[i]
		} else {
			right[i/2] = array[i]
		}

	}
	return left, right
}

func findDifference(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}
