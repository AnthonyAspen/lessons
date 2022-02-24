package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Employee struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

var pool = newPool()

func main() {
	client := pool.Get()
	defer client.Close()

	db_data, err := json.Marshal(Employee{"Alex", "San Franco"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("SET", "id123", db_data, 0)
	if err != nil {
		log.Fatal(err)
	}

	value, err := client.Do("GET", "id123")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s \n", value)

}
