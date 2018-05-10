package payment

import (
	"fmt"
	"github.com/aaa/go/lib/myNet"
)

// CreateReceipt will create receipt
func CreateReceipt(orderNo, receiptno, site string) bool {
	url := fmt.Sprintf("%sController/Action?orderNo=%s&receiptNo=%s&proccessor=%s&buildPDF=true",
		site,
		orderNo,
		receiptno,
		"")
	fmt.Println("Receipt: " + url)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"
	req := myNet.NewMyRequest()
	res, e := req.SetURL(url).
		SetMethod("get").
		SetHeaders(headers).
		//SetPostData([]byte(`{"orderId": ` + orderId + `, "receiptNo": ` + receiptno + `}`)).
		Send()

	if e != nil {
		fmt.Println(e)
		return false
	}
	defer res.RawResponse.Body.Close()
	return true
}
