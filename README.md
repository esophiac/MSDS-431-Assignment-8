# Assignment Week 8 - Modern Applied Statistics with Go
Text

describe your efforts in finding R and Go packages for the method

Review your process of building the Go implementation. What did you do to improve the performance of the Go implementation of the selected statistical emthod?

Review your experiences with testing, benchmarking, software profiling, and logging

## Background
Andrew Gelamn and Aki Vehrati wrote the article, “What Are the Most Important Statistical Ideas of the Past 50 Years?" which provides a list of significant statistical ideas and how they relate to modern computing and exploratory data analysis. Based on the article, bootstrapping was identified as a method of testing the performance of Go and R due to its significance in  modern applied statistics. 

Full source: Gelman, Andrew, and Aki Vehtari. 2022. “What Are the Most Important Statistical Ideas of the Past 50 Years?" *Journal of the American Statistical Association*, 116: 2087–2097. Available online at [https://arxiv.org/abs/2012.00174](https://arxiv.org/abs/2012.00174).

### R Programming Tools
To demonstrate the validity of the packages selected below, the R_week_8 program runs bootstrap sampling on a normal distribution, left-skew, and right-skew. Once this is completed, the program moves on to perform the same analysis that the main.go program performs; logging the performance of the program at different sample sizes and numbers of bootstrap samples. The full logs for the R program can be found at R_logs.txt in the repository.

**peakRAM** is a library that is used to track the memory usage of R programs. For information about the package can be found at [https://cran.r-project.org/web/packages/peakRAM/index.html](https://cran.r-project.org/web/packages/peakRAM/index.html). **Note:** The memory usage reported by this package is in MiB.

**boot** is a library that is used to perform functions and analysis for bootstrap sampling, based on the the book "Bootstrap Methods and Their Application" by A. C. Davison and D. V. Hinkley (1997, CUP), originally written by Angelo Canty for S. More information about the library can be found at [https://cran.r-project.org/web/packages/boot/index.html](https://cran.r-project.org/web/packages/boot/index.html).

### Go Programming Tools

I used the gonum package [here](gonum.org/v1/gonum/stat) instead of previous packages explored in this course because it had more statistics that I was interested in calculating.

For random number generation, I used the Mersenne Twister in Go package at [https://github.com/seehuhn/mt19937](https://github.com/seehuhn/mt19937) to ensure that the random numbers generate in Go used the same algorithm as the numbers generated in the R program. This also ensures that an comparisons drawn from memory usage will be comparable.

All other packages were from the standard library, which included: bufio, fmt, math, math/rand, os, runtime, sort, and time.

## Recommendation to Management
Overall, I would recommend that management proceed forward with using Go in place of R because of the improvements to speed and memory usage. The comparison of the time it takes to run the program for an initial sample size of 100 with varying boot samples is below.

| Boot Samples | R | Go |
|----------|-----------|------------|
| 10 | <0.01 sec | <0.01 sec |
| 100 | <0.01 sec | 0.001 sec|
| 1000 | 0.03 sec | 0.005 sec |
| 100000 | 0.25 sec | 0.037 sec |

One of the main drivers of the improved performance in Go is the use of goroutines and channels to handle the bootstrap sampling. Additional savings could be possible if reasonable deployment of concurreny is made an enterprise standard.

How easy is it to use Go in place of R? How much money will the firm save in cloud computing costs?

Under what circumstances would it make sense for the firm to use Go in place of R for the selected statistical method?

Select a cloud provider of infrastructure as a service (IaS). Note the cloud costs for virtual machine (compute engine) services. What percentage of cloud computing costs might be saved with a move from R to Go?

## Roles of Programs and Data
These are the programs in the repository. Data information is in the background section.

- Go
    - main.go: runs the bootstrapping program
    - go.mod: defines the module's properties
    - go.sum: record of the library the project depends on
    - main_test.go: tests and benchmarks the fuctions in the main.go file
    - Go_result_log.txt: outputs the logs that will demonstrate Go's functionality for bootstrap sampling
    - assignment_8: the application for this program. It will generate the logs without requiring user input.
- R
    - R_logs.txt: logs for exploring the memory and time requirements of bootstrapping
    - R_week_8.R: R file used to explore the memory and time requirements of bootstrapping
- README.md: the readme file for the repository

## Application
An executable for this project was created using Windows. To create your own executable, run **go build** in the same directory as the go program. For more information, see the Gopher documentation on creating an executable [here](https://go.dev/doc/tutorial/compile-install).

## Use of AI
AI was not used for this assignment.