package eg_1

// 工厂方法模式
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

type RedisCache struct {
	data map[string]string
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		data: make(map[string]string),
	}
}

func (r *RedisCache) Set(key, value string) {
	r.data[key] = value
}

func (r *RedisCache) Get(key string) string {
	return r.data[key]
}

type MemCache struct {
	data map[string]string
}

func NewMemCache() *MemCache {
	return &MemCache{
		data: make(map[string]string),
	}
}

func (m *MemCache) Set(key, value string) {
	m.data[key] = value
}

func (m *MemCache) Get(key string) string {
	return m.data[key]
}

//type cacheType int
//
//const (
//	redis cacheType = iota
//	mem
//)

type CacheFactory interface {
	Create() Cache
}

type RedisCacheFactory struct {
}

func (rf RedisCacheFactory) Create() Cache {
	return NewRedisCache()
}

type MemCacheFactory struct {
}

func (mf MemCacheFactory) Create() Cache {
	return NewMemCache()
}
