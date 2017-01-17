package client

import (
	"net"
	// "net/http"
)

//Client 指定一个客户端机器
type Client struct {
	id      int //机器唯一标识
	uid     int //uid
	ctype   int //
	conn    net.Conn
	address string
}

func (c *Client) Write(data []byte) {
	conn := c.conn
	conn.Write(data)
}

type loopFunc func([]byte)

func (c *Client) Loop(callback loopFunc) error {
	conn := c.conn
	var buf [512]byte
	for {
		_, err := conn.Read(buf[0:])
		if err != nil {
			return err
		}

		callback(buf[0:])
	}
}
