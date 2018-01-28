package ftp

import (
	"net"
)

//ServerReceive server receive and save the contents
func ServerReceive(c net.Conn, filename string) error {
	return receive(c, filename)
}

//ClientSend client send contents to server
func ClientSend(c net.Conn, filename string) error {
	return send(c, filename)
}
