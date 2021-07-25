# UnixConn tools

Some stuff commonly used made more simple.

## socketpair

Calling

```go
	pair, err := unixconn.Pair()
```

Allows pair to contain two `net.Conn` that are connected to each other. Those
will actually be `*net.UnixConn` instances and I might just change the function
to match this.

## cred

Golang already makes this easier thanks to `GetsockoptUcred` but I feel this
can be even better.

```go
	uid, gid, pid, err := unixconn.Cred(c)
```
