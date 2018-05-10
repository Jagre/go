package payment

import (
	"encoding/json"
	"fmt"
	"github.com/aaa/go/lib/myNet"
)

//OrderQueryRequest 參數
type orderQueryRequest struct {
	OrderID         int64
	UserID          int
	UserName        string
	OrgID           int
	OrgCode         string
	IP              string
	IsShowExtraInfo bool
}

//PaymentInfo  收款頁面結果（主要取CollectionID）
type paymentInfo struct {
	OrderID   int64
	OrderNo   string
	Receipted paymentReceiptedInfo
}

type paymentReceiptedInfo struct {
	TotalReceivedAmount float32
	Collections         []paymentCollectionInfo
}

type paymentCollectionInfo struct {
	CollectionID int64
}

//GetCollectionIDs will parse collectionIds from json result of the payment page
func GetCollectionIDs(site string, orderid int64) []int64 {
	url := site + "Controller/Action"
	req := myNet.NewMyRequest()

	param := orderQueryRequest{
		OrderID: orderid,
	}
	headers := make(map[string]string, 2)
	headers["Content-Type"] = "application/json; charset=utf-8"

	res, e := req.SetURL(url).
		SetMethod("post").
		SetPostData(param).
		SetHeaders(headers).
		Send()
	if e != nil {
		fmt.Println(e)
	}
	defer res.RawResponse.Body.Close()

	pi := new(paymentInfo)
	json.Unmarshal([]byte(res.Body), pi)

	var collectionIds []int64
	for _, c := range pi.Receipted.Collections {
		collectionIds = append(collectionIds, c.CollectionID)
	}
	return collectionIds
}
