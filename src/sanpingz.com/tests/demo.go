package main

import (
		"fmt"
		"time"
		"runtime"
	//	"tests"
		)

var c chan int

func foo() (ret int) {
	defer func() {
		ret ++
	}()
	return 0
}

func bar(args ...int) {
	for i, n := range args {
		fmt.Printf("args[%d] = %d\n", i, n)
	}
}

func long_loop(n int) (count float64) {
	for i:=1; i<=n; i++ {
		count += float64(i)
	}
	println(count)
	c <- 1
	return
}

func main() {
	runtime.GOMAXPROCS(4)
	c = make(chan int)
	/*
	fmt.Printf("Hello, world.\n")
	fmt.Printf("Sqrt(100) = %v.\n", tests.Sqrt(100))
	for i:=0; i<=10; i++ {
		println(i*(i+1))	
	}
	defer fmt.Printf("\n")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	println(foo())
	go bar(1,4,2,8,5,7)
	fmt.Println("I'm waiting")
	time.Sleep(3 * time.Second)
	*/
	start := time.Now().UnixNano()
	go long_loop(5000000)
	<-c
	end := time.Now().UnixNano()
	fmt.Printf("Run time: %0.2fms\n", float32(end-start)/1000000)
}
