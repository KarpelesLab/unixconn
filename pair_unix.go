package unixconn

import (
	"net"
	"os"
	"syscall"
)

// Pair will return a pair of interconnected sockets that can be used for
// testing and more.
func Pair() ([]net.Conn, error) {
	pair, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}
	// make into files
	f1 := os.NewFile(uintptr(pair[0]), "")
	f2 := os.NewFile(uintptr(pair[1]), "")
	defer f1.Close()
	defer f2.Close()
	link1, err := net.FileConn(f1)
	if err != nil {
		return nil, err
	}
	link2, err := net.FileConn(f2)
	if err != nil {
		return nil, err
	}
	return []net.Conn{link1, link2}, nil
}
