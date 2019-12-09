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

//新建redis连接
func New(options *Options) redis.Conn {
	if conn != nil {
		return conn
	}
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
	return connect
}

//连接redis,复用连接池
func Connect(options *Options) redis.Conn {
	if conn != nil {
		return conn
	}
	conn = New(options)
	return conn
}

//关闭redis
func Close() error {
	if conn != nil {
		return conn.Close()
	}
	return nil
}
