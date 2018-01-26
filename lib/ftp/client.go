package ftp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Client() {
	addr := "192.18.56.45:8080"
	//connect to server
	conn, e := net.Dial("tcp", addr)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 10240)
	n, e := conn.Read(buf)
	if e != nil && e != io.EOF {
		fmt.Println(e)
	}
	fmt.Println(string(buf[:n]))
	var cmd string
	reader := bufio.NewReader(os.Stdin)
	for {
		cmdline, e := reader.ReadString('\n')
		if e != nil {
			fmt.Println(e)
		}
		fmt.Sscan(cmdline, &cmd)
		if len(cmdline) == 1 {
			continue
		}
		go sending(&conn, cmdline)
	}
}

func sending(conn *net.Conn, cmdline string) {
	//send cmd to server
	_, e := (*conn).Write([]byte(cmdline))
	if e != nil {
		fmt.Println(e)
	}

	buf := make([]byte, 256)
	//read content of server returned
	_, e = (*conn).Read(buf)
	if e != nil {
		if e != io.EOF {
			fmt.Println(e)
		}
	}
	if len(buf) > 0 {
		fmt.Println(string(buf))
	}
}
