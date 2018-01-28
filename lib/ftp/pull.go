package ftp

import (
	"net"
)

//ServerSend server send file contents to client
//param: c never close
func ServerSend(c net.Conn, filename string) error {
	return send(c, filename)
}

//ClientReceive client will receive and save the file that was transfered from server
func ClientReceive(c net.Conn, filename string) error {
	return receive(c, filename)
}
