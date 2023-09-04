package eg_1

import (
	"fmt"
	"testing"
)

func TestRedisCache(t *testing.T) {
	var redisCacheFactory CacheFactory
	redisCacheFactory = RedisCacheFactory{}
	redisCache := redisCacheFactory.Create()
	redisCache.Set("key1", "value1")
	value := redisCache.Get("key1")
	fmt.Println(">> redis output is:", value)
}

func TestMemCache(t *testing.T) {
	var memCacheFactory CacheFactory
	memCacheFactory = MemCacheFactory{}
	memCache := memCacheFactory.Create()
	memCache.Set("key1", "value1")
	value := memCache.Get("key1")
	fmt.Println(">> mem output is:", value)
}
