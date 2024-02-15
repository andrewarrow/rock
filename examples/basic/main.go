package main

import (
	"fmt"

	"codeberg.org/andrewarrow/rock/cluster"
)

func main() {
	ip := "127.0.0.1"
	port := "30001"
	poolSize := 2
	c := cluster.NewClient(poolSize, ip, port)
	c.ConnectAll()
	c.Set("foo", "bar")
	r := c.Get("foo")
	fmt.Println(r)
}
