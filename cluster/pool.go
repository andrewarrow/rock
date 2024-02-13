package cluster

import "net"

func (c *Client) TakeFromPool() *ClientConnection {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.connections) == 0 {
		cc := ClientConnection{}
		cc.hosts = map[string]net.Conn{}
		cc.buffer = make([]byte, MAX_SIZE*2)
		cc.Connect(c.rip)
		return &cc
	}

	conn := c.connections[0]
	c.connections = c.connections[1:]

	return conn
}

func (c *Client) PlaceBackInPool(cc *ClientConnection) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connections = append(c.connections, cc)
}
