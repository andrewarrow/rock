# rock
Redis Orchestration of Cluster Konnections 

```
ip := os.Getenv("REDIS_CLUSTER_IP")
port := os.Getenv("REDIS_CLUSTER_PORT")
poolSize := 2
c := cluster.NewClient(poolSize, ip, port)
c.ConnectAll()
r := c.Info()
fmt.Println(r)
c.Set("test", "foo")
c.Get("test")
```

# why not goredis
I could not get goredis poolsize configs to work for me.
This is a very simple thread-safe redis cluster client.

# how it works

```
type Client struct {
	mu          sync.Mutex
	rip         string
	connections []*ClientConnection
}
```

A client has a pool of N *ClientConnections.

```
type ClientConnection struct {
	hosts  map[string]net.Conn
	target string
	buffer []byte
}
```

Each ClientConnection has a map of redis node hostname+port and the
current "target" it just connected to after the last MOV.

For example:

```
hosts["127.0.0.1:30001"] = conn1
hosts["127.0.0.1:30002"] = conn2
```

When it gets a MOV response with 127.0.0.1:30003 and 127.0.0.1:30003 is not
in the map of hosts, it will connect and add it:

hosts["127.0.0.1:30003"] = conn3

# calling commands

You can make calls like:

```
c.Set("test", "foo")
c.Get("test")
c.SAdd("foo", "bar")
c.SMembers("foo")
c.SRem("foo", "bar")
```

And they are thread safe because;

```
cc := c.TakeFromPool()
defer c.PlaceBackInPool(cc)
```

Before each command it gets it's own ClientConnection from the pool
and is the only one using that ClientConnection until it returns it.
