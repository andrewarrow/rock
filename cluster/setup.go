package cluster

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type Client struct {
	mu          sync.Mutex
	rip         string
	connections []*ClientConnection
}

type ClientConnection struct {
	hosts  map[string]net.Conn
	target string
	buffer []byte
}

func NewClient(poolSize int, ip, port string) *Client {
	c := Client{}
	c.connections = []*ClientConnection{}
	for i := 0; i < poolSize; i++ {
		cc := ClientConnection{}
		cc.hosts = map[string]net.Conn{}
		cc.buffer = make([]byte, MAX_SIZE)
		c.connections = append(c.connections, &cc)
	}
	c.rip = ip + ":" + port
	return &c
}

func (c *Client) ConnectAll() {
	for _, item := range c.connections {
		item.Connect(c.rip)
	}
}

func (c *ClientConnection) Connect(target string) bool {
	fmt.Println("dialing...", target)
	tcp, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return false
	}
	tcp.(*net.TCPConn).SetKeepAlive(true)
	tcp.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	c.target = target
	fmt.Println("target", target)
	c.hosts[target] = tcp
	return true
}

func (c *ClientConnection) RunCommand(command string) (string, bool, error) {
	var isThere net.Conn
	isThere = c.hosts[c.target]
	if isThere == nil {
		if c.Connect(c.target) == false {
			return "", false, fmt.Errorf("connect did not work")
		}
	}
	//fmt.Println("target2", c.target, command)
	_, err := c.hosts[c.target].Write([]byte(command))
	if err != nil {
		fmt.Println("rc1", err)
		c.hosts[c.target] = nil
		return "", false, err
	}

	return c.ReadAll()
	/*
		response, err := c.ReadAll()
		if err != nil {
			fmt.Println("rc1", err)
			c.hosts[c.target] = nil
			return "", false, err
		}
		//for i, item := range response {
		//	fmt.Println("target2b", command, i, item)
		//}

		first := response[0]

		reply := ""
		if strings.HasPrefix(first, "$-1") {
			reply = ""
		} else if strings.HasPrefix(first, "*-1") {
			reply = ""
		} else if strings.HasPrefix(first, "*0") {
			reply = ""
		} else if strings.HasPrefix(first, "$") {
			reply = strings.TrimSpace(strings.Join(response[1:], ""))
		} else if strings.HasPrefix(first, "*") {
		//		*3\r\n           # Array with 3 elements
			//	$5\r\nHello\r\n   # First element: Bulk string "Hello" with length 5
		//		:42\r\n           # Second element: Integer reply 42
			//	$11\r\nWorld!\r\n
			buffer := []string{}
			for _, item := range response[2:] {
				if !strings.HasPrefix(item, "$") {
					if strings.HasPrefix(item, ":") {
						buffer = append(buffer, item[1:])
					} else {
						buffer = append(buffer, item)
					}
				}
			}
			reply = strings.Join(buffer, ",")
		} else if strings.HasPrefix(first, ":") {
			reply = strings.TrimSpace(strings.Join(response[1:], ""))
		} else if strings.HasPrefix(first, "-MOVED") {
			c.handleMoved(first)
			return "", true, nil
		} else {
			reply = strings.TrimSpace(first)
		}
	*/
	//return reply, false, nil
}

func (c *Client) Close() {
	//TODO
}

func (c *ClientConnection) handleMoved(response string) {
	tokens := strings.Split(response, " ")
	moved := tokens[len(tokens)-1]
	//tokens = strings.Split(moved, ":")
	//c.target = c.rip + ":" + tokens[1]
	c.target = moved
}
