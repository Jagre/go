package payment

import (
	"fmt"
	"os"
	"strconv"
	"time"
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

//Log is log error message
func Log(err error) {
	msg := fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
	var f *os.File
	var e error
	fileName := "./errorlog.txt"
	if !fileIsExist(fileName) {
		f, e = os.Create(fileName)
		if e != nil {
			fmt.Println(e)
		}
	} else {
		f, e = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0777)
		if e != nil {
			fmt.Println(e)
		}
	}
	defer f.Close()

	off, _ := f.Seek(0, os.SEEK_END)
	f.WriteAt([]byte(msg), off)
	//ioutil.WriteFile(`./errorLog.txt`, []byte(msg), 0777)
}

//Info is log operation data
func Info(msg string) {
	msg = fmt.Sprintf("%s: %s", time.Now().Format("2006-01-02 15:04:05"), msg)
	var f *os.File
	var e error
	fileName := "./info.txt"
	if !fileIsExist(fileName) {
		f, e = os.Create(fileName)
		if e != nil {
			fmt.Println(e)
		}
	} else {
		f, e = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0777)
		if e != nil {
			fmt.Println(e)
		}
	}
	defer f.Close()

	off, _ := f.Seek(0, os.SEEK_END)
	f.WriteAt([]byte(msg), off)
	//ioutil.WriteFile(`./info.txt`, []byte(msg), 0777)
}

func fileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
