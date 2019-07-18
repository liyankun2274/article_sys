package utils

import (
	"github.com/boj/redistore"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"time"
)

// 创建 redis session
func GetRedisSession(redis_pool *redis.Pool, redis_prefix string) (echo.MiddlewareFunc, error) {
	store, err := redistore.NewRediStoreWithPool(redis_pool, []byte("serect"))
	if err != nil {
		return nil, err
	}
	store.SetKeyPrefix(redis_prefix)

	return session.Middleware(store), nil
}

// 创建redis pool
func NewRedisPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
