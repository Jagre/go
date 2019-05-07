package payment

import (
	"errors"
	"fmt"
	"github.com/aaa/go/lib/sql"
	"time"
)

type ReceiptInfo struct {
	OrderId          int64
	ReceiptNo        string
	HasReceiptCotent bool
}

func GetReceipts(conn *sql.MSSQLConnection) *[]ReceiptInfo {
	receipts := []ReceiptInfo{}
	db, e := conn.Open()
	if e != nil {
		Log(e)
		return nil
	}
	defer db.Close()

	before5Day := time.Now().AddDate(0, 0, -7)
	//取出所有未生成PDF的有效收据的订单
	sql := fmt.Sprintf(`
		Select	distinct orderId, accountReceiptId as receitpNo, receiptContent
		From	DepartureGroupOrderReceiptLog (nolock) 
		Where	ReceiptPDFUrl = '' 
				and ReceiptStatus > 0 
				and CreatedDate > '%s' 
		Order by OrderID desc`, before5Day.Format("2006-01-02"))
	rows, e := db.Query(sql)
	if e != nil {
		Log(e)
		return nil
	}
	for rows.Next() {
		orderId := int64(0)
		receiptNo := ""
		receiptContent := ""

		// 查询结果字段和声明变量数量相等，否则数据为空。
		rows.Scan(&orderId, &receiptNo, &receiptContent)
		receipts = append(receipts, ReceiptInfo{OrderId: orderId, ReceiptNo: receiptNo, HasReceiptCotent: receiptContent > ""})
	}
	if len(receipts) > 0 {
		orderId := int64(0)
		for _, item := range receipts {
			if orderId != item.OrderId {
				orderId = item.OrderId
				Info(fmt.Sprintf("OrderID: %d\n", orderId))
			}
		}
	}

	return &receipts
}

func GetOrder(conn *sql.MSSQLConnection, orderid int64) bool {
	db, e := conn.Open()
	if e != nil {
		Log(e)
		return false
	}
	defer db.Close()

	sql := fmt.Sprintf(`Select distinct orderid From DepartureGroupOrderReceiptLog (nolock) Where OrderID = %d and ReceiptPDFUrl = '' and ReceiptStatus > 0`, orderid)
	rows, e := db.Query(sql)
	if e != nil {
		Log(e)
		return false
	}
	for rows.Next() {
		orderid := int64(0)
		// 查询结果字段和声明变量数量相等，否则数据为空。
		rows.Scan(&orderid)
		if orderid > 0 {
			Log(errors.New(fmt.Sprintf("Error, OrderID(%d) no PDF yet\n", orderid)))
			return false
		}
	}

	return true
}

func GetReceiptsByOrderId(conn *sql.MSSQLConnection, orderid int64) []string {
	receiptNos := make([]string, 0)
	db, e := conn.Open()
	if e != nil {
		Log(e)
		return nil
	}
	defer db.Close()

	sql := fmt.Sprintf(`Select distinct AccountReceiptID as receiptno From DepartureGroupOrderReceiptLog (nolock) Where OrderID = %d and ReceiptPDFUrl = '' and ReceiptStatus > 0`, orderid)
	rows, e := db.Query(sql)
	if e != nil {
		Log(e)
		return nil
	}
	for rows.Next() {
		receiptno := ""
		// 查询结果字段和声明变量数量相等，否则数据为空。
		rows.Scan(&receiptno)
		receiptNos = append(receiptNos, receiptno)
	}
	return receiptNos
}

func GetOrderStatus(conn *sql.MSSQLConnection, orderid int64) (PaymentStatus, OrderStatus, SalesChannel byte) {
	db, err := conn.Open()
	if err != nil {
		Log(err)
		return
	}
	defer db.Close()

	// 执行SQL语句
	var query = fmt.Sprintf("select PaymentStatus, OrderStatus, SalesChannel from DepartureGroupOrder (nolock) where orderId=%d", orderid)
	rows, err := db.Query(query)
	if err != nil {
		Log(err)
		return
	}
	for rows.Next() {
		// 查询结果字段和声明变量数量相等，否则数据为空。
		rows.Scan(&PaymentStatus, &OrderStatus, &SalesChannel)
		Info(fmt.Sprintf("PaymentStatus: %d; OrderStatus: %d; SalesChannel: %d\n", PaymentStatus, OrderStatus, SalesChannel))
	}
	return PaymentStatus, OrderStatus, SalesChannel
}
