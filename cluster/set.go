package cluster

import "fmt"

const MAX_SIZE = 1024

func (c *Client) Set(k, v string) error {
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)

	command := fmt.Sprintf("SET %s %q\r\n", k, v)
	return cc.TryCommand(command)
}

func (c *Client) Get(k string) string {
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)
	command := fmt.Sprintf("GET %s\r\n", k)
	return cc.TryCommandWithReply(command)
}

func (c *Client) GetWithError(k string) (string, error) {
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)
	command := fmt.Sprintf("GET %s\r\n", k)
	return cc.TryCommandWithReplyAndError(command)
}
