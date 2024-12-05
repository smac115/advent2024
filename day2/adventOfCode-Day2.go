package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
		// split each line into number array
		row := strings.Split(scanner.Text(), " ")
		intRow := make([]int, len(row))
		for i := 0; i < len(row); i++ {
			intRow[i], err = strconv.Atoi(row[i])
		}
		if isProperNumberDistance(intRow) && isProperOrder(intRow) {
			count++
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return count
}

func isProperNumberDistance(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if math.Abs(float64(arr[i]-arr[i+1])) > 3 || math.Abs(float64(arr[i]-arr[i+1])) < 1 {
			return false
		}
	}
	return true
}
func isProperOrder(arr []int) bool {
	var ascending bool = false
	var descending bool = false

	if arr[0] > arr[1] {
		descending = true
	}
	if arr[0] < arr[1] {
		ascending = true
	}
	if arr[0] == arr[1] {
		return false
	}

	for i := 1; i < len(arr)-1; i++ {
		if ascending && arr[i] > arr[i+1] {
			return false
		}
		if descending && arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}
