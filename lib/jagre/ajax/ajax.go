package ajax

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// AjaxParam is the param of requesting
type AjaxParam struct {
	URL      string
	Method   string //Get Post PostForm
	Data     []byte
	Header   map[string]string
	Sessions map[string]interface{}
	Cookies  []http.Cookie
}

// AjaxResult is the result of requested
type AjaxResult struct {
	Data     []byte
	Header   map[string]string
	Sessions map[string]interface{}
	Cookies  []http.Cookie
}

//Ajax request the url
func (a *AjaxParam) Ajax() (*AjaxResult, error) {
	if len(a.Method) == 0 {
		a.Method = "GET"
	}
	a.Method = strings.ToUpper(a.Method)
	// if a.Method == "GET" {
	// 	return requesting(a)
	// } else if a.Method == "POST" {
	// 	return requesting(a)
	// }

	return requesting(a)
}

// Get url's get
func requesting(a *AjaxParam) (*AjaxResult, error) {
	//create Request obj
	var req *http.Request
	var e error
	if a.Method == "GET" {
		req, e = http.NewRequest(a.Method, a.URL, nil)
	} else if a.Method == "POST" {
		body := bytes.NewReader(a.Data)
		req, e = http.NewRequest(a.Method, a.URL, body)
	}
	if e != nil {
		return &AjaxResult{}, e
	}

	//Set cookies
	for _, v := range a.Cookies {
		req.AddCookie(&v)
	}

	for k, v := range a.Header {
		req.Header.Add(k, v)
	}

	cli := http.Client{}
	resp, e := cli.Do(req)
	if e != nil {
		return &AjaxResult{}, e
	}
	defer resp.Body.Close()
	return result(resp)
}

// // Post url's post
// func post(a *AjaxParam) (*AjaxResult, error) {
// 	req, e := http.NewRequest(a.Method, a.URL, nil)
// 	if e != nil {
// 		return &AjaxResult{}, e
// 	}
// 	for _, v := range a.Cookies {
// 		req.AddCookie(&v)
// 	}

// 	cli := http.Client{}
// 	resp, e := cli.Do(req)
// 	return &AjaxResult{}, nil
// }

func result(resp *http.Response) (*AjaxResult, error) {
	aresult := AjaxResult{}
	//Get cookies
	for i, cookie := range resp.Cookies() {
		aresult.Cookies[i] = *cookie
	}

	//Get message body
	data, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return &AjaxResult{}, e
	}
	//Set value
	aresult.Data = data
	return &aresult, e
}

func convertType() {
	//type assertion
	// if v, ok := param.Data.(string); ok {
	// 	json.Unmarshal([]byte(v), param.Data)
	// } else if v, ok := param.Data.([]byte); ok {
	// 	json.Unmarshal(v, param.Data)
	// }
	// if _, ok := param.Data.(map[string]interface{}); !ok {
	// 	panic("param type is error, just string json || []byte json || map[string]interface{}")
	// }

	// //interface{}转为普通类型必须通过类型断言：v, ok := obj.(T) || t := obj.(type)
	// values := url.Values{}
	// for k, v := range param.Data.(map[string]interface{}) {
	// 	v2 := ""
	// 	switch t := v.(type) {
	// 	case string:
	// 		v2, _ := v.(string)
	// 	case int:
	// 		temp, _ := v.(int)
	// 		v2 = string(temp)
	// 	}
	// 	values.Set(k, v2)
	// }
}

//Convert result to next requestion's param
func (a *AjaxResult) Convert() *AjaxParam {
	p := AjaxParam{
		Cookies:  a.Cookies,
		Sessions: a.Sessions,
		Header:   a.Header}

	return &p
}
