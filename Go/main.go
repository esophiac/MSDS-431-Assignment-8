package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
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

// Write a string to a specific file
func writeLine(newText string, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)

	writer.WriteString(newText)
	writer.WriteString("\n")

	writer.Flush()

	return nil
}

//Determine standard error/confidence interval for the bootstrapped statistic from the bootstrapped distribution

func main() {

	// create a random generator to use in all the functions
	rand.New(mt19937.New())

	var sample_size = []int{100, 1000, 10000, 100000}

	var num_boots = []int{100, 1000, 10000, 100000}

	for _, sample := range sample_size {

		for _, nums := range num_boots {

			fmt.Println("processing %v size and %v boot samples", sample, nums)

			var m1, m2 runtime.MemStats
			runtime.GC()
			runtime.ReadMemStats(&m1)

			testSlice := randSlice(sample, 10)
			testBoots := boots(testSlice, nums)
			testMedians := bootsMedian(testBoots)
			testSE := medianSE(testMedians)

			runtime.ReadMemStats(&m2)

			line1 := "Sample Size: " + string(sample) + " Boot Samples: " + string(nums) + "\n"
			writeLine(line1, "Go_result_log.txt")

			line2 := "Median Standard Error: " + string(int64(testSE))
			writeLine(line2, "Go_result_log.txt")

			line3 := "Cumulative Bytes Allocated: " + string(m2.TotalAlloc-m1.TotalAlloc)
			writeLine(line3, "Go_result_log.txt")

			line4 := "Cumulative Heap Objects Allocated: " + string(m2.Mallocs-m1.Mallocs)
			writeLine(line4, "Go_result_log.txt")

		}
	}

	fmt.Println("processing complete")

}
