package libclient2

import "net"

type Client interface {

	// send a JSON command.
	Send(command string, params map[string]interface{}) bool

	// write a slice of bytes to the socket.
	write(b []byte) (int, error)

	// fetch the socket.
	socket() *net.UnixConn
}
