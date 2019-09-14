package main

import (
	"testing"
)

func initClient(t *testing.T) *Client {
	c := NewClient("localhost:6379")
	reply, err := c.Auth("123456")
	if err != nil {
		t.Error(err)
	} else {
		str := reply.(string)
		if str != "OK" {
			t.Fail()
		}
	}
	return c
}

func TestClient_Select(t *testing.T) {
	c := initClient(t)
	defer c.Close()

	reply, err := c.Select(1)
	if err != nil {
		t.Error(err)
	} else {
		str := reply.(string)
		if str != "OK" {
			t.Fail()
		}
	}
}
