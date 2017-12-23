package myNet

import (
	"io/ioutil"
	"net/http"
)

//MyResponse is response object
type MyResponse struct {
	RawResponse *http.Response
	Header      *map[string]string
	Cookies     *map[string]string
	Body        string
}

//NewMyResponse will create a instance of the type MyResponse
func NewMyResponse() *MyResponse {
	return new(MyResponse)
}

//IsOK express response's status is OK
func (res *MyResponse) IsOK() bool {
	return res.RawResponse.StatusCode == 200
}

func (res *MyResponse) parseHeader() {
	header := map[string]string{}
	for k, v := range res.RawResponse.Header {
		header[k] = v[0]
	}
	res.Header = &header
}

func (res *MyResponse) parseCookies() {
	rawCookies := res.RawResponse.Cookies()
	cookies := map[string]string{}
	for _, v := range rawCookies {
		cookies[v.Name] = v.Value
	}
	res.Cookies = &cookies
}

func (res *MyResponse) parseBody() {
	body, e := ioutil.ReadAll(res.RawResponse.Body)
	if e != nil {
		panic(e.Error())
	}
	res.Body = string(body)
}
