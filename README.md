# UnixConn tools

Some stuff commonly used made more simple.

## socketpair

Calling

```go
	pair, err := unixconn.Pair()
```

Allows pair to contain two `*net.UnixConn` that are connected to each other.
Those can be used as replacement to TCP connections in test code, for example.

## cred

Golang already makes this easier thanks to `GetsockoptUcred` but I feel this
can be even better.

```go
	uid, gid, pid, err := unixconn.Cred(c)
```
