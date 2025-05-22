package main

import (
	"fmt"
	"math/rand"

	"github.com/seehuhn/mt19937"
	"gonum.org/v1/gonum/stat"
)

// create a slice of random floats that is lngth long
func randSlice(lngth int, max int) []float64 {

	var result []float64

	for i := 0; i < lngth; i++ {
		result = append(result, float64(rand.Intn(max))) // Generates random integers between 0 and max
	}

	return result
}

// given a slice of ints, create a new slice that is a sample with replacement
// the same size as the initial slice
func sampledSlice(inital []float64) (final []float64) {

	sampleRange := len(inital)

	for i := 0; i < sampleRange; i++ {
		final = append(final, inital[rand.Intn(sampleRange)])
	}

	return final
}

// Resample the data with replacement n times
func boots(floatSlice []float64, n int) (bootsResult [][]float64) {

	for i := 0; i < n; i++ {

		// create a new slice with replacement
		sample := sampledSlice(floatSlice)

		// add it to the result
		bootsResult = append(bootsResult, sample)

	}

	return bootsResult

}

// Compute desired statistic n times to generate a distribution of estimated statistics
func bootsMean(bootsResult [][]float64) (meanSlice []float64) {

	for _, value := range bootsResult {
		meanSlice = append(meanSlice, stat.Mean(value, nil))
	}

	return meanSlice
}

//Determine standard error/confidence interval for the bootstrapped statistic from the bootstrapped distribution

func main() {

	// create a random generator to use in all the functions
	rand.New(mt19937.New())

	newSlice := randSlice(5, 10)
	fmt.Println(newSlice)

	newSample := boots(newSlice, 5)
	fmt.Println(newSample)

	newMean := bootsMean(newSample)
	fmt.Println(newMean)

}
