package main

import (
	"fmt"
	"github.com/qingwenjie/go-redis"
)

func main() {
	options := &redis.Options{
		Protocol: "tcp",
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		Database: 1,
	}
	conn := redis.Connect(options)
	reply, err := conn.Do("SELECT", 1)
	fmt.Println(reply, err)
}
