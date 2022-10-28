package main

import (
	"fmt"
	"reflect"
)

type A struct{}

func main() {
	var a interface{}
	a = 1
	i, ok := a.(int)
	fmt.Println(i, ok)
	Type(a)
	Type2(a)

	a = "hello world"
	s := a.(string)
	fmt.Println(s)
	Type(a)
	Type2(a)

	a = 3.14159267
	f, ok := a.(int)
	fmt.Println(f, ok)
	Type(a)
	Type2(a)
}

func Type(i interface{}) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.String:
		fmt.Println("string kind")
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("number kind")
	case reflect.Float32, reflect.Float64:
		fmt.Println("float type")
	}
}

func Type2(i interface{}) {
	switch i.(type) {
	case int, int16, int32, int64:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case float32, float64:
		fmt.Println("float type")
	case A:
		fmt.Println("struct A type")
	}
}
