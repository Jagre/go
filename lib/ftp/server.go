package ftp

import (
	"bufio"
	"fmt"
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
	for {
		c.Write([]byte(fmt.Sprintf("Welcome to connect to FTP server | %v", c.LocalAddr().String())))
		rd := bufio.NewReader(c)
		_, cmd, params, e := parseCmd(rd)
		if e != nil {
			fmt.Println(e)
			continue
		}
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

//returns: commandname, params, e
// func parseCmdServer(c net.Conn) (string, []string, error) {
// 	// buf := []byte{}
// 	// _, e = c.Read(buf)
// 	// if e != nil {
// 	// 	return cmd, params, e
// 	// }

// 	reader := bufio.NewReader(c)
// 	cmdline, e := reader.ReadString('\n')
// 	if e != nil {
// 		return "", nil, e
// 	}
// 	if len(cmdline) == 0 {
// 		return "", nil, errors.New("No command contents")
// 	}

// 	cmd := strings.Fields(cmdline)[0]
// 	cmd = strings.ToLower(cmd)
// 	params := []string{}
// 	for i := 1; i < len(strings.Fields(cmdline)); i++ {
// 		params = append(params, strings.Fields(cmdline)[i])
// 	}
// 	return cmd, params, nil
// }
