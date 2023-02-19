package xredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type ZpRedis struct {
	conn redis.Conn
}

func (zpr *ZpRedis) Close() error {
	fmt.Println("zpredis close")
	return zpr.conn.Close()
}

func (zpr *ZpRedis) Err() error {
	return zpr.conn.Err()
}

func (zpr *ZpRedis) Do(commandName string, args ...interface{}) (interface{}, error) {
	start := time.Now()
	reply, err := zpr.conn.Do(commandName, args...)
	end := time.Now()
	diff := end.Sub(start)
	timeout := time.Second

	if diff > timeout {
		fmt.Printf("ZpRedis slow log. diff(%s) start(%d) end(%d) err(%s) cmd(%s) args(%+v)", diff, start.Unix(), end.Unix(), err, commandName, args)
	}

	return reply, err
}

func (zpr *ZpRedis) Send(commandName string, args ...interface{}) error {
	return zpr.conn.Send(commandName, args...)
}

func (zpr *ZpRedis) Flush() error {
	return zpr.conn.Flush()
}

func (zpr *ZpRedis) Receive() (interface{}, error) {
	return zpr.conn.Receive()
}
func zpRedisDial(network, address string, options ...redis.DialOption) (redis.Conn, error) {
	conn, err := redis.Dial(network, address, options...)
	if err != nil {
		return nil, err
	}
	return &ZpRedis{conn: conn}, nil
}

func makePool(addr string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: time.Minute * 10,
		MaxActive:   20,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return zpRedisDial("tcp", addr, redis.DialReadTimeout(time.Millisecond*200))
			//return redis.Dial("tcp", addr, redis.DialReadTimeout(time.Second/50))
		},
	}

}

var defaultPool *redis.Pool

func init() {
	fmt.Println("init redis pool")
	defaultPool = makePool("127.0.0.1:6379")
}

func GetRedis() redis.Conn {
	fmt.Println("zpredis GetRedis")
	return defaultPool.Get()
}
