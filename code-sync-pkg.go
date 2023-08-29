package main

import (
	"fmt"
	"sync"
	"time"
)

func syncMutex() {
	var sc sync.Mutex
	sc.Lock()
	go func() {
		fmt.Println("hello sync...")
		sc.Unlock()
	}()
	sc.Lock()
}

func syncWaitgroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("hello,waitgroup...")
		}()
	}
	wg.Wait()
	fmt.Println("end...")
}

// lock()and unlock()
func syncRWMutexNormal() {
	var mu sync.RWMutex
	mu.Lock()

	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			mu.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mu.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	mu.Unlock()
	time.Sleep(time.Second)
	for _, c := range channels {
		fmt.Printf("这是值 %v \n", <-c)
	}
}

func syncRWMutexR() {
	var mu sync.RWMutex
	mu.Lock()

	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("*> Not read lock: ", i)
			mu.RLock() //Rlock
			fmt.Println("#> Read Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("^> Unlock the read lock: ", i)
			mu.RUnlock() //Runlock
			c <- i
		}(i, channels[i])
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock....")
	mu.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		fmt.Printf("值：%v\n", <-c)
	}
}

// =================go的消息队列====================
func syncQueue() {
	c := sync.NewCond(&sync.Mutex{})

	queue := make([]interface{}, 0, 10)
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println(">>>>Removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("****Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

// ===========go的单例模式====================
func syncOnce() {
	var (
		count int
		once  sync.Once
	)

	increment := func() {
		count++
		fmt.Printf("第%v次执行\n", count)
	}

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		fmt.Printf("循环到：%v\n", i)
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf(">>> count is %d \n", count)
}

// ==============sync.pool线程池============
var pool *sync.Pool

type Person1 struct {
	Name string
}

func NewHandlerPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println(">> Creating a new Person")
			return new(Person1)
		},
	}
}

func syncPool() {
	NewHandlerPool()
	p := pool.Get().(*Person1)
	fmt.Printf("## 首次从 pool 里获得：%v\n", p)

	p.Name = "first"
	fmt.Printf("** 设置 p.Name = %s\n", p.Name)

	pool.Put(p)
	fmt.Printf("&& Pool 里已有一个对象，调用 Get获得: %v\n", pool.Get().(*Person1))

	fmt.Println("@@ 此时对象里未知，调用 Get获得: ", pool.Get().(*Person1))
}

func main() {
	//syncMutex()
	//syncWaitgroup()
	//syncRWMutexNormal()
	//syncRWMutexR()
	//syncQueue()
	//syncOnce()
	syncPool()
}
