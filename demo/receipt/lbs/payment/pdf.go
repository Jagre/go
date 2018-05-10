package payment

import (
	"fmt"
	"github.com/aaa/go/lib/myNet"
	"strings"
)

// CreatePdf will create pdf file
//oid : orderId
//cid : collectionId
func CreatePdf(orderid int64, receiptNo, pdfSite, bookingAPI string) (bool, string) {
	url := fmt.Sprintf("%sController/Action?orderId=%d&receiptNo=%s",
		pdfSite,
		orderid,
		receiptNo)
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
		return false, ""
	}
	defer res.RawResponse.Body.Close()
	pdfURL := strings.Replace(res.Body, `"`, "", -1)

	updatePdfURL(orderid, receiptNo, pdfURL, bookingAPI)

	return res.IsOK(), res.Body
}

func updatePdfURL(orderid int64, receiptNo, pdfURL, bookingAPI string) bool {
	url := fmt.Sprintf("%sOrderReceipt/UpdatePDFUrl?orderId=%d&receiptNo=%s&pdfUrl=%s&proccessor=System",
		bookingAPI,
		orderid,
		receiptNo,
		pdfURL)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"
	req := myNet.NewMyRequest()
	res, e := req.SetURL(url).SetMethod("get").SetHeaders(headers).Send()
	if e != nil {
		fmt.Println(e)
		return false
	}

	defer res.RawResponse.Body.Close()

	return res.IsOK()
}
