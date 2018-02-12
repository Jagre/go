package ftp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Client() {
	//addr := "192.168.1.105:8080"
	addr := "172.18.23.45:8080"
	//connect to server
	conn, e := net.Dial("tcp", addr)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 256)
	n, e := conn.Read(buf)
	if e != nil && e != io.EOF {
		fmt.Println(e)
	}
	fmt.Println(string(buf[:n]))
	rd := bufio.NewReader(os.Stdin)
	for {
		cmdline, cmd, params, e := parseCmd(rd)
		if e != nil {
			fmt.Println(e)
			continue
		}
		// fmt.Sscan(cmdline, &cmd)
		// if len(cmdline) == 1 {
		// 	continue
		// }
		if len(cmdline) > 0 {
			go clientHandler(conn, cmdline, cmd, params)
		}
	}
}

func clientHandler(c net.Conn, cmdline, cmd string, params []string) {
	sendCmd(c, cmdline)
	switch cmd {
	case "pull":
		if len(params) < 2 {
			fmt.Println("Haven't specified the file name that you wanna save")
			return
		}
		e := ClientReceive(c, params[1])
		if e != nil {
			fmt.Println(e)
		}
	case "push":
		if len(params) < 1 {
			fmt.Println("Haven't specified the file name that you wanna upload")
			return
		}

		e := ClientSend(c, params[0])
		if e != nil {
			fmt.Println(e)
		}
	default:
		receiveMsg(c)
	}

}

func sendCmd(c net.Conn, cmdline string) {
	//send cmd to server
	_, e := c.Write([]byte(cmdline))
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("shit")
}

func receiveMsg(c net.Conn) {
	msg := make([]byte, 256)
	//read message of server returned
	_, e := c.Read(msg)
	if e != nil {
		if e != io.EOF {
			fmt.Println(e)
		}
	}
	if len(msg) > 0 {
		fmt.Println(string(msg))
	}
}
