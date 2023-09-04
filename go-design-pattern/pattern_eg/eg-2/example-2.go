package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//装饰器模式-模拟http中间件
/*
实现一个http server
实现一个 handler:hello
实现中间件的功能
1> 记录请求的url和方法
2> 记录请求的网络的地址
3> 记录方法的执行时间

*/

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("记录请求的网络地址：%s", r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Println("logging end...")
	})
}

func timeRecord(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		next.ServeHTTP(w, r)
		endTime := time.Since(startTime)
		log.Printf("记录方法的执行时间 %s", endTime)
		log.Println("timeRecod end...")

	})
}

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("记录请求的url和方法：%s,%s", r.URL, r.Method)
		next.ServeHTTP(w, r)
		log.Println("tracing end...")
	})
}

func hello2(w http.ResponseWriter, r *http.Request) {
	//log.Printf("记录请求的url和方法：%s，%s", r.URL, r.Method)
	//log.Printf("记录请求的网络地址%s", r.RemoteAddr)
	//
	//startTime := time.Now()
	fmt.Fprintf(w, "hello")
	//endTime := time.Since(startTime)
	//log.Printf("记录方法的执行时间:%s", endTime)
}

func main() {
	http.Handle("/", tracing(logging(timeRecord(http.HandlerFunc(hello2)))))
	http.ListenAndServe(":8080", nil)
}
