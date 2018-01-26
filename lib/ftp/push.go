package ftp

import (
	"net"
)

//Push client upload file to the server (run at client)
func Push(c net.Conn, filename string) error {
	contents, e := Read(filename)
	if e != nil {
		return e
	}

	_, e = c.Write(contents)
	if e != nil {
		return e
	}
	return nil
}

func upload(conn net.Conn, cmdline string) {
	// fields := strings.Fields(cmdline)
	// _, path := fields[0], fields[1]
	// if strings.Index(path, "/") < 0 || strings.Index(path, "\\") < 0 {
	// 	curtDir, _ := os.Getwd()
	// 	path = filepath.Join(curtDir, path)
	// }
	// fmt.Println("Uploading file......to " + path)
	// // f, e := os.Open(path)
	// // if e != nil {
	// // 	fmt.Println(e)
	// // }
	// // defer f.Close()
	// content := make([]byte, 256)
	// _, e := conn.Read(content)
	// if e != nil {
	// 	fmt.Println("Upload Error: " + e.Error())
	// 	return
	// }
	// file, e := os.Create(path)
	// if e != nil {
	// 	fmt.Println("Upload creating file Error: " + e.Error())
	// 	return
	// }
	// _, e = file.Write(content)
	// if e != nil {
	// 	fmt.Println("Upload Writing content to file Error: " + e.Error())
	// 	return
	// }
}
