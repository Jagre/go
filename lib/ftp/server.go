package ftp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
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
		go parseConn(conn)
	}
}

func parseConn(conn net.Conn) {
	defer conn.Close()
	for {
		conn.Write([]byte("******Welcome to connect to my FTP server******"))
		reader := bufio.NewReader(conn)
		for {
			cmdline, e := reader.ReadString('\n')
			if e == io.EOF {
				//conn.Close()
				fmt.Println(e)
			}
			cmdline = strings.TrimSpace(cmdline)
			if len(cmdline) == 0 {
				continue //wait for input
			}
			cmd := strings.Fields(cmdline)[0]
			switch cmd {
			case "pull":
				download(conn, cmdline)
			case "push":
				upload(conn, cmdline)
			case "exit":
				return
			default:
				fmt.Println("Invaild command")
			}
		}
	}
}

//returns: commandname, params, e
func pasreCmd(c net.Conn) (string, []string, error) {
	// buf := []byte{}
	// _, e = c.Read(buf)
	// if e != nil {
	// 	return cmd, params, e
	// }

	reader := bufio.NewReader(c)
	cmdline, e := reader.ReadString('\n')
	if e != nil {
		return "", nil, e
	}
	if len(cmdline) == 0 {
		return "", nil, errors.New("No command contents")
	}

	cmd := strings.Fields(cmdline)[0]
	params := []string{}
	for i := 1; i < len(strings.Fields(cmdline)); i++ {
		params = append(params, strings.Fields(cmdline)[i])
	}
	return cmd, params, nil
}
