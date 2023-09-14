package main

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	S *RedisCli
}

func getEnv() *Env {
	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	pwd := os.Getenv("APP_REDIS_PASSWD")
	if pwd == "" {
		pwd = ""
	}

	dbS := os.Getenv("APP_REDIS_DB")
	if dbS == "" {
		dbS = "0"
	}

	db, err := strconv.Atoi(dbS)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connect to redis (addr:%s ,pwd:%s ,db:%v)", addr, pwd, db)
	r := NewRidisCli(addr, pwd, db)
	return &Env{S: r}
}
