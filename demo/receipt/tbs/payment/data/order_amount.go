package data

import (
	"fmt"
	"github.com/aaa/go/demo/receipt/tbs/payment"
	"github.com/aaa/go/lib/sql"
)

// GetOrderIDsByReceiptTime get orders' id by the receipt time
func GetOrderIDsByReceiptTime(conn *sql.MSSQLConnection, from, to string) *[]int64 {
	db, e := conn.Open()
	if e != nil {
		payment.Log(e)
		return nil
	}
	defer db.Close()

	//取出所有未生成PDF的有效收据的订单
	query := fmt.Sprintf(`
			Select	distinct OrderID
			From	DepartureGroupOrderReceiptLog (nolock) 
			Where	ReceiptStatus > 0 
					and CreatedDate >= '%s' 
					and CreatedDate <= '%s' 
			Order by OrderID desc`, from, to)
	rows, e := db.Query(query)
	if e != nil {
		payment.Log(e)
		return nil
	}
	orderids := make([]int64, 0)
	for rows.Next() {
		orderid := int64(0)
		// 查询结果字段和声明变量数量相等，否则数据为空。
		rows.Scan(&orderid)
		orderids = append(orderids, orderid)
	}
	return &orderids
}

// Order is order entity
type Order struct {
	OrderID            int64
	OrderNo            string
	TotalAmount        float32
	TotalReceiveAmount float32
	TotalTICCostAmount float32
	PaymentStatus      byte
	OrderStatus        byte
	SalesChannel       byte
}

// GetOrderByID get the order entity by orderId
func GetOrderByID(conn *sql.MSSQLConnection, orderid int64) *Order {
	db, err := conn.Open()
	if err != nil {
		payment.Log(err)
		return nil
	}
	defer db.Close()
	// 执行SQL语句
	query := fmt.Sprintf(`
			select	OrderID, OrderNo, TotalAmount, TotalReceiveAmount, TotalTICCostAmount, PaymentStatus, OrderStatus, SalesChannel 
			from	DepartureGroupOrder (nolock) 
			where 	OrderStatus < 3
					AND OrderID=%d`, orderid)
	rows, err := db.Query(query)
	if err != nil {
		payment.Log(err)
		return nil
	}
	order := new(Order)
	for rows.Next() {
		rows.Scan(&order.OrderID,
			&order.OrderNo,
			&order.TotalAmount,
			&order.TotalReceiveAmount,
			&order.TotalTICCostAmount,
			&order.PaymentStatus,
			&order.OrderStatus,
			&order.SalesChannel)

		//Log to inforamtion file
		payment.Info(fmt.Sprintf("OrderID: %d; OrderNo: %s; TotalReceiveAmount: %f; PaymentStatus: %d; OrderStatus: %d; SalesChannel: %d\n", order.OrderID, order.OrderNo, order.TotalReceiveAmount, order.PaymentStatus, order.OrderStatus, order.SalesChannel))
	}
	return order
}

// IsNeedTic will check data whether count TIC
func IsNeedTic(conn *sql.MSSQLConnection, orderid int64) (bool, error) {
	db, err := conn.Open()
	if err != nil {
		payment.Log(err)
		return false, err
	}
	defer db.Close()
	// 执行SQL语句
	query := fmt.Sprintf(`
		Select	COUNT(0) As isNeedTic
		From	DepartureGroupOrderChargeDetail (nolock) 
		Where  	ChargeType = 16
				AND OrderID=%d`, orderid)
	isNeedTic := false
	err = db.QueryRow(query).Scan(&isNeedTic)
	if err != nil {
		payment.Log(err)
		return false, nil
	}
	return isNeedTic, nil
}

// UpdateTICChargeDetail update tic item
func UpdateTICChargeDetail(conn *sql.MSSQLConnection, orderid int64, ticAmount float32) bool {
	db, err := conn.Open()
	if err != nil {
		payment.Log(err)
		return false
	}
	defer db.Close()
	script := `
		UPDATE  DepartureGroupOrderChargeDetail With (Rowlock) 
		SET     Price = $1,
				OriginalAmount = $1,
				HKDAmount = $1
		WHERE   OrderID = $2 
				AND ChargeType = 16`

	script = sql.SQLScript(script).Parse(ticAmount, orderid)
	_, err = db.Exec(script)
	if err != nil {
		payment.Log(fmt.Errorf("UpdateTICChargeDetail.Exec: %s", err.Error()))
		return false
	}
	return true
}

// UpdateOrder will update order amounts & order status
func UpdateOrder(conn *sql.MSSQLConnection, orderid int64, totalAmount, totalReceiveAmount, totalPromoCodeAmount, totalTicAmount, totalCouponAmount, totalEmoneyAmount float32) bool {

	db, err := conn.Open()
	if err != nil {
		payment.Log(err)
		return false
	}

	script := fmt.Sprintf(`
		Update 	DepartureGroupOrder with (rowlock)
		SET		TotalAmount = %f,
				TotalReceiveAmount = %f,
				TotalMarketCodeAmount = %f,
				TotalTICCostAmount = %f,
				TotalCouponAmount = %f,
				TotalEmoneyAmount = %f,
				PaymentStatus = CASE
									WHEN $2 >= $1 THEN 2
									WHEN $2 > 0 AND $2 < $1 THEN 1
									WHEN $2 <= 0 THEN 0
								END,
				CurrentPaymentStatus = 1,
				OrderStatus = 2
		Where	OrderID = $3`, totalAmount, totalReceiveAmount, totalPromoCodeAmount, totalTicAmount, totalCouponAmount, totalEmoneyAmount)

	script = sql.SQLScript(script).Parse(totalAmount, totalReceiveAmount, orderid)
	_, err = db.Exec(script)
	if err != nil {
		payment.Log(fmt.Errorf("UpdateOrder.Exec: %s", err.Error()))
		return false
	}
	return true
}
