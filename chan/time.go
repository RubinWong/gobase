package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	Timeout(func() { fmt.Println("time.After 1 second") })
	fmt.Println(time.Now().Unix())
	// ConsumeAndProduce()
}

func ConsumeAndProduce() {
	var c chan int = make(chan int, 10)
	done := make(chan bool)
	go Produce(c, done)
	go Consume(c)
	time.Sleep(time.Second * 2)
}

func Consume(c <-chan int) {
	// for {
	// 	if i, ok := <-c; ok {
	// 		fmt.Println(i)
	// 	} else {
	// 		return
	// 	}
	// }
	for i := range c {
		fmt.Println(i, time.Now().UnixNano())
	}
}

func Produce(c chan<- int, done chan bool) {
	tk := time.NewTicker(time.Millisecond * 500)
	to := time.NewTimer(time.Second)
	for {
		select {
		case <-tk.C:
			c <- rand.Int()
		case <-to.C:
			c <- 100
			// default:
			// 	fmt.Println("default")
		}
	}
}

func Timeout(f func()) {
	to := time.NewTimer(time.Second)
	<-to.C

	f()
}
