package main

import "fmt"

func main() {
	// Deadlock()
	Deadlock2()
}

func Deadlock() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("successfully send")
}

func Deadlock2() {
	ch := make(chan int, 1)
	ch <- 10
	ch <- 20
	fmt.Println("successfully send")
}
