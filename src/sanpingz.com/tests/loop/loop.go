package main

import (
		"fmt"
		"time"
		)

func main() {
	N := 1000000000
	i := 0
	sum := 0
	start := time.Now().UnixNano()
	for {
		i++
		sum += i
		if i == N {
			break
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("(Go) Run time for %d loop: %0.2fms\n", N, float32(end-start)/1000000)
}
