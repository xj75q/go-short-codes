package main

import (
	"fmt"
	"sync"
)

func sync1() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello,world")
		mu.Unlock()
	}()
	mu.Lock()
}

func sync2() {
	done := make(chan int)
	go func() {
		fmt.Println("hello,world,2")
		<-done
	}()
	done <- 1
}

func sync3() {
	done := make(chan int, 1)
	go func() {
		fmt.Println("hello,world,3")
		done <- 1
	}()
	<-done
}

func sync4() {
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("hello,world4")
			done <- 1
		}()
	}

	for i := 0; i < cap(done); i++ {
		fmt.Println("=========")
		<-done
	}
}

func sync5() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("hello,world,5")
			wg.Done()
		}()
	}
	//等待第N个后台完成
	wg.Wait()
}

func main() {
	//sync1()
	//sync2()
	//sync3()
	//sync4()
	sync5()
}
