package cluster

import (
	"fmt"
	"strconv"
	"strings"
)

func fixForIntReply(s string) string {
	numString := strings.TrimSpace(s)
	return numString[1:]
}

func (c *ClientConnection) ReadFirst() (string, error) {
	n, err := c.hosts[c.target].Read(c.buffer)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	first := string(c.buffer[0:n])
	//fmt.Println(first)
	return first, nil
}

func (c *ClientConnection) ReadToLimit(first string) (string, error) {
	sizeBuffer := []byte{}
	asBytes := []byte(first[1:])
	for _, char := range asBytes {
		if char == 13 {
			break
		}
		sizeBuffer = append(sizeBuffer, char)
	}
	sizeString := string(sizeBuffer)
	sizeLimit, _ := strconv.Atoi(sizeString)

	data := []byte(strings.TrimSpace(first[len(sizeString)+1:]))
	//fmt.Println(string(data))

	if len(data) == sizeLimit {
		return string(data), nil
	}

	for {
		n, err := c.hosts[c.target].Read(c.buffer)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		thing := c.buffer[0:n]
		data = append(data, thing...)

		//fmt.Println("b", len(data), sizeLimit)
		if len(data)-2 == sizeLimit {
			break
		}
	}

	return string(data), nil
}

func (c *ClientConnection) ReadMembers(first string) (string, error) {
	lines := strings.Split(first[1:len(first)-2], "\r\n")
	if len(lines) == 0 {
		return "", fmt.Errorf("size")
	}
	total, _ := strconv.Atoi(lines[0])
	complete := len(lines)-1 == total*2
	for i, line := range lines {
		fmt.Println(i, line)
	}
	fmt.Println(complete, total, len(lines))
	/*
		*3\r\n           # Array with 3 elements
		$5\r\nHello\r\n   # First element: Bulk string "Hello" with length 5
		:42\r\n           # Second element: Integer reply 42
		$11\r\nWorld!\r\n
	*/
	/*
		for _, item := range response[2:] {
			if !strings.HasPrefix(item, "$") {
				if strings.HasPrefix(item, ":") {
					buffer = append(buffer, item[1:])
				} else {
					buffer = append(buffer, item)
				}
			}
		}*/
	return string(""), nil
}
