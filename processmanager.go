package libclient2

import "net"

type ProcessManagerClient struct {
	sock *net.UnixConn
}

func (client *ProcessManagerClient) Run() {
}

// GENERIC METHODS

func (client *ProcessManagerClient) Send(command string, params map[string]interface{}) bool {
	return genericSend(client, command, params)
}

func (client *ProcessManagerClient) write(b []byte) (int, error) {
	return genericWrite(client, b)
}

func (client *ProcessManagerClient) socket() *net.UnixConn {
	return client.sock
}
