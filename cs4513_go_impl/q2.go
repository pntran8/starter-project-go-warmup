package cs4513_go_impl

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

/*
Do NOT modify function signature.

Sum numbers from channel `nums` and output sum to `out`.
You should only output to `out` once.
*/
func sumWorker(nums chan int, out chan int) { 
	// nums is an int channel and out is an int channel
	// TODO: implement me
	// HINT: use for loop over `nums`

	//every worker adds together the numbers in their numChan and then puts the sum in their outChan

	sum := 0

	for i := range nums { // loops over the stuff in nums channel
		sum += i
	}

	out <- sum // puts the sum in the "out" channel
}

/*
Do NOT modify function signature.

Read integers from the file `fileName` and return sum of all values.
This function must launch `num` go routines running `sumWorker` to find the sum of the values concurrently.

You should use `checkError` to handle potential errors.
*/
func Sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorker`
	// HINT: use buffered channels for splitting numbers between workers
	
	// read from OS
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	// read vals
	vals, err := readInts(file)
	checkError(err)

	if num <= 0 {
		return 0
	}

	numsChan := make(chan int, num) //nums channel
	outChan := make(chan int, num) //out channel

	for i := 0; i < num; i++ { // loop over num to sum
		go sumWorker(numsChan, outChan) //make a num worker to add the numbers
	}

	for _, j := range vals {
		numsChan <- j //adds "j" to the numbers channel
	}

	close (numsChan)

	total := 0
	for i := 0; i < num; i++ {
		total += <- outChan //adds from the "out" channel
	}

	return total
}

/*
Do NOT modify this function.
Read a list of integers separated by whitespace from `r`.
Return the integers successfully read with no error, or
an empty slice of integers and the error that occurred.
*/
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
