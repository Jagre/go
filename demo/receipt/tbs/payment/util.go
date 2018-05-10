package payment

import (
	"strconv"
)

//ToOrderNo convert orderId to orderNo
func ToOrderNo(o interface{}) string {
	orderid := 0
	switch o.(type) {
	case string:
		o1 := o.(string)
		if len(o1) == 8 {
			return o1
		}
		orderid, _ = strconv.Atoi(o1)
	case int, int64:
		o1 := o.(int)
		if o1 > 80000000 {
			return strconv.Itoa(orderid)
		}
		orderid = o1
	}
	orderid = 80000000 + orderid
	return strconv.Itoa(orderid)
}

//ToOrderID convert orderNo to orderId
func ToOrderID(o interface{}) int64 {
	//o, _ := strconv.Atoi(orderno)
	orderno := int64(0)
	switch o.(type) {
	case string:
		o1 := o.(string)
		orderno, _ := strconv.Atoi(o1)
		if len(o1) < 8 {
			return int64(orderno)
		}
	case int, int64:
		o1 := o.(int64)
		if o1 < 10000000 {
			return o1
		}
		orderno = o1
	}

	return int64(orderno - 80000000)
}
