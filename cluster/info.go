package cluster

import (
	"fmt"
	"net"
	"strconv"
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
	buffer := make([]byte, 1024)
	data := []byte{}

	sizeLimit := 0
	for {
		n, err := c.hosts[c.target].Read(buffer)
		if err != nil {
			fmt.Println(err)
			return ""
		}

		thing := buffer[0:n]
		data = append(data, thing...)

		if len(thing) < 4 {
			return ""
		}
		if thing[0] == '$' && sizeLimit == 0 {
			sizeBuffer := []byte{}
			for _, char := range thing[1:] {
				if char == 13 {
					break
				}
				sizeBuffer = append(sizeBuffer, char)
			}
			sizeString := string(sizeBuffer)
			sizeLimit, _ = strconv.Atoi(sizeString)
			data = data[len(sizeString)+1+2:]
		}

		if len(data)-2 == sizeLimit {
			break
		}
	}

	raw := string(data)
	return raw
}
