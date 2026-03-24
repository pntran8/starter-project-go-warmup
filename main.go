package main

import (
	"fmt"
	impl "github.com/pntran8/starter-project-go-warmup/cs4513_go_impl"
	"time"
)

func main() {
	// TODO: implement me
	// HINT: need to import the cs4513_go_impl package
	// HINT: use the time package for measurement

	files := []string{ //specify the test files
		"./cs4513_go_test/q2_test1.txt",
		"./cs4513_go_test/q2_test2.txt",
		"./cs4513_go_test/q2_test3.txt",
		"./cs4513_go_test/q2_test5.txt",
		"./cs4513_go_test/q2_test5.txt",
	}

	goroutines := []int{ 1, 2, 3, 5, 10} //make the num of goroutines
	runs := 100 // num of timed runs

	// need to loop over the files and goroutines
	for _, file := range files {
		fmt.Println("Running workload file: ", file)

		// now need to loop over the goroutine counts
		for _, num := range goroutines {
			// warm up run (not timed one)
			result := impl.Sum(num, file)

			// timed runs
			start := time.Now()
			for i := 0; i < runs; i++ {
				_ = impl.Sum(num, file)
			}
			end := time.Since(start)

			avgTime := end.Seconds() / float64(runs) //avg time is this/num of timed runs

			//print out
			fmt.Printf("result: %d, num of workers: %d, avg time over 100 runs: %.8f seconds\n", result, num, avgTime)
		}
	}
}