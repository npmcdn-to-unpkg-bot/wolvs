package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	redisAddress   = flag.String("redis-address", "localhost:6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
	redisPool      *redis.Pool
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Gorilla!\n"))
}

func InsertHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	value := vars["value"]
	c := redisPool.Get()
	defer c.Close()

	status, err := c.Do("SET", key, value)

	if err != nil {
		message := fmt.Sprintf("Could not SET %s:%s", key, value)

		http.Error(writer, message, 400)
	} else {
		writer.WriteHeader(200)
		writer.Write([]byte(fmt.Sprintf("%#v\n", status)))
	}
}

func GetHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]

	c := redisPool.Get()
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))

	if err != nil {
		message := fmt.Sprintf("Could not GET %s", key)

		http.Error(writer, message, 400)
	} else {
		writer.Write([]byte(value + "\n"))
	}
}

func InitializeRedis() {
	redisPool = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}

		return c, err
	}, *maxConnections)
}

func main() {
	flag.Parse()
	InitializeRedis()
	defer redisPool.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/add/{key}/{value}", InsertHandler).Methods("GET")
	router.HandleFunc("/add/{key}", GetHandler).Methods("GET")

	n := negroni.Classic()

	n.UseHandler(router)
	n.Run(":8080")
}
