package cluster

import (
	"fmt"
	"io"
)

func (c *ClientConnection) handleErr(s string, err error) error {
	if err != nil && err != io.EOF {
		fmt.Println("Error:", s, err)
		return nil
	}
	return nil
}
