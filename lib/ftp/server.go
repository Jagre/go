package ftp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Server() {
	addr := ":8080"
	listener, e := net.Listen("tcp", addr)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("Error: " + e.Error())
			//os.Exit(1)
		}
		fmt.Printf("Client: %v | CONNECTED\n", conn.RemoteAddr().String())
		go serverHandler(conn)
	}
}

func serverHandler(c net.Conn) {
	defer c.Close()
	c.Write([]byte(fmt.Sprintf("Welcome to connect to FTP server | %v", c.LocalAddr().String())))
	rd := bufio.NewReader(c)
	for {
		_, cmd, params, e := parseCmd(rd)
		fmt.Println(cmd)
		if e != nil {
			if e != io.EOF {
				fmt.Printf("%v\nClinet: %v | DISCONNECTING\n", e, c.RemoteAddr().String())
			}
			return
		}

		fmt.Println(cmd)
		switch cmd {
		case "pull":
			if len(params) == 0 {
				msg := "Haven't speicified file name"
				fmt.Println(msg)
				c.Write([]byte(msg))
				continue
			}
			e = ServerSend(c, params[0])
			if e != nil {
				fmt.Println(e)
				c.Write([]byte(e.Error()))
			}
		case "push":
			if len(params) < 2 {
				msg := "Haven't speicified the distination filename that save on server"
				fmt.Println(msg)
				c.Write([]byte(msg))
				continue
			}
			e = ServerReceive(c, params[1])
			if e != nil {
				fmt.Println(e)
				c.Write([]byte(e.Error()))
			}
		case "exit":
			return
		default:
			fmt.Println("Invaild command")
		}

	}
}
