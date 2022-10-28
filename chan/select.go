package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		Select1()
	}
}

func Select1() {
	c1 := make(chan string, 1)
	c2 := make(chan int, 1)

	c1 <- "hello"
	c2 <- 1

	select {
	case <-c1:
		fmt.Println("receive string value")
	case <-c2:
		fmt.Println("receive int value")
	}
}
