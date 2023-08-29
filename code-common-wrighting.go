package main

import "fmt"

// ===============一般函数====================
func sample_1() {
	echoHello_1()
}

func echoHello_1() {
	fmt.Println("hello,1")
}

// ================匿名函数====================
func sample_2() {
	go func() {
		fmt.Println("hello,world")
	}()
}

// ===============函数赋值给变量================
func sample_3() {
	echoHello := func() {
		fmt.Println("hello,3")
	}
	echoHello()
}

func main() {
	//sample_1()
	sample_3()
}
