package ftp

import (
	"net"
)

//Pull client download file from server (run at server)
//param: c never close
func Pull(c net.Conn, filename string) error {
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

func download(conn net.Conn, cmdline string) {
	// fields := strings.Fields(cmdline)
	// _, path := fields[0], fields[1]
	// if strings.Index(path, "/") < 0 || strings.Index(path, "\\") < 0 {
	// 	curtDir, _ := os.Getwd()
	// 	path = filepath.Join(curtDir, path)
	// }
	// fmt.Println("Downloading file......from " + path)
	// f, e := os.Open(path)
	// if e != nil {
	// 	fmt.Println(e)
	// }
	// defer f.Close()

	// content, e := ioutil.ReadAll(f)
	// if e != nil && e != io.EOF {
	// 	fmt.Println(e)
	// 	os.Exit(1)
	// }
	// //Write content to client
	// _, e = conn.Write(content)
	// if e != nil {
	// 	fmt.Println("Download tranport to client raise error: " + e.Error())
	// }
	// fmt.Println(string(content))
}
