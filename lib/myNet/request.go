//https://github.com/mikemintang/go-curl/blob/master/request.go

package myNet

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//MyRequest reqeust object
type MyRequest struct {
	client     *http.Client
	request    *http.Request
	RawRequest *http.Request
	URL        string
	Method     string //get; post
	header     map[string]string
	cookies    map[string]string
	rawQuery   string    //RawQuery
	postData   io.Reader //read byte to
}

//NewMyRequest is creation the instance of MyRequest
func NewMyRequest() *MyRequest {
	req := new(MyRequest)
	req.client = new(http.Client)
	return req
}

//SetURL method
func (req *MyRequest) SetURL(url string) *MyRequest {
	req.URL = url
	return req
}

//SetMethod method
func (req *MyRequest) SetMethod(method string) *MyRequest {
	if method == "" {
		method = "GET"
	}
	req.Method = strings.ToUpper(method)
	return req
}

// SetHeaders method
func (req *MyRequest) SetHeaders(header map[string]string) *MyRequest {
	if header == nil {
		return req
	}
	req.header = header
	return req
}

//setHeaders method is internal method
func (req *MyRequest) setHeaders() {
	if req.header != nil {
		for k, v := range req.header {
			req.request.Header.Set(k, v)
		}
	}
}

//SetCookies method
func (req *MyRequest) SetCookies(cookies map[string]string) *MyRequest {
	if cookies == nil {
		return req
	}
	req.cookies = cookies
	return req
}

//setCookies method is internal method
func (req *MyRequest) setCookies() {
	if req.cookies != nil {
		for k, v := range req.cookies {
			//TODO: jagre, mybe there is some issue that cookie's param is not completion
			req.request.AddCookie(&http.Cookie{
				Name:  k,
				Value: v})
		}
	}
}

//SetTLS set Transport Layer Security Protocol configuration for https access
//TLS is Transport Layer Security Protocol that base on SSL3.0(Secure Socket Layer Protocol)
//How to declare http.Transport:
// transport = &http.Transport{
//     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// }
// req.client.Transport = transport
func (req *MyRequest) SetTLS(transport *http.Transport) *MyRequest {
	if transport != nil {
		req.client.Transport = transport
	}
	return req
}

//setTLS set Transport Layer Security Protocol configuration for https access
//[NOTICE] just for https access
//TLS is Transport Layer Security Protocol that base on SSL3.0(Secure Socket Layer Protocol)
func (req *MyRequest) setTLS() *MyRequest {
	if strings.HasPrefix(req.URL, "https") {
		transport := http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		req.client.Transport = &transport
	}
	return req
}

//SetQueries method
func (req *MyRequest) SetQueries(queries map[string]string) *MyRequest {
	if queries == nil {
		return req
	}
	values := url.Values{}
	for k, v := range queries {
		values.Set(k, v)
	}
	req.rawQuery = values.Encode()
	return req
}

//SetPostData method
func (req *MyRequest) SetPostData(postData interface{}) *MyRequest {
	if postData != nil {
		data, e := json.Marshal(postData)
		if e != nil {
			panic("Parse postData raising error. ERR: " + e.Error())
		} else {
			req.postData = bytes.NewReader(data)
		}
	} else {
		req.postData = nil
	}
	return req
}

// Send request
func (req *MyRequest) Send() (*MyResponse, error) {
	if len(req.URL) == 0 {
		panic("No request url")
	}

	req2, e := http.NewRequest(req.Method, req.URL, req.postData)
	if e != nil {
		panic(e.Error())
	}
	req.request = req2

	req.setHeaders()
	req.setCookies()
	req.setTLS()
	if len(req.rawQuery) > 0 {
		req.request.URL.RawQuery = req.rawQuery
	}

	req.RawRequest = req.request
	rawRes, e := req.client.Do(req.request)
	if e != nil {
		return nil, e
	}
	defer rawRes.Body.Close()

	res := new(MyResponse)
	res.RawResponse = rawRes
	if !res.IsOK() {
		return nil, errors.New(rawRes.Status)
	}

	res.parseHeader()
	res.parseCookies()
	res.parseBody()
	return res, e
}
