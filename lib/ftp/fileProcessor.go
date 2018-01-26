package ftp

import (
	"os"
)

//Read content from the file name that was specified by youself
func Read(filename string) ([]byte, error) {

	f, e := os.Open(filename)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	contents := []byte{}
	_, e = f.Read(contents)
	if e != nil {
		return nil, e
	}
	return contents, nil
}

//Write content to the file name that was specified by youself
func Write(contents []byte, filename string) error {
	f, e := os.Create(filename)
	if e != nil {
		return e
	}
	defer f.Close()
	_, e = f.Write(contents)
	if e != nil {
		return e
	}
	return nil
}
