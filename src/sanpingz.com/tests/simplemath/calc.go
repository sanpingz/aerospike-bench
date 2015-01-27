package simplemath

import "fmt"

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum // send results to channel
}

func C3sum() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan // receive results
	fmt.Println("Result:", sum1, sum2, sum1+sum2)
}
