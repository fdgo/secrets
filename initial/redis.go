package initial

import (
	"github.com/go-redis/redis"
	"sync"
)

var clientInstance *redis.Client
var redisLock sync.Mutex

func InitRedis() *redis.Client {
	clientInstance = redis.NewClient(&redis.Options{
		Addr:     "120.27.239.127:6379",
		Password: "",
		DB:       0,
	})
	return clientInstance
	//defer client.Close()
}

func RedisClient() *redis.Client {
	if clientInstance != nil {
		return clientInstance
	}
	redisLock.Lock()
	defer redisLock.Unlock()
	if clientInstance != nil {
		return clientInstance
	}
	return InitRedis()
}
