package main

import (
	"github.com/go-redis/redis"
	"time"
)

const (
	URLIDKEY           = "next.url.id"
	ShortlinkKey       = "shortlink:%s:url"
	URLHashKey         = "uslhash:%s:url"
	shortlinkDetailKey = "shortlink:%s:detail"
)

type RedisCli struct {
	Cli *redis.Client
}

type UrlDetail struct {
	Url                string        `json:"url"`
	CreatedAt          string        `json:"created_at"`
	ExpirationInMiutes time.Duration `json:"expiration_in_miutes"`
}

func NewRidisCli(addr string, pwd string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}
}
