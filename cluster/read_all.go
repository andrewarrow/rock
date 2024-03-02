package cluster

import (
	"fmt"
	"strings"
)

func (c *ClientConnection) ReadAll() (string, bool, error) {
	//data := []byte{}
	//sizeLimit := 0
	reply := ""

	for {
		n, err := c.hosts[c.target].Read(c.buffer)
		if err != nil {
			fmt.Println(err)
			return "", false, err
		}

		thing := c.buffer[0:n]
		first := string(thing)
		fmt.Println(first)

		if strings.HasPrefix(first, "$-1") {
		} else if strings.HasPrefix(first, "*-1") {
		} else if strings.HasPrefix(first, "*0") {
		} else if strings.HasPrefix(first, "$") {
		} else if strings.HasPrefix(first, "*") {
		} else if strings.HasPrefix(first, ":") {
		} else if strings.HasPrefix(first, "-MOVED") {
			c.handleMoved(first)
			return "", true, nil
		} else {
		}
	}

	//raw := string(data)
	//tokens := strings.Split(raw, "\r\n")
	//return tokens[0 : len(tokens)-1], nil
	return reply, false, nil
}
