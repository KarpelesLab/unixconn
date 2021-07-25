package unixconn

import (
	"errors"
	"net"
	"syscall"
)

// if you can think of a better name for this interface I'm open to suggestions
type syscallGettable interface {
	SyscallConn() (syscall.RawConn, error)
}

// Cred returns details on the connected unix connection, useful for
// identifying the connecting process.
func Cred(c net.Conn) (uid, gid, pid int, err error) {
	uid, gid, pid = -1, -1, -1

	var rc syscall.RawConn
	if ic, ok := c.(syscallGettable); !ok {
		err = errors.New("connection does not provides syscall interface")
		return
	} else {
		rc, err = ic.SyscallConn()
		if err != nil {
			return
		}
	}

	err = rc.Control(func(fd uintptr) {
		ucred, _ := syscall.GetsockoptUcred(int(fd), syscall.SOL_SOCKET, syscall.SO_PEERCRED)
		uid = int(ucred.Uid)
		gid = int(ucred.Gid)
		pid = int(ucred.Pid)
	})
	return
}
