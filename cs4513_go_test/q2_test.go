package cs4513_go_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"project-go-warmup/cs4513_go_impl"
)

func runSumTest(t *testing.T, fileName string, num int, expectedSum int) {
	t.Helper()
	result := cs4513_go_impl.Sum(num, fileName)
	if result != expectedSum {
		t.Fatal(fmt.Sprintf(
			"Sum of %s failed: got %d, expected %d\n", fileName, result, expectedSum))
	}
}

func writeTempIntsFile(t *testing.T, contents string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "nums.txt")
	err := os.WriteFile(path, []byte(contents), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	return path
}

func TestSumSequentialAscending1000(t *testing.T) {
	runSumTest(t, "q2_test1.txt", 1, 499500)
}

func TestSumConcurrentAscending1000(t *testing.T) {
	runSumTest(t, "q2_test1.txt", 10, 499500)
}

func TestSumSequentialMixed1000(t *testing.T) {
	runSumTest(t, "q2_test2.txt", 1, 117652)
}

func TestSumConcurrentMixed1000(t *testing.T) {
	runSumTest(t, "q2_test2.txt", 10, 117652)
}

func TestSumSequentialAscending2000(t *testing.T) {
	runSumTest(t, "q2_test3.txt", 1, 617152)
}

func TestSumConcurrentAscending2000(t *testing.T) {
	runSumTest(t, "q2_test3.txt", 10, 617152)
}

func TestSumSequentialAscending10000(t *testing.T) {
	runSumTest(t, "q2_test4.txt", 1, 4995000)
}

func TestSumConcurrentAscending10000(t *testing.T) {
	runSumTest(t, "q2_test4.txt", 10, 4995000)
}

func TestSumSequentialAscending100000(t *testing.T) {
	runSumTest(t, "q2_test5.txt", 1, 49950000)
}

func TestSumConcurrentAscending100000(t *testing.T) {
	runSumTest(t, "q2_test5.txt", 10, 49950000)
}

func TestSumEmptyFile(t *testing.T) {
	fileName := writeTempIntsFile(t, "")
	runSumTest(t, fileName, 1, 0)
}

func TestSumMoreWorkersThanNumbers(t *testing.T) {
	fileName := writeTempIntsFile(t, "4 5 6")
	runSumTest(t, fileName, 10, 15)
}

func TestSumSingleNegativeValue(t *testing.T) {
	fileName := writeTempIntsFile(t, "-7")
	runSumTest(t, fileName, 3, -7)
}
