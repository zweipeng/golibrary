package redis

import (
	xtime "github.com/zweipeng/golibrary/time"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Pool struct {
	redis.Pool
}

type Config struct {
	Addr           string
	Password       string
	ClientName	   string
	MaxIdle        int
	MaxActive      int
	IdleTimeout    xtime.Duration
	DialTimeout    xtime.Duration
	ReadTimeout    xtime.Duration
	WriteTimeout   xtime.Duration
}

//creat a new pool
func NewPool(c *Config) (p * Pool)  {
	if c.DialTimeout <=0 || c.WriteTimeout <=0 || c.ReadTimeout <=0 {
		panic("must config redis timeout")
	}

	dialFunc := func() (redis.Conn, error) {
		return redis.Dial(
			"tcp",
			c.Addr,
			redis.DialPassword(c.Password),
			redis.DialClientName(c.ClientName),
			redis.DialConnectTimeout(time.Duration(c.DialTimeout)),
			redis.DialReadTimeout(time.Duration(c.ReadTimeout)),
			redis.DialWriteTimeout(time.Duration(c.WriteTimeout)))
	}

	return &Pool{redis.Pool{
		Dial:            dialFunc,
		MaxIdle:         c.MaxIdle,
		MaxActive:       c.MaxActive,
		IdleTimeout:     time.Duration(c.IdleTimeout),
	}}
}

type Error string

func (err Error) Error() string { return string(err) }
