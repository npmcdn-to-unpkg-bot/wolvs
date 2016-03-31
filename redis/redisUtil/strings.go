package redisUtil

import (
	"flag"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	redisAddress   = flag.String("redis-address", "localhost:6379", "Address to the Redis server")
	maxConnections = flag.Int("redis-max-connections", 10, "Max connections to Redis")
	redisPool      *redis.Pool
)

//Initialize will initialize the redis pool using the flags redis-address and redis-max-connections
func Initialize() {
	redisPool = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}

		return c, err
	}, *maxConnections)
}

//Close will close the redis connection pool made in InitializeRedis
func Close() {
	redisPool.Close()
}

//SetString will set a key and value in redis with no expiry
func SetString(key string, value string) (string, error) {
	c := redisPool.Get()
	defer c.Close()

	status, err := c.Do("SET", key, value)

	if err != nil {
		return fmt.Sprintf("Could not SET %s:%s", key, value), err
	}

	return fmt.Sprintf("%#v\n", status), err
}

//SetStringWithExpiry will set a key and a value with the provided expiry.
func SetStringWithExpiry(key string, value string, expiryInSeconds int) (string, error) {
	status, err := SetString(key, value)

	if expiryInSeconds > 0 {
		c := redisPool.Get()
		defer c.Close()

		c.Do("EXPIRE", key, expiryInSeconds)
	}

	return status, err
}

//GetString will get the contents of a redis key's value as a string
func GetString(key string) (string, error) {
	c := redisPool.Get()
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))

	if err != nil {
		message := fmt.Sprintf("Could not GET %s", key)
		return message, err
	}
	return value + "\n", err
}
