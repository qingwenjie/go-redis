package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/qingwenjie/goredis"
)

func main() {
	options := &goredis.Options{
		Protocol: "tcp",
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		Database: 1,
	}
	conn := goredis.Connect(options)
	reply, err := conn.Do("SELECT", 1)
	fmt.Println(reply, err)

	reply1, err := redis.Int(conn.Do("HSET", "myhash", "field1", "Hello"))
	fmt.Println(reply1, err)

	reply2, err := redis.String(conn.Do("HGET", "myhash", "field1"))
	fmt.Println(reply2, err)
}
