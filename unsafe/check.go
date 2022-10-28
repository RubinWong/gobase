package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(fmt.Sprintf("0x%X", (uintptr)(unsafe.Pointer(&a))))

	b := a[:10]
	fmt.Println(fmt.Sprintf("0x%X", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))
	PrintSlice(b, "m")

	Append1()

	Append2()

	Append3()
}

func PrintSlice(i []int, tag string) {
	fmt.Println(tag, " ", fmt.Sprintf("0x%X", (*reflect.SliceHeader)(unsafe.Pointer(&i)).Data))
}

func Append1() {
	var i []int
	i = append(i, 1)
	PrintSlice(i, "A1")

	i = append(i, 2)
	PrintSlice(i, "A1")

	i = append(i, 3)
	PrintSlice(i, "A1")

	i = append(i, 4)
	PrintSlice(i, "A1")

	i = append(i, 5)
	PrintSlice(i, "A1")
}

func Append2() {
	var i []int
	i = append(i, 1, 2, 3, 4)
	PrintSlice(i, "A2")

	i = append(i, 5, 6, 7, 8)
	PrintSlice(i, "A2")
}

func Append3() {
	var i []int = make([]int, 4)
	i = append(i, 1)
	PrintSlice(i, "A3")

	i = append(i, 2)
	PrintSlice(i, "A3")

	i = append(i, 3)
	PrintSlice(i, "A3")

	i = append(i, 4)
	PrintSlice(i, "A3")

	i = append(i, 5)
	PrintSlice(i, "A3")
}
