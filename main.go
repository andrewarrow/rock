package main

import (
	"fmt"
	"math/rand"
	"os"
	"rock/cluster"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "" {
	} else if command == "info" {
		ip := os.Getenv("REDIS_CLUSTER_IP")
		port := os.Getenv("REDIS_CLUSTER_PORT")
		c := cluster.NewClient(2, ip, port)
		c.ConnectAll()
		r := c.Info()
		fmt.Println(r)
	}
}
