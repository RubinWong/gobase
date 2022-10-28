package main

import (
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	// 创建trace文件
	runtime.GOMAXPROCS(4)
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 200000; i++ {
		go func() {
			wg.Add(1)
			Cal()
			wg.Done()
		}()
	}

	wg.Wait()
}

func Cal() {
	sum := 1
	for i := 1; i < 100000; i++ {
		sum += i
	}
}
