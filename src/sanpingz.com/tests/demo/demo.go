package main

import (
	"fmt"
	//. "sanpingz.com/simplemath"
	//"time"
)

const (
	PI = 3.14159265
	v1 = 2 << iota
	v2 = iota
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 这个常量没有导出
)

type PersonInfo struct {
	ID    string
	Name  string
	Email string
}

type Integer int

func (a Integer) less(b Integer) bool {
	return a < b
}

func (a *Integer) add(b Integer) Integer {
	*a += b
	return *a
}

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("I'm Foo")
}

func (base *Base) Bar() {
	fmt.Println("I'm Bar")
}

type NBase struct {
	Base
}

func (nbase *NBase) Bar() {
	nbase.Base.Bar()
	fmt.Println("I'm new added Bar")
}

func strloop(str string) {
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}

func main() {
	/*
		fmt.Printf("Hello world\n")
		fmt.Printf("Sqrt(%v) = %v\n", 142857, Sqrt(142857))
		fmt.Printf("PI = %v, v1 = %v, v2 = %v, Monday = %v\n", PI, v1, v2, Monday)
		strloop("Hello, 世界")
		var (
			arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			sli []int   = arr[:5]
		)
		for _, v := range sli {
			fmt.Print(v, " ")
		}
		fmt.Println()

		sli = make([]int, 5, 10)
		sli = append(sli, 1, 2, 3)
		fmt.Printf("len(sli) = %d, cap(sli) = %d\n", len(sli), cap(sli))
		var imap map[string]PersonInfo
		imap = make(map[string]PersonInfo)
		imap["1001"] = PersonInfo{"1001", "Calvin", "calvin.zhang@163.com"}
		imap["1002"] = PersonInfo{"1002", "Andy", "andy.zhang@163.com"}
		imap["1003"] = PersonInfo{"1003", "Kevin", "kevin.zhang@163.com"}

		if person, ok := imap["1002"]; ok {
			fmt.Println(person)
		}

		N := 1000000000
		start := time.Now().Unix()
		sum := 0
		for {
			sum++
			if sum == N {
				break
			}
		}
		fmt.Printf("(Go) Runtime for %v loop: %ds\n", sum, time.Now().Unix()-start)
		foo := func() func() {
			var i int = 100
			return func() {
				fmt.Printf("Closure: N = %d, i = %d\n", N, i)
			}
		}()

		foo()
		defer foo()
		N = 10000
		foo()

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Runtime error caught: %v\n", r)
			}
		}()
			panic(404)
			fmt.Println("exit?\n")
	*/

	var (
		a Integer = 1
		b Integer = 2
	)
	if a.less(b) {
		fmt.Println(a, "less than", b)
	}
	a.add(b)
	fmt.Println("a =", a, ", a.add(b) =", a.add(b))

	var (
		base  Base
		nbase NBase
	)
	base.Bar()
	nbase.Bar()
	nbase.Foo()
}
