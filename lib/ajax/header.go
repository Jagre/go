package ajax

//GetDefaultHeader get the default Header
func (a *AjaxParam) GetDefaultHeader() *AjaxParam {
	a.Header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
	a.Header["Accept-Encoding"] = "gzip, deflate, sdch"
	a.Header["Accept-Language"] = "zh-CN,zh;q=0.8"
	a.Header["Connection"] = "keep-alive"
	a.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.10 Safari/537.36"
	//Host, Referer need to custom by self

	return a
}

//GetJSONHeader get the requestion header by json reqeusting
func (a *AjaxParam) GetJSONHeader() *AjaxParam {
	a.Header["Accept"] = "application/json, text/javascript, */*; q=0.01"
	a.Header["Accept-Encoding"] = "gzip, deflate"
	a.Header["Accept-Language"] = "zh-CN,zh;q=0.8"
	a.Header["Connection"] = "keep-alive"
	//a.Header["Content-Type"] = "application/json"
	//a.Header["X-Requested-With"] = "XMLHttpRequest"
	a.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.10 Safari/537.36"
	//Host, Referer need to custom by self

	return a
}
