package payment

import (
	"encoding/json"
	"fmt"
	"github.com/aaa/go/lib/myNet"
	"github.com/buger/jsonparser"
)

// Result is Acount API's result sturcture
// type result struct {
// 	head head
// 	data AccountAmountDto
// }

// // Head account's head
// type head struct {
// 	Code    int
// 	Message string
// }

// AccountAmountDto is amount entity
type AccountAmountDto struct {
	TotalAmount         float32
	PaymentAmountDic    map[string]float32
	PromotionCodeAmount float32
	EMoneyAmount        float32
}

// GetAccountAmounts will get the account's paymented
func GetAccountAmounts(url, orderno string) *AccountAmountDto {
	url = url + fmt.Sprintf("Collection/GetExistingAmount?orderno=%s", orderno)
	req := myNet.NewMyRequest()

	headers := make(map[string]string, 2)
	headers["Content-Type"] = "application/json; charset=utf-8"

	res, e := req.SetURL(url).
		SetMethod("get").
		SetHeaders(headers).
		SetTLS(nil).
		Send()
	if e != nil {
		Log(e)
		return nil
	}
	defer res.RawResponse.Body.Close()
	// take off shell
	data, _, _, _ := jsonparser.Get([]byte(res.Body), "data")
	fmt.Printf("%v\n", string(data))

	dto := &AccountAmountDto{}
	json.Unmarshal(data, dto)
	return dto
}
