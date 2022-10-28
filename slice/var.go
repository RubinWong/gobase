package main

import "fmt"

func main() {
	var (
		// 默认值创建
		a []int            // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
		b = []int{}        // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
		c = []int{1, 2, 3} // 有3个元素的切片, len和cap都为3

		// 从数组或者其它slice创建
		d = c[:2]         // 有2个元素的切片, len为2, cap为3
		e = c[0:2:cap(c)] // 有2个元素的切片, len为2, cap为3
		f = c[:0]         // 有0个元素的切片, len为0, cap为3

		// make创建
		g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
		i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	)
	PrintSlice("a", a)
	PrintSlice("b", b)
	PrintSlice("c", c)
	PrintSlice("d", d)
	PrintSlice("e", e)
	PrintSlice("f", f)
	PrintSlice("g", g)
	PrintSlice("g", h)
	PrintSlice("i", i)

	var j [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var k []int = j[4:6]
	PrintSlice("k", k)

	var l []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var m []int = l[4:6]
	PrintSlice("m", m)
}

func PrintSlice(n string, i []int) {
	fmt.Print(n, " ")
	fmt.Print("nil: ", i == nil, "\tlen: ", len(i), "\tcap: ", cap(i), "\t value: ")
	for _, v := range i {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
