package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// testing the randSlice function in main.go
func TestRandSlice(t *testing.T) {

	rand.Seed(1)

	in1, in2 := 5, 5
	expected := []float64{1, 2, 2, 4, 1}
	out := randSlice(in1, in2)

	if reflect.DeepEqual(expected, out) != true {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// testing the sampledSlice function in main.go
func TestSampledSlice(t *testing.T) {

	rand.Seed(1)

	in := []float64{5, 4, 3, 2, 1}
	expected := []float64{4, 3, 3, 1, 4}
	out := sampledSlice(in)

	if reflect.DeepEqual(expected, out) != true {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// testing the median function in main.go
func TestMedian(t *testing.T) {

	in := []float64{3, 2, 1}
	expected := float64(2)
	out := median(in)

	if expected != out {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// testing the medianSE function in main.go
func TestMedianSE(t *testing.T) {

	rand.Seed(1)

	in := []float64{3, 2, 1}
	expected := 0.5773502691896258
	out := medianSE(in)

	if expected != out {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// benchmarking the workflow without all of the writing conditions
func BenchmarkWithout(b *testing.B) {

	sample := 100
	nums := 100

	dataChannel := make(chan []float64)

	testSlice := randSlice(sample, 10)
	go boots(dataChannel, testSlice, nums)
	testMedians := bootsMedian(dataChannel)
	testSE := medianSE(testMedians)

	fmt.Println("Median Standard Error:", testSE)

}
