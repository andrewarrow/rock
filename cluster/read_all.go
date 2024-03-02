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

	data := []byte(first[len(sizeString)+1:])

	fmt.Println(len(data), sizeLimit)
	if len(data)-2 == sizeLimit {
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

		fmt.Println("b", len(data), sizeLimit)
		if len(data)-2 == sizeLimit {
			break
		}
	}

	return string(data), nil
}
