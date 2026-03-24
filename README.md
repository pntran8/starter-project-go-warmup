# CS4513: Warmup The Go Basic

## Student

Phuong Tran (pntran1@wpi.edu)

## Two Problems

1. What is the difference between unbuffered and buffered channels? And why do you choose one over the other for this assignment?

Unbuffered channels are when you don't set an initial capacity. So `channelUnbuff := make(chan int)`. This means that it can't hold any values and that adding in a value blocks until a goroutine is ready to take it. Also, taking from the channel is a block until there's a value set inside. So it really only works if the goroutine is ready to put something in AND if something is ready to take out that value.

A buffered channel is when you do set an initial capacity. So `channelBuff := make(chan int, n)`, where n is a number. This means that the channel can store up to n things. Adding stuff to the channel only blocks if the buffer is full. And taking from the channel only blocks if the buffer is empty. This mean that a goroutine can leave something in the channel for later.

For this assignment, we're asked to create multiple goroutines that each sum up numbers from a num channel. So that means that each worker can send part of the total sum to the out channel. If one goroutine finishes before another, we don't want to wait for all goroutines to complete before this one can put something in the channel. Which means we want something where the goroutine can send stuff to the channel without relying on something being ready to receive it. We also want to put all of our numbers into the channel at once because this keeps the goroutines busy so they don't have to wait around for each one.

For the out channel, a buffer is also useful because the goroutines can put their sums into the channel without having to wait for the main goroutine to receive it.

If we wanted each goroutine to wait until the main goroutine is ready to read, then we would use unbuffered because it blocks the other goroutines from doing stuff to the channel before the main one is ready.

2. Briefly explain how you approached the two problems.

For the Top K Words problem, I checked if the file is empty before opening it. I converted the file text into lower case and removed the non alphanumeric characters before splitting them into words. I think I also could've done it the other way, where I sliced it and then lower case & regexp it, but this method would handle puncuation and stuff more consistently since it does it all together instead of separately. Then I would just loop through the slices and make sure there's no puncuations and that it's longer than my charThreshold before adding it to my "counts" map, which keeps track of what words appeared how many times.

Then I sorted it and returned that.

For the Parallel Sum one, I did the Sum func first. I made the numsChan and outChan, which is the channel to put all my numbers in and the channel to put all my sums in. Then I made the numsChan worker goroutine to sum the numbers from nums until it couldn't anymore and put the partial sums into outChan. Then I put all numbers into numsChan and closed it. Then I added all the sums from outChan.

And then I made the sumWorker func to add numbers.

## Measurement

Running workload file: ./cs4513_go_test/q2_test1.txt  
result: 499500, num of workers: 1, avg time over 100 runs: 0.00067856 seconds  
result: 499500, num of workers: 2, avg time over 100 runs: 0.00044842 seconds  
result: 499500, num of workers: 3, avg time over 100 runs: 0.00056960 seconds  
result: 499500, num of workers: 5, avg time over 100 runs: 0.00050429 seconds  
result: 499500, num of workers: 10, avg time over 100 runs: 0.00039877 seconds  
Running workload file: ./cs4513_go_test/q2_test2.txt  
result: 117652, num of workers: 1, avg time over 100 runs: 0.00033575 seconds  
result: 117652, num of workers: 2, avg time over 100 runs: 0.00031844 seconds  
result: 117652, num of workers: 3, avg time over 100 runs: 0.00031041 seconds  
result: 117652, num of workers: 5, avg time over 100 runs: 0.00032537 seconds  
result: 117652, num of workers: 10, avg time over 100 runs: 0.00031060 seconds  
Running workload file: ./cs4513_go_test/q2_test3.txt  
result: 617152, num of workers: 1, avg time over 100 runs: 0.00055195 seconds  
result: 617152, num of workers: 2, avg time over 100 runs: 0.00051126 seconds  
result: 617152, num of workers: 3, avg time over 100 runs: 0.00055697 seconds  
result: 617152, num of workers: 5, avg time over 100 runs: 0.00057793 seconds  
result: 617152, num of workers: 10, avg time over 100 runs: 0.00055885 seconds  
Running workload file: ./cs4513_go_test/q2_test5.txt  
result: 49950000, num of workers: 1, avg time over 100 runs: 0.02323286 seconds  
result: 49950000, num of workers: 2, avg time over 100 runs: 0.02255467 seconds  
result: 49950000, num of workers: 3, avg time over 100 runs: 0.02305630 seconds  
result: 49950000, num of workers: 5, avg time over 100 runs: 0.02396701 seconds  
result: 49950000, num of workers: 10, avg time over 100 runs: 0.02315370 seconds  
Running workload file: ./cs4513_go_test/q2_test5.txt  
result: 49950000, num of workers: 1, avg time over 100 runs: 0.02323192 seconds  
result: 49950000, num of workers: 2, avg time over 100 runs: 0.02297677 seconds  
result: 49950000, num of workers: 3, avg time over 100 runs: 0.02306543 seconds  
result: 49950000, num of workers: 5, avg time over 100 runs: 0.02484330 seconds  
result: 49950000, num of workers: 10, avg time over 100 runs: 0.02411801 seconds

## Observations and Explanations

Regardless of how many workers there are, it all got the same results, so that means that the functions are doing what we expect them to do. Increasing the number of workers doesn't always decrease the average time, because sometimes it even increases it slightly. This might be because for smaller files, creating more goroutines and channels take up the majority of the time, so it's not really worth it to make that many goroutines/channels. But for larger files, more workers helped to decrease the average time by a little bit, but then at a certain point, having more workers didn't really help a lot. It might be because by that point, the work per goroutine isn't a lot so adding more workers isn't worth it again.

For additional measurement scenarios, measuring really really large files might be worth a try, because then you can see how the number of goroutines change as the file size substantially increases.

## Errata

In multiple spots I had to change it from `project-go-warmup` to `starter-project-go-warmup`, like in `q1_test.go` line 12 & `q2_test.go` line 9 & also redo the `go.mod` file to `starter-project-go-warmup`.

Also, I was able to run `go run .` on my local just fine, but the autograder fails to do so.