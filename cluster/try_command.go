package cluster

import "fmt"

func (c *ClientConnection) TryCommand(command string) error {
	_, moved, err := c.RunCommand(command)
	if err != nil {
		return err
	}
	if moved {
		_, _, err = c.RunCommand(command)
		return err
	}
	return nil
}

func (c *ClientConnection) TryCommandWithReply(command string) string {
	reply, moved, err := c.RunCommand(command)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if moved {
		reply, _, err = c.RunCommand(command)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return reply
	}
	return reply
}

func (c *ClientConnection) TryCommandWithReplyAndError(command string) (string, error) {
	reply, moved, err := c.RunCommand(command)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if moved {
		reply, _, err = c.RunCommand(command)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return reply, nil
	}
	return reply, nil
}
