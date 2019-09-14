package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Client struct {
	conn redis.Conn
}

func NewClient(address string) *Client {
	conn, err := redis.Dial("tcp", address,
		redis.DialReadTimeout(30*time.Second),
		redis.DialWriteTimeout(30*time.Second),
		redis.DialConnectTimeout(30*time.Second),
		redis.DialKeepAlive(5*time.Minute))
	if err != nil {
		panic(err)
	}
	return &Client{
		conn: conn,
	}
}

func (c *Client) Auth(password string) (reply interface{}, err error) {
	return c.conn.Do("AUTH", password)
}

func (c *Client) Select(dbID int) (reply interface{}, err error) {
	return c.conn.Do("SELECT", dbID)
}

func (c *Client) Close() {
	if err := c.conn.Close(); err != nil {
		panic(err)
	}
}
