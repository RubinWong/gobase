package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	Cast()

	ModifySlice()

	ModifyStruct()
}

func Cast() {
	var u64 uint64 = 200
	// var i int64 = u64
	// var i64 *int64 = &u64
	var i64 *int64 = (*int64)(unsafe.Pointer(&u64))
	fmt.Println(u64, *i64)
	*i64 = 100
	fmt.Println(u64, *i64)
}

func ModifySlice() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// s := "1234567890"   //  8bit  1byte
	b := a[:10]

	// reflect.StringHeader
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	i := (*int)(unsafe.Pointer(p.Data + 8))
	*i = 110
	PrintSlice("i", b)
	PrintSlice("i", a[:])
}

func ModifyStruct() {
	type Programmer struct {
		name     string
		age      int
		language string
	}

	p := Programmer{
		name:     "Airce",
		age:      35,
		language: "CPP",
	}

	name := (*string)(unsafe.Pointer(&p))
	*name = "Zeus"
	age := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p)) + unsafe.Sizeof(string(""))))
	*age = 25

	language := (*string)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p)) + unsafe.Sizeof(string("")) + unsafe.Sizeof(int(0))))
	*language = "golang"

	fmt.Println(p)
}

// func Cast2() {
// 	s := "12334578"
// 	b := []byte(s)
// }

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	sliceHeader := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&stringHeader))
}

func PrintSlice(n string, i []int) {
	fmt.Print(n, " ")
	fmt.Print("nil: ", i == nil, "\tlen: ", len(i), "\tcap: ", cap(i), "\t value: ")
	for _, v := range i {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
