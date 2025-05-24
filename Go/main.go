package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"time"

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

// Resample the data with replacement n times and send with a channel
func boots(ch chan []float64, floatSlice []float64, n int) {

	for i := 0; i < n; i++ {

		// create a new slice with replacement
		sample := sampledSlice(floatSlice)

		// add it to the result
		ch <- sample
	}
	close(ch)
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
// Receivedd channel
func bootsMedian(ch chan []float64) (medianSlice []float64) {

	for data := range ch {

		medianVal := median(data)
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

func main() {

	// create a random generator to use in all the functions
	rand.New(mt19937.New())

	var sample_size = []int{10, 100, 1000, 10000}

	var num_boots = []int{10, 100, 1000, 10000}

	for _, sample := range sample_size {

		for _, nums := range num_boots {

			writeLine("-------------", "Go_result_log.txt")

			title := fmt.Sprintf("processing %v size and %v boot samples", sample, nums)
			writeLine(title, "Go_result_log.txt")
			fmt.Println(title)

			start := time.Now()

			var m1, m2 runtime.MemStats
			runtime.GC()
			runtime.ReadMemStats(&m1)

			dataChannel := make(chan []float64)

			testSlice := randSlice(sample, 10)
			go boots(dataChannel, testSlice, nums)
			testMedians := bootsMedian(dataChannel)
			testSE := medianSE(testMedians)

			runtime.ReadMemStats(&m2)

			elapsed := time.Since(start)

			line1 := fmt.Sprintf("Sample Size: %v Boot Samples: %v", sample, nums)
			writeLine(line1, "Go_result_log.txt")

			line2 := fmt.Sprintf("Median Standard Error:  %v", testSE)
			writeLine(line2, "Go_result_log.txt")

			totAlloc := m2.TotalAlloc - m1.TotalAlloc
			line3 := fmt.Sprintf("Cumulative Bytes Allocated: %v", totAlloc)
			writeLine(line3, "Go_result_log.txt")

			heapObj := m2.Mallocs - m1.Mallocs
			line4 := fmt.Sprintf("Cumulative Heap Objects Allocated: %v", heapObj)
			writeLine(line4, "Go_result_log.txt")

			line5 := fmt.Sprintf("Time to Compute: %v", elapsed)
			writeLine(line5, "Go_result_log.txt")

			writeLine("-------------", "Go_result_log.txt")

		}
	}

	fmt.Println("processing complete")

}
