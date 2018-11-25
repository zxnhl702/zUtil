// Package zUtil 发送请求的函数以及数据结构第二版
package zUtil

import (
    "crypto/tls"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"
)

// ReqParam2 请求结构体
type ReqParam2 struct {
    Addr        string          // 请求地址
    Fulladdr    string          // 完整请求地址
    Param       url.Values      // 请求参数(放在请求url上的)
    Body        url.Values      // 请求参数(放在请求body内的)
    Body2       string          // 请求参数(放在请求body内的 已经拼完的字符串)
    Cookies     []*http.Cookie  // cookie
    ContentType string          // contenttype
    Method      string          // 请求方法get/post
    Header      http.Header     // 请求头
    UserAgent   string          // UserAgent
}

// NewReqParam2 新建请求结构体
func NewReqParam2() *ReqParam2 {
    return &ReqParam2{
        Addr        : "",
        Fulladdr    : "",
        Param       : make(url.Values),
        Body        : make(url.Values),
        Body2       : "",
        Cookies     : make([]*http.Cookie, 0),
        ContentType : ContentTypeUrlencoded,
        Method      : "",
        Header      : make(http.Header),
        UserAgent   : "",
    }
}

// ============================
// GET请求
// ============================

// SendGetRequest 发送get请求(client模式 不带cookie)
func (r *ReqParam2)SendGetRequest() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = r.getFullURL(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 发送请求
    data, err = r.doSendGetWithURL(r.Fulladdr, r.Header)
    return data, err
}

// SendGetRequestWithCookie 发送get请求(client模式 带cookie)
func (r *ReqParam2)SendGetRequestWithCookie() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = r.getFullURL(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 生成请求
    req, err := r.genGet()
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 发送请求
    data, err = r.doSendWithReq(req)
    return data, err
}

// SendGetRequestToGetDataAndCookie 发送get请求 获取返回数据和cookie
func (r *ReqParam2)SendGetRequestToGetDataAndCookie() ([]byte, []*http.Cookie, error) {
    var data []byte
    var err error
    cookies := make([]*http.Cookie, 0)
    // 完整的请求地址
    r.Fulladdr, err = r.getFullURL(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, cookies, err
    }
    // 生成请求
    req, err := r.genGet()
    if nil != err {
        log.Println(err)
        return data, cookies, err
    }
    // 发送请求
    resp, err := r.doSendWithReqToGetResp(req)
    if nil != err {
        log.Println(err)
        return data, cookies, err
    }
    defer resp.Body.Close()
    // 获取cookie
    cookies = resp.Cookies()
    // 获取返回数据
    data, err = ioutil.ReadAll(resp.Body)
    return data, cookies, err
}

// 生成get请求
func (r *ReqParam2)genGet() (*http.Request, error) {
    // 生成请求
    req, err := http.NewRequest("GET", r.Fulladdr, nil)
    if nil != err {
        log.Println("生成请求失败")
        return nil, err
    }
    // 设置请求头
    // 自定义的请求头
    for k, v := range r.Header {
        for _, vv := range v {
            req.Header.Set(k, vv)
        }
    }
    // 设置cookies
    for _, c := range r.Cookies {
        req.AddCookie(c)
    }
    // 设置Content-Type
    if ContentTypeUrlencoded != r.ContentType {
        req.Header.Set("Content-Type", r.ContentType)
    }
    // 设置user agent
    if "" != r.UserAgent {
        req.Header.Set("User-Agent", r.UserAgent)
    }
    return req, nil
}

// 发送post请求 p post请求body中的参数  apiurl 完整的请求地址 带自定义请求头
func (r *ReqParam2)doSendGetWithURL(apiurl string, header http.Header) ([]byte, error) {
    var rtn []byte
    // 指定不验证服务器端证书 生成客户端
    tr := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // 生成请求
    req, err := http.NewRequest("GET", apiurl, nil)
    if nil != err {
        log.Println("生成请求失败")
        return rtn, err
    }
    // 设置请求头
    // 自定义的请求头
    for k, v := range header {
        for _, vv := range v {
            req.Header.Set(k, vv)
        }
    }
    // 发送post请求
    resp, err := client.Do(req)
    if nil != err {
        log.Println("发送请求失败")
        return rtn, err
    }
    log.Println("resp.Cookies:", resp.Cookies())
    defer resp.Body.Close()
    // 读取返回参数
    rtn, err = ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if nil != err {
        log.Println("返回数据读取失败")
        return rtn, err
    }
    return rtn, nil
}

// ============================
// POST请求
// ============================

// SendPostRequest 发送post请求(client模式 不带cookie)
func (r *ReqParam2)SendPostRequest() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = r.getFullURL(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 拼请求body参数
    r.setRequestBody2()
    // 发送请求
    data, err = r.doSendPostWithURL(r.Body2, r.ContentType, r.Fulladdr, r.Header)
    return data, err
}

// SendPostRequestWithCookie 发送get请求(client模式 带cookie)
func (r *ReqParam2)SendPostRequestWithCookie() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = r.getFullURL(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 拼请求body参数
    r.setRequestBody2()
    // 生成请求
    req, err := r.genPost()
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 发送请求
    data, err = r.doSendWithReq(req)
    return data, err
}

// SendPostRequestToGetDataAndCookie 发送post请求 获取返回数据和cookie
func (r *ReqParam2)SendPostRequestToGetDataAndCookie() ([]byte, []*http.Cookie, error) {
    var data []byte
    var err error
    cookies := make([]*http.Cookie, 0)
    // 完整的请求地址
    r.Fulladdr, err = r.getFullURL(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, cookies, err
    }
    // 拼请求body参数
    r.setRequestBody2()
    // 生成请求
    req, err := r.genPost()
    if nil != err {
        log.Println(err)
        return data, cookies, err
    }
    // 发送请求
    resp, err := r.doSendWithReqToGetResp(req)
    if nil != err {
        log.Println(err)
        return data, cookies, err
    }
    defer resp.Body.Close()
    // 获取cookie
    cookies = resp.Cookies()
    // 获取返回数据
    data, err = ioutil.ReadAll(resp.Body)
    return data, cookies, err
}

// 生成post请求
func (r *ReqParam2)genPost() (*http.Request, error) {
    // 生成请求
    req, err := http.NewRequest("POST", r.Fulladdr, strings.NewReader(r.Body2))
    if nil != err {
        log.Println("生成请求失败")
        return nil, err
    }
    // 设置请求头 设置Content-Type
    req.Header.Set("Content-Type", r.ContentType)
    // 自定义的请求头
    for k, v := range r.Header {
        for _, vv := range v {
            req.Header.Set(k, vv)
        }
    }
    // 设置cookies
    for _, c := range r.Cookies {
        req.AddCookie(c)
    }
    // 设置user agent
    if "" != r.UserAgent {
        req.Header.Set("User-Agent", r.UserAgent)
    }
    return req, nil
}

// 发送post请求 p post请求body中的参数  apiurl 完整的请求地址 带自定义请求头
func (r *ReqParam2)doSendPostWithURL(p, contentType, apiurl string, header http.Header) ([]byte, error) {
    var rtn []byte
    // 指定不验证服务器端证书 生成客户端
    tr := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // 生成请求
    req, err := http.NewRequest("POST", apiurl, strings.NewReader(p))
    if nil != err {
        log.Println("生成请求失败")
        return rtn, err
    }
    // 设置请求头
    req.Header.Set("Content-Type", contentType)
    // 自定义的请求头
    for k, v := range header {
        for _, vv := range v {
            req.Header.Set(k, vv)
        }
    }
    // 发送post请求
    resp, err := client.Do(req)
    if nil != err {
        log.Println("发送请求失败")
        return rtn, err
    }
    defer resp.Body.Close()
    // 读取返回参数
    rtn, err = ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if nil != err {
        log.Println("返回数据读取失败")
        return rtn, err
    }
    return rtn, nil
}

// ============================
// 其他函数
// ============================

// 发送请求返回请求结果 req 拼装完成的http请求结构体
func (r *ReqParam2)doSendWithReq(req *http.Request) ([]byte, error) {
    var rtn []byte
    // 指定不验证服务器端证书 生成客户端
    tr := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // 发送post请求
    resp, err := client.Do(req)
    if nil != err {
        log.Println("发送请求失败")
        return rtn, err
    }
    defer resp.Body.Close()
    // 读取返回参数
    rtn, err = ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if nil != err {
        log.Println("返回数据读取失败")
        return rtn, err
    }
    return rtn, nil
}

// 发送post请求 获取http响应结构体 req 拼装完成的http请求结构体
func (r *ReqParam2)doSendWithReqToGetResp(req *http.Request) (*http.Response, error) {
    // 指定不验证服务器端证书 生成客户端
    tr := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // 发送post请求
    resp, err := client.Do(req)
    if nil != err {
        log.Println("发送请求失败")
        return nil, err
    }
    return resp, nil
}

// 设置请求body中的数据
func (r *ReqParam2)setRequestBody2() {
    if nil != r.Body && len(r.Body) > 0 {
        r.Body2 = r.Body.Encode()
    }
}

// 设置请求的cookie
func (r *ReqParam2)addCookie(cookie *http.Cookie) {
    r.Cookies = append(r.Cookies, cookie)
}

// 设置请求的cookie
func (r *ReqParam2)setCookie(key, value, path string) {
    cookie := new(http.Cookie)
    cookie.Name = key
    cookie.Value = value
    if "" != path {
        cookie.Path = path
    }
    r.Cookies = append(r.Cookies, cookie)
}

// getFullURL 生成带参数的完整链接
func (r *ReqParam2)getFullURL(addr string, params url.Values) (string, error) {
    // 解析链接地址
    u, err := url.Parse(addr)
    if nil != err {
        log.Printf("getFullURL 解析链接失败 链接：%s", addr)
        return "", err
    }
    // 添加参数
    q := u.Query()
    if nil != params && len(params) > 0 {
        for k, v := range params {
            q.Add(k, v[0])
        }
    }
    // 参数拼到链接地址上
    u.RawQuery = q.Encode()
    return u.String(), nil
}