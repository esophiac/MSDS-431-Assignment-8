# Assignment Week 8 - Modern Applied Statistics with Go
For this assignment, I chose to implement bootstrap sampling and find the standard error of the median. I first created the program in R to validate the methodology for bootstrap sampling with different shapes of sample. Then I tested how the sample size and bootstrap sampling impacts the memory usage and runtime of the programs in R anf Go.

*Describe your efforts in finding R and Go packages for the method.*
R has specific packages already created for handling bootstrap sampling. Go did not. In fact, one of the decision points for the Go program was to determine which statistics package to use. I ultimately decided on **gonum** because it had the greatest number of functions I was interested in using, but none of them were completely holistic. However, Go had much better benchmarking and logging capailities. I eventually found **peakRAM** for R, but it took a while.

*Review your process of building the Go implementation. What did you do to improve the performance of the Go implementation of the selected statistical method?*
The big performance changer for my Go program was implementing channels and goroutines. When I first created the program, it did not have any concurrency, and so ran much slower than the R program. I decided to implement the go routine on the bootstrapping itself and sending the resulting samples to be computed into a median.

*Review your experiences with testing, benchmarking, software profiling, and logging.*
Logging was one of the main goals of both of the programs because I wanted to make sure I could accurately compare the two programs. R does not have as robust logging and memory tracking; it records everything in MiB while Go recorded everything in bytes.

## Recommendation to Management
Overall, I would recommend that management proceed forward with using Go in place of R because of the improvements to speed and memory usage. The comparison of the time it takes to run the program for an initial sample size of 100 with varying boot samples is below.

| Boot Samples | R | Go |
|----------|-----------|------------|
| 10 | <0.01 sec | <0.01 sec |
| 100 | <0.01 sec | 0.001 sec|
| 1000 | 0.03 sec | 0.005 sec |
| 100000 | 0.25 sec | 0.037 sec |

One of the main drivers of the improved performance in Go is the use of goroutines and channels to handle the bootstrap sampling. Additional savings could be possible if reasonable deployment of concurreny is made an enterprise standard.

*How easy is it to use Go in place of R? How much money will the firm save in cloud computing costs?*
Go is easier to use than R because of the building-blocks nature that the language encourages. Given that many costs are determined by time using resources, we estimate that the company could reduce costs by almost 85.2% by switching to go. This comparison was made by looking at how much time it takes to process a 100-size sample with 100000 bootstrap samples.

*Under what circumstances would it make sense for the firm to use Go in place of R for the selected statistical method?*
It would make sense for the firm to use Go in cases where analysis can be done on discrete units that do not require information from the step before them to proceed. It is important to note that concurrency is not parallelism, though the shift to cloud resources may give the firm an opportunity to take advantage of both paradigms.

*Select a cloud provider of infrastructure as a service (IaS). Note the cloud costs for virtual machine (compute engine) services. What percentage of cloud computing costs might be saved with a move from R to Go?*
Based on the pricing guide by [Google Cloud](https://cloud.google.com/compute/vm-instance-pricing#section-1), a standard work-week (40 hours without overtime) for one person of computing would cost $11.09. If Go is able to deliver on the 85.2% reduction in cost, then the company would save $9.44 every week.

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