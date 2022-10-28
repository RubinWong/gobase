package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c chan int = make(chan int, 10)
	done := make(chan bool)
	go Produce(c, done)
	go Consume(c)
	time.Sleep(time.Second * 2)
}

func Consume(c <-chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func Produce(c chan<- int, done chan bool) {
	for {
		c <- rand.Int()
	}
}
