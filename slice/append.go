package main

import (
	"fmt"
	"time"
)

func main() {
	var c int = 100000
	Measure(Append1, c)
	Measure(Append2, c)
	Measure(Append3, c)
}

func Measure(f func(), times int) {
	now := time.Now().UnixNano()
	for i := 0; i < times; i++ {
		f()
	}

	fmt.Println(time.Now().UnixNano() - now)
}

// func append(slice, ...interface{}) sice
func Append1() {
	var i []int      // nil
	i = append(i, 1) // alloc 1
	i = append(i, 2) // aaloc 2, copy , append
	i = append(i, 3) // alloc 4, copy  , append
	i = append(i, 4) // apppend
}

func Append2() {
	var i []int               // nil
	i = append(i, 1, 2, 3, 4) // alloc 4, append
}

func Append3() {
	var i []int = make([]int, 0, 1024)
	i = append(i, 1, 2, 3, 4) // apend
}
