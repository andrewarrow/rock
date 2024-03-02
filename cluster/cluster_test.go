package cluster

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var c *Client

func TestInfo(t *testing.T) {
	reply := c.Info()
	fmt.Println(reply)
}

func TestSet(t *testing.T) {
	c.Set("test", "foo")
	reply := c.Get("test")
	fmt.Println(reply)
	if reply != "foo" {
		t.Errorf("get returned %s, expected %s", reply, "foo")
	}
}

func TestSAdd(t *testing.T) {
	c.SAdd("set1", "val1")
	c.SAdd("set1", "val2")
	reply := c.SMembers("set1")
	fmt.Println("hi", reply)
	tokens := strings.Split(reply, ",")
	m := map[string]bool{}
	for _, token := range tokens {
		m[token] = true
	}
	if !m["val1"] || !m["val2"] {
		t.Errorf("get returned %v, expected %s", m, "val1,val2")
	}
	c.SRem("set1", "val1")
	reply = c.SMembers("set1")
	fmt.Println(reply)
	if reply != "val2" {
		t.Errorf("get returned %s, expected %s", reply, "val2")
	}
	cc := c.TakeFromPool()
	defer c.PlaceBackInPool(cc)
	if cc.hosts["127.0.0.1:30002"] == nil {
		t.Errorf("MOV should have connected to 30002")
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
