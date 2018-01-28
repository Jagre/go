package ftp

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

//parsceCmd will parse command
//return: cmdline, cmd name, cmd parameters, error
func parseCmd(rd *bufio.Reader) (string, string, []string, error) {
	cmdline, e := rd.ReadString('\n')
	if e != nil {
		return "", "", nil, e
	}
	if len(cmdline) == 0 {
		return "", "", nil, errors.New("No command contents")
	}

	cmd := strings.Fields(cmdline)[0]
	cmd = strings.ToLower(cmd)
	params := []string{}
	for i := 1; i < len(strings.Fields(cmdline)); i++ {
		params = append(params, strings.Fields(cmdline)[i])
	}
	return cmdline, cmd, params, nil
}

func send(c net.Conn, filename string) error {
	//read contents from local by filename
	contents, e := Read(filename)
	if e != nil {
		return e
	}
	//write the contents to the distination machine
	_, e = c.Write(contents)
	if e != nil {
		return e
	}

	return nil
}

func receive(c net.Conn, filename string) error {
	contents := make([]byte, 512)
	//read content from distination machine
	_, e := c.Read(contents)
	if e != nil {
		return e
	}

	//save contents on local by filename
	e = Write(contents, filename)
	if e != nil {
		return e
	}
	return nil
}
