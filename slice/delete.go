package main

import "fmt"

func main() {
	Delete1()
	Delete2()
	Delete3()
	Delete4()
}

func Delete1() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	b := a[:10]
	PrintSlice("d1", b)

	b = b[:len(b)-3] // 删除尾部N个元素
	PrintSlice("d1", b)

	b = b[3:] // 删除开头N个元素
	PrintSlice("d1", b)
}

func Delete2() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	b := a[:10]
	PrintSlice("d2", b)

	b = append(b[:0], b[:]...) // 删除开头N个元素
	PrintSlice("d2", b)

	b = append(b[:0], b[3:]...) // 删除开头N个元素

	PrintSlice("d2", b)
}

func Delete3() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	b := a[:10]
	PrintSlice("d3", b)

	b = b[:copy(b, b[:len(b)-3])] // 删除尾部N个元素
	PrintSlice("d3", b)

	b = b[:copy(b, b[3:])] // 删除开头N个元素
	PrintSlice("d3", b)
}

func Delete4() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	b := a[:10]
	PrintSlice("d4-1", b)

	b = append(b[:4], b[4+1:]...) // 删除中间第四个元素
	PrintSlice("d4-1", b)

	b = append(b[:4], b[4+3:]...) // 删除中间第四个元素开始的N个元素
	PrintSlice("d4-1", b)

	c := a[:10]
	PrintSlice("d4-2", c)

	c = c[:4+copy(c[4:], c[4+1:])] // 删除中间1个元素
	PrintSlice("d4-2", c)

	c = c[:4+copy(c[4:], c[4+3:])] // 删除中间N个元素
	PrintSlice("d4-2", c)
}

func PrintSlice(n string, i []int) {
	fmt.Print(n, " ")
	fmt.Print("nil: ", i == nil, "\tlen: ", len(i), "\tcap: ", cap(i), "\t value: ")
	for _, v := range i {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
