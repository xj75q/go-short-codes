package main

import (
	"log"
	"net/http"
	"time"
)

type MiddleWare struct {
}

func (m *MiddleWare) LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := t1.Sub(time.Now())
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2)
	}
	return http.HandlerFunc(fn)
}

func (m *MiddleWare) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic:%v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
