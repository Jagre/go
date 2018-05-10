package payment

import (
	"strconv"
	"strings"
)

//ToOrderNo convert orderId to orderNo
func ToOrderNo(o interface{}) string {
	orderid := 0
	switch o.(type) {
	case string:
		o1 := o.(string)
		if len(o1) == 9 {
			return o1
		}
		orderid, _ = strconv.Atoi(o1)
	case int, int64:
		o1 := o.(int)
		orderid = o1
	}
	strOrderid := strconv.Itoa(orderid)
	return "B" + strings.Repeat("0", 8-len(strOrderid)) + strOrderid
}

//ToOrderID convert orderNo to orderId
func ToOrderID(o interface{}) int64 {
	//o, _ := strconv.Atoi(orderno)
	orderid := int64(0)
	switch o.(type) {
	case string:
		o1 := o.(string)
		o1 = strings.Replace(o1, "B", "", -1)
		tempOrderid, _ := strconv.Atoi(o1)
		orderid = int64(tempOrderid)
	case int, int64:
		orderid = o.(int64)
	}
	return int64(orderid)
}
