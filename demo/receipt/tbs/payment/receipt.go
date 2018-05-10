package payment

import (
	"fmt"
	"github.com/aaa/go/lib/myNet"
)

// CreateReceipt will create receipt
func CreateReceipt(orderno, receiptno, operator, employeeid, site string) bool {
	url := fmt.Sprintf("%sController/Action?orderNo=%s&receiptNo=%s&processor=%s&employeeId=%s",
		site,
		orderno,
		receiptno,
		operator,
		employeeid)
	fmt.Println("Receipt: " + url)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"
	req := myNet.NewMyRequest()
	res, e := req.SetURL(url).
		SetMethod("get").
		SetHeaders(headers).
		Send()
	if e != nil {
		fmt.Println(e)
		return false
	}
	defer res.RawResponse.Body.Close()
	return true
}
