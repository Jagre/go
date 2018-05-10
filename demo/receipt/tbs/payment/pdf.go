package payment

import (
	"fmt"
	"github.com/aaa/go/lib/myNet"
)

// CreatePdf will create pdf file
//oid : orderId
//cid : collectionId
func CreatePdf(orderid, collectionid int64, site string) bool {
	//orderNo := strings.
	url := fmt.Sprintf("%sController/Action?orderNo=80%d&collectionId=%d&processor=%s",
		site,
		orderid,
		collectionid,
		"System")
	fmt.Println("PDF: " + url)

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

	return res.IsOK()
}
