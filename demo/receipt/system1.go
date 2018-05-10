package receipt

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

//re-create receipt
func main() {
	params := *parseParam()
	testurl := ""
	//produrl := ""
	for _, p := range params {
		//fmt.Printf("OrderNo: %s, ReceiptNO: %s, Operator: %s \n", p.OrderNo, p.ReceiptNO, p.Operator)
		creating(testurl, p.OrderNo, p.ReceiptNO, p.Operator, "001")
		fmt.Printf("ReceiptNo(%s) was been sent! \n", p.ReceiptNO)
		time.Sleep(time.Second * time.Duration(1))
	}
}

func creating(baseURL, orderNo, receiptNo, operator, employeeID string) {
	url := fmt.Sprintf("%s?orderNo=%s&receiptNo=%s&processor=%s&employeeId=%s", baseURL, orderNo, receiptNo, operator, employeeID)
	res, e := http.Get(url)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%s is ok\n", receiptNo)
	}
	defer res.Body.Close()
}

type tbsReceiptParam struct {
	OrderNo   string
	ReceiptNO string
	Operator  string
}

// Notice: the param formatter as follow
// 90001234,RCXO0001234,System
// 90001235,RCXO0001235,System
func parseParam() *[]tbsReceiptParam {

	params := []tbsReceiptParam{}
	f, e := os.Open("./param.txt")
	if e != nil {
		fmt.Printf("When getting receipt's param from the file \"./param.txt\" raise error: %s\n", e)
	}
	defer f.Close()

	//read content row by row
	reader := bufio.NewReader(f)
	for {
		line, _, e := reader.ReadLine()
		if line == nil && e != nil {
			break
		}

		items := strings.Split(string(line), ",")
		param := tbsReceiptParam{OrderNo: items[0], ReceiptNO: items[1], Operator: items[2]}
		params = append(params, param)
	}
	return &params
}
