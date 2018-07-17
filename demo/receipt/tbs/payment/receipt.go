package payment

import (
	"fmt"
	"github.com/aaa/go/lib/myNet"
)

// CreateReceipt will create receipt
func CreateReceipt(orderid int64, receiptno, operator, employeeid, url string) bool {
	url = fmt.Sprintf("%sReceipt/CreateReceipt?orderNo=80%d&receiptNo=%s&processor=%s&employeeId=%s",
		url,
		orderid,
		receiptno,
		operator,
		employeeid)
	Info(fmt.Sprintf("Receipt: %s\n", url))

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"
	req := myNet.NewMyRequest()
	res, e := req.SetURL(url).
		SetMethod("get").
		SetHeaders(headers).
		Send()
	if e != nil {
		fmt.Println(e)
		Log(e)
		return false
	}
	defer res.RawResponse.Body.Close()
	return res.IsOK()
}
