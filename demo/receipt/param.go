package receipt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//ReadOrders will read orderNos from file
func ReadOrders(fullPath string) []string {
	if len(fullPath) == 0 {
		fullPath = "./orderNos.txt"
	}
	return readLines(fullPath)
}

// ReceiptParam is param when the creating receipt
type ReceiptParam struct {
	OrderNo   string
	ReceiptNo string
}

// ReadReceiptParams will get receipt param
func ReadReceiptParams(fullPath string) []ReceiptParam {
	if len(fullPath) == 0 {
		fullPath = "./receipts.txt"
	}
	lines := readLines(fullPath)
	var result []ReceiptParam
	for _, l := range lines {
		items := strings.Split(l, ",")
		result = append(result, ReceiptParam{OrderNo: items[0], ReceiptNo: items[1]})
	}
	return result
}

func readLines(fullPath string) []string {
	f, e := os.Open(fullPath)
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var lines []string
	for {
		line, _, e := reader.ReadLine()
		if e != nil {
			if e != io.EOF {
				fmt.Println(e)
				os.Exit(1)
			}
			break
		}
		if len(line) > 0 {
			lines = append(lines, string(line))
		}
	}
	return lines
}
