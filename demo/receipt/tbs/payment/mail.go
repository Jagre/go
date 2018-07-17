package payment

import (
	"fmt"
	"github.com/aaa/go/lib/myNet"
)

// MailSender will create receipt
func MailSender(param interface{}, url string) bool {
	Info(fmt.Sprintf("Mail2: %s\n", url))

	//define the struct model and set values. eg:
	// param := mailParam{
	// 	OrderID:   236626,
	// 	MailType:  12,
	// 	ProcessBy: "haha",
	// }

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"
	req := myNet.NewMyRequest()

	res, e := req.SetURL(url).
		SetMethod("post").
		SetPostData(param).
		SetHeaders(headers).
		SetTLS(nil).
		Send()

	if e != nil {
		Log(e)
		return false
	}
	defer res.RawResponse.Body.Close()
	return true
}

func GetMailParam(url string, orderId int64, mailType byte) string {
	url = fmt.Sprintf("%s?orderId=%d&mailType=%d", url, orderId, mailType)
	Info(fmt.Sprintf("Mail1: %s\n", url))

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"
	req := myNet.NewMyRequest()

	res, e := req.SetURL(url).
		SetMethod("get").
		SetHeaders(headers).
		SetTLS(nil).
		Send()

	if e != nil {
		Log(e)
		return ""
	}
	defer res.RawResponse.Body.Close()
	return res.Body
}
