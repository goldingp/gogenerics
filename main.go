package main

import "fmt"

// Represents an integer or a rational number.
type number interface {
	int64 | float64
}

// The entry point of the module.
func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", sumInts(ints), sumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		sumIntsOrFloats[string, int64](ints),
		sumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with interface constraint: %v and %v\n",
		sumNumbers(ints),
		sumNumbers(floats))
}

// Returns the sum of the rational number values in map m.
func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Returns the sum of the integer number values in map m.
func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// Returns the sum of the values in map m with integer or rational numbers
// supported as the value type.
func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Returns the sum of the values in map m with a supported value type
// of the number interface.
func sumNumbers[K comparable, V number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
