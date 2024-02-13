package cluster

import "fmt"

func (c *Client) SAdd(k, v string) error {
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)
	command := fmt.Sprintf("SADD %s %s\r\n", k, escape(v))
	return cc.TryCommand(command)
}

func (c *Client) SMembers(k string) string {
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)
	command := fmt.Sprintf("SMEMBERS %s\r\n", k)
	return cc.TryCommandWithReply(command)
}

func (c *Client) SRem(k, v string) error {
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)
	command := fmt.Sprintf("SREM %s %s\r\n", k, escape(v))
	return cc.TryCommand(command)
}
