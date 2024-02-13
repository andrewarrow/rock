package cluster

import (
	"fmt"
	"io"
	"net"
)

func (c *Client) Info() string {
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)
	return cc.Info()
}

func (c *ClientConnection) Info() string {

	command := fmt.Sprintf("info\r\n")
	var isThere net.Conn
	isThere = c.hosts[c.target]
	if isThere == nil {
		if c.Connect(c.target) == false {
			return ""
		}
	}
	_, err := c.hosts[c.target].Write([]byte(command))
	c.handleErr("today"+command, err)
	reply := c.ReadInfoAll()
	return reply
}

func (c *ClientConnection) ReadInfoAll() string {
	buffer := make([]byte, 1024*900)
	data := []byte{}

	for {
		n, err := c.hosts[c.target].Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			fmt.Println(err)
			break
		}

		if n > 0 {
			//fmt.Println("n", n)
			data = append(data, buffer[0:n]...)
			break
		}
	}

	raw := string(data)
	return raw
}
