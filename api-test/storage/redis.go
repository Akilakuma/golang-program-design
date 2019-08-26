package storage

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

// PushRedis 對redis 做rpsuh
func PushRedis() error {

	conn := getConnectionPool()
	redisC := conn.Get()
	defer redisC.Close()

	_, err := redisC.Do("RPUSH", "data_queue", "num")
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

// getConnectionPool 實質建立連線
func getConnectionPool() *redis.Pool {

	return &redis.Pool{
		Wait:        true,
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Println(err.Error())
			}
			return err
		},
	}
}
