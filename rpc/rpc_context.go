package rpc

import "github.com/dyike/nfs_server/vfs"

type RPCContext struct {
	LocalPort   uint16
	ClientAddr  string
	Auth        AuthUnix
	VFS         vfs.NFSFileSystem
	MountSignal chan bool
}
