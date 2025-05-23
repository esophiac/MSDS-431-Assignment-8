## import packages to use
library(boot)
library(peakRAM)

## create functions to use in boot to calculate mean and median
## mean
samplemean <- function(x, d) {
  return(mean(x[d]))
}

## median
samplemedian <- function(x, d) {
  return(median(x[d]))
}


## demonstrate that the methods work for different distributions and log the results

## creating a new file and writing to it
filename <- 'R_logs.txt'
wd <- getwd()
print(cat("Printing", filename, "to", wd))

cat("R Log Information", file=filename, append=FALSE, sep='\n')
cat("\n", file=filename, append=TRUE, sep='')

i = 1

## set up the data to be used to demonstrate that the boot method works
## population mean: 10, standard deviation: 5, size: 100
set.seed(123)
data_norm <- round(rnorm(100, 10, 5))

## create data with a left skew
data_left <- round(rbeta(100,5,1)*10)

## create data with a right skew
data_right <- round(rbeta(100,1,5)*10)

## list of the data to iterate through later
data_list <- list(data_norm, data_left, data_right)
data_type <- c("normalized", "left-skew", "right-skew")

## iterate through all of the data distributions and write the results to a file
i = 1

for (data in data_list) {
  print(cat("Printing for sample size = ", data_type[[i]]))
  cat("-------------", file=filename, append=TRUE, sep='\n')
  cat("Bootstrapping for ", data_type[[i]], file=filename, append=TRUE, sep='\n')
  
  ## calculate the RAM and time it takes for each function to run
  mem1 <- peakRAM({
    ## running the boot function with the data generated - mean with 100 replications
    boot_mean <- boot(data, samplemean, R=100)
    calc_mean <- boot_mean$t0
    cat("Bootstrapping Mean:", calc_mean, file=filename, append=TRUE, sep='\n')
    ## std error of mean = standard deviation divided by the square root of the sample size
    calc_mean_SE <- sd(boot_mean$t) / 10 	
    cat("Bootstrapping Mean SE:", calc_mean_SE, file=filename, append=TRUE, sep='\n')
  })
  
  
  ## write RAM info to file
  cat("RAM to Calculate Mean and SE:", mem1$Total_RAM_Used_MiB*(1048576), file=filename, append=TRUE, sep='\n')
  cat("Peak RAM to Calculate Mean and SE:", mem1$Peak_RAM_Used_MiB*(1048576), file=filename, append=TRUE, sep='\n')
  cat("Time: ", mem1$Elapsed_Time_sec, file=filename, append=TRUE, sep='\n')
  
  ## calculate the RAM and time it takes for each function to run
  mem2 <- peakRAM({
    ## running the boot function with the data generated - median with 100 replications
    b_median <- boot(data, samplemedian, R=100)
    calc_median <- b_median$t0
    cat("Bootstrapping Median:", calc_median, file=filename, append=TRUE, sep='\n')
    calc_median_SE <- sd(b_median$t) / 10
    cat("Bootstrapping Median SE:", calc_median_SE, file=filename, append=TRUE, sep='\n')
  })
  
  ## write RAM info to file
  cat("RAM to Calculate Media and SE: ", mem2$Total_RAM_Used_MiB*(1048576), file=filename, append=TRUE, sep='\n')
  cat("Peak RAM to Calculate Media and SE:", mem2$Peak_RAM_Used_MiB*(1048576), file=filename, append=TRUE, sep='\n')
  cat("Time: ", mem2$Elapsed_Time_sec, file=filename, append=TRUE, sep='\n')
  
  cat("-------------", file=filename, append=TRUE, sep='\n')
  
  i = i + 1
}

## iterate through different sample sizes and number of bootstrap samples to get memory and time requirments
sample_size <- c(10, 100, 1000, 10000)
num_boots <- c(10, 100, 1000, 10000)

for (size in sample_size) {
  print(cat("Printing for sample size = ", size))
  ## create a normally distrbuted, random dataset
  ## population mean: 10, standard deviation: 5, size: 100
  set.seed(123)
  data_size <- round(rnorm(size, 10, 5))
  
  ## select a new boot size
  for (numb in num_boots) {
    ## new entry in log
    cat("Sample size:", size, "Number of Boot Samples:", numb, file=filename, append=TRUE, sep=' ')
    
    mem3 <- peakRAM({
      ## running the boot function with the data generated - median with numb replications
      b_median <- boot(data, samplemedian, R=numb)
      calc_median <- b_median$t0
      cat("Bootstrapping Median:", calc_median, file=filename, append=TRUE, sep='\n')
      calc_median_SE <- sd(b_median$t) / sqrt(size) 
      cat("Bootstrapping Median SE:", calc_median_SE, file=filename, append=TRUE, sep='\n')
    })
    
    ## write RAM info to file
    cat("RAM to Calculate Median and SE:", mem3$Total_RAM_Used_MiB, file=filename, append=TRUE, sep='\n')
    cat("Peak RAM to Calculate Median and SE:", mem3$Peak_RAM_Used_MiB, file=filename, append=TRUE, sep='\n')
    cat("Time: ", mem3$Elapsed_Time_sec, file=filename, append=TRUE, sep='\n')
    
    cat("-------------", file=filename, append=TRUE, sep='\n')
    
  }
  
}

print(cat("Printing", filename, "to", wd))
