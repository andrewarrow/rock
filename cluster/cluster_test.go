package cluster

import (
	"fmt"
	"os"
	"testing"
)

var c *Client

func TestSet(t *testing.T) {
	c.Set("test", "foo")
	reply := c.Get("test")
	fmt.Println(reply)
	if reply != "foo" {
		t.Errorf("get returned %s, expected %s", reply, "foo")
	}
}

func setup() {
	ip := "127.0.0.1"
	port := "30001"
	poolSize := 2
	c = NewClient(poolSize, ip, port)
	c.ConnectAll()
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}
