package main

import "math"

func main() {

}

func parseArray(str string) [][]int {
	var temp string = str
	for len(temp) > 0 {

	}
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
		if ascending && arr[i] < arr[i+1] {
			return false
		}
		if descending && arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
