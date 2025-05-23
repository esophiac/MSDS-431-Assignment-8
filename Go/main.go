package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"

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

// Compute mean n times to generate a distribution of estimated statistics
func bootsMean(bootsResult [][]float64) (meanSlice []float64) {

	for _, value := range bootsResult {

		meanVal := stat.Mean(value, nil)
		meanSlice = append(meanSlice, meanVal)
	}

	return meanSlice
}

// Compute the median of a slice of floats
func median(sliceFloats []float64) (medianVal float64) {

	sort.Float64s(sliceFloats)
	medianVal = (stat.Quantile(0.5, stat.Empirical, sliceFloats, nil))

	return medianVal
}

// Compute median n times to generate a distribution of estimated statistics
func bootsMedian(bootsResult [][]float64) (medianSlice []float64) {

	for _, value := range bootsResult {

		medianVal := median(value)
		medianSlice = append(medianSlice, medianVal)
	}

	return medianSlice
}

// Compute the standard error of a slice
// Determine standard error for the bootstrapped statistic from the bootstrapped distribution
// Standard error is the standard deviation divided by the square root of the sample size
func medianSE(medianSlice []float64) (seMedian float64) {

	_, stdDev := stat.MeanStdDev(medianSlice, nil)

	floatLen := float64(len(medianSlice))

	seMedian = stdDev / math.Sqrt(floatLen)

	return seMedian

}

//Determine standard error/confidence interval for the bootstrapped statistic from the bootstrapped distribution

func main() {

	// create a random generator to use in all the functions
	rand.New(mt19937.New())

	//var sample_size = []int{100, 1000, 10000, 100000}

	//var num_boots = []int{100, 1000, 10000, 100000}

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	testSlice := randSlice(100, 10)
	testSampled := sampledSlice(testSlice)
	testBoots := boots(testSampled, 100)
	testMedians := bootsMedian(testBoots)
	testSE := medianSE(testMedians)
	fmt.Println("Median SE:", testSE)

	runtime.ReadMemStats(&m2)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

}
