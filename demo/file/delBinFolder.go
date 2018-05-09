package main

import (
	"fmt"
	"os"
)

func main() {
	var root = os.Args[1]
	fmt.Println(root)
	Del(root, "bin")
	Del(root, "obj")

}

func Del(root, toBeDel string) bool {
	rootDir, e := os.Open(root)

	defer rootDir.Close()
	if e != nil {
		fmt.Println(e.Error())
		return true
	}
	fis, e := rootDir.Readdir(0)
	if e != nil {
		fmt.Println(e.Error())
		return false
	}
	if len(fis) == 0 {
		fmt.Println(root + "---------- is empty")
		return true
	}

	for _, fi := range fis {

		if fi.IsDir() {
			tempRoot := root + `\` + fi.Name()
			fmt.Println(tempRoot)
			if fi.Name() == toBeDel {
				e = os.RemoveAll(tempRoot)
				if e != nil {
					fmt.Println(e.Error())
					continue
				}
				fmt.Println("----------delete")
			} else {
				Del(tempRoot, toBeDel)
			}
		}
	}
	return true
}
