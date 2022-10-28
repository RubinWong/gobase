package main

import (
	"fmt"
	"reflect"
)

type any = interface{}

func main() {
	var a interface{}
	a = 1
	fmt.Println(a, reflect.TypeOf(a).Name())

	a = "hello world"
	fmt.Println(a, reflect.TypeOf(a).Name())

	a = 3.14159267
	fmt.Println(a, reflect.TypeOf(a).Name())
}
