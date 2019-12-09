package goredis

import (
	"github.com/gomodule/redigo/redis"
)

var (
	conn redis.Conn
)

type Options struct {
	Protocol string
	Addr     string
	Password string
	Database int32
}

//连接redis
func Connect(options *Options) redis.Conn {
	if options == nil {
		return nil
	}
	dialOption := redis.DialPassword(options.Password)
	var connect redis.Conn
	if c, err := redis.Dial(options.Protocol, options.Addr, dialOption); err != nil {
		return nil
	} else {
		connect = c
	}
	if _, err := connect.Do("SELECT", options.Database); err != nil {
		return nil
	}
	conn = connect
	return conn
}

//关闭redis
func Close() error {
	if conn != nil {
		return conn.Close()
	}
	return nil
}
