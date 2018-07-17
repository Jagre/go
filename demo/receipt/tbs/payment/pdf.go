package payment

import (
	"fmt"
	"github.com/aaa/go/lib/myNet"
)

// CreatePdf will create pdf file
//oid : orderId
//cid : collectionId
func CreatePdf(orderid, collectionid int64, url string) bool {
	//orderNo := strings.
	url = fmt.Sprintf("%sReceipt/CreateReceiptPDFByCollectionId?orderNo=80%d&collectionId=%d&processor=%s", //"%sController/Action?orderNo=80%d&collectionId=%d&processor=%s",
		url,
		orderid,
		collectionid,
		"System")
	Info(fmt.Sprintf("PDF: %s\n", url))

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

func CreatePdfByReceipt(orderid int64, receiptNo, url string) bool {
	//orderNo := strings.
	url = fmt.Sprintf("%sReceipt/CreateReceiptPDFByReceiptNo?orderNo=80%d&receiptNo=%s&processor=%s", //"%sController/Action?orderNo=80%d&collectionId=%d&processor=%s",
		url,
		orderid,
		receiptNo,
		"System")
	Info(fmt.Sprintf("PDF: %s\n", url))

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
