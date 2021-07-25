package unixconn

// Cred returns details on the connected unix connection, useful for
// identifying the connecting process.
func Cred(c *net.UnixConn) (uid, gid, pid int, err error) {
	uid, gid, pid = -1, -1, -1

	var rc syscall.RawConn
	rc, err = c.SyscallConn()
	if err != nil {
		return
	}

	err = rc.Control(func(fd uintptr) {
		ucred, _ := syscall.GetsockoptUcred(int(fd), syscall.SOL_SOCKET, syscall.SO_PEERCRED)
		uid = int(ucred.Uid)
		gid = int(ucred.Gid)
		pid = int(ucred.Pid)
	})
	return
}
