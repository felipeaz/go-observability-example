package redis

import (
	"fmt"
	"log"

	redigo "github.com/gomodule/redigo/redis"
)

type Redis struct {
	addr string
	port string
}

func New(host, port string) *Redis {
	cache := &Redis{
		addr: host,
		port: port,
	}

	return cache
}

func (r *Redis) connect() (redigo.Conn, error) {
	conn, err := redigo.Dial("tcp", r.getHost())
	if err != nil {
		log.Fatalf(_failedToConnectToRedisServer, err)
	}

	return conn, err
}

func (r *Redis) Set(key string, value interface{}) error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do(_setAction, key, value)
	if err != nil {
		log.Printf(_failedToSetKey, key, value, err)
	}

	return err
}

func (r *Redis) Get(key string) ([]byte, error) {
	conn, err := r.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	data, err := redigo.Bytes(conn.Do(_getAction, key))
	if err != nil {
		log.Printf(_failedToGetKey, key, err)
	}
	return data, err
}

func (r *Redis) Remove(key string) error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do(_deleteAction, key)
	if err != nil {
		log.Printf(_failedToRemoveKey, key, err)
	}

	return err
}

func (r *Redis) getHost() string {
	return fmt.Sprintf("%s:%s", r.addr, r.port)
}
