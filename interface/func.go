package main

import (
	"fmt"
	"reflect"
)

func main() {
	PrintType(int(1))
	PrintType("1234567")
	PrintType([]byte("1234567"))

	a := []int{1, 2, 3, 4, 5, 6}
	// PrintMulInterface(int(1), int(2), int(3))

	PrintMulInterface2(a)
}

func PrintType(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Println("bool")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case complex64:
		fmt.Println("comlex64")
	case complex128:
		fmt.Println("complex128")
	case int:
		fmt.Println("int")
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case int32:
		fmt.Println("int32")
	case int64:
		fmt.Println("int64")
	case uint:
		fmt.Println("uint")
	case uint8:
		fmt.Println("uint8")
	case uint16:
		fmt.Println("uint16")
	case uint32:
		fmt.Println("uint32")
	case uint64:
		fmt.Println("uint64")
	case uintptr:
		fmt.Println("uintptr")
	case string:
		fmt.Println("string")
	case []byte:
		fmt.Println("[]byte")
	case reflect.Value:
		fmt.Println("reflect.Value")
	default:
		fmt.Println("unknowed type")
	}
}

func PrintMulInterface(i []interface{}) {
	for _, v := range i {
		PrintType(v)
	}
}

func PrintMulInterface2(i ...interface{}) {
	// for _, v := range i {
	// 	PrintType(v)
	// }
	PrintMulInterface(i)
}

func PrintMulInterfac3(i *interface{}) {
}
