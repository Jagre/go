package find

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var reg *regexp.Regexp

func Finding(rootPath, pattern string, isRegular, isIgnoreCase bool) {
	if isIgnoreCase {
		isRegular = true
		pattern = fmt.Sprintf(`(?i)%s`, pattern)
	}

	if isRegular {
		reg = regexp.MustCompile(pattern)
		walkDirByRegular(rootPath)
	} else {
		walkDirByString(rootPath, pattern)
	}
}

func walkDirByRegular(rootPath string) {
	filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			return nil
		}
		fileName := d.Name()
		if strings.HasSuffix(fileName, ".csproj") ||
			strings.HasSuffix(fileName, ".cshtml") ||
			strings.HasSuffix(fileName, ".cs") ||
			strings.HasSuffix(fileName, ".js") ||
			strings.HasSuffix(fileName, ".json") {
			opFileByRegular(path)
		}
		return nil
	})
}

func opFileByRegular(path string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	loc := reg.FindIndex(f)
	if loc != nil {
		fmt.Printf("	%s (%d, %d)\r\n", path, loc[0], loc[1])
	}
}

func walkDirByString(rootPath, pattern string) {
	fs.WalkDir(os.DirFS(rootPath), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal()
		}
		if d.IsDir() {
			return nil
		}

		fileName := d.Name()
		if strings.HasSuffix(fileName, ".csproj") ||
			strings.HasSuffix(fileName, ".cshtml") ||
			strings.HasSuffix(fileName, ".cs") ||
			strings.HasSuffix(fileName, ".js") ||
			strings.HasSuffix(fileName, ".json") {
			opFileByString(path, pattern)
		}
		return nil
	})
}

func opFileByString(path, pattern string) {
	// in order to read file line by line
	f, err := os.OpenFile(path, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rowIndex := 0
	for scanner.Scan() {
		// read a line content by scanner
		line := scanner.Text()
		rowIndex++
		if len(line) > 0 {
			colIndex := strings.Index(line, pattern)
			if colIndex >= 0 {
				fmt.Printf("	%s (%d, %d)\r\n", path, rowIndex, colIndex)
				return
			}
		}
	}
}
