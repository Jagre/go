package main

//https://github.com/lnx/stockforces

import (
	"encoding/json"
	"fmt"

	//"github.com/lxn/go-winapi"
	//"github.com/lxn/win"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	//SinaURLPattern is sina url pattern
	//SinaURLPattern = "https://finance.sina.com.cn/realstock/company/%s/nc.shtml"
	SinaURLPattern = "http://hq.sinajs.cn/rn=%s&list=%s"
)

func main() {
	stocks := reading()
	for {
		if len(stocks) > 0 {
			msgs := ""
			for name, no := range stocks {
				//you can try to run with goroutine
				stockDatas := getStockDataOnSina(no)
				open, _ := strconv.ParseFloat(stockDatas[1], 10)
				yesterday, _ := strconv.ParseFloat(stockDatas[2], 10)
				price, _ := strconv.ParseFloat(stockDatas[3], 10)
				up, _ := strconv.ParseFloat(stockDatas[4], 10)
				down, _ := strconv.ParseFloat(stockDatas[5], 10)
				scope := price - yesterday
				scopePercent := (scope / price) * 100
				msg := fmt.Sprintf("%s: %3.2f(%3.2f, %3.2f%%), Open: %3.2f, U-D: %3.2f - %3.2f", name, price, scope, scopePercent, open, up, down)
				msgs = fmt.Sprintf("%s\n\n%s", msgs, msg)
			}
			if len(msgs) > 0 {
				ShowMessage(msgs)
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func getStockDataOnSina(stockNo string) []string {
	tick := strconv.FormatInt(time.Now().UnixNano(), 10)[:13] //nanoseconds
	url := fmt.Sprintf(SinaURLPattern, tick, stockNo)
	res, _ := http.Get(url)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	result := strings.Split(string(body), "=")[1]
	return strings.Split(result, ",")
}

func reading() map[string]string {
	content, err := ioutil.ReadFile("./stock.json")
	if err != nil {
		panic(err.Error())
	}
	stocks := make(map[string]string, 3)
	json.Unmarshal(content, &stocks)
	return stocks
}

func toUint16(_str string) *uint16 {

	return syscall.StringToUTF16Ptr(_str)
}

func toString(_n int32) string {
	return strconv.Itoa(int(_n))
}

//ShowMessage will show your message
func ShowMessage(msg string) {
	//var hwnd win.HWND
	//xScreen := winapi.GetSystemMetrics(winapi.SM_CXSCREEN)
	//yScreen := winapi.GetSystemMetrics(winapi.SM_CYSCREEN)

	//winapi.MessageBox(hwnd, toUint16(msg), toUint16(title), winapi.MB_OK)
	//win.MessageBox(hwnd, hwnd, toUint16(msg), toUint16(title), win.MB_OK)
	fmt.Printf("%s\n", msg)
}
