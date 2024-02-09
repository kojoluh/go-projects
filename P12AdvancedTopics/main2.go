package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1.5, 2.5, 3.5}
	fmt.Println(sumSlice[float32](float32Slice))

	var stringSlice = []string {"new", "false"}
	fmt.Println(isEmpty[string](stringSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T 
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}