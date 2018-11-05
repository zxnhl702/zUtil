// 发送请求的函数以及数据结构
package zUtil

import (
    "crypto/tls"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"
)

// http contenttype的常量定义
const (
    // x-www-form-urlencoded
    ContentTypeUrlencoded = "application/x-www-form-urlencoded;charset=utf-8"
    // multipart/form-data
    ContentTypeFormdata = "multipart/form-data"
    // application/json
    ContentTypeJson = "application/json;charset=utf-8"
    // text/xml
    ContentTypeXML = "text/xml;charset=utf-8"
)

// 请求结构体
type ReqParam struct {
    Addr        string          // 请求地址
    Fulladdr    string          // 完整请求地址
    Param       url.Values      // 请求参数(放在请求url上的)
    Body        url.Values      // 请求参数(放在请求body内的)
    Body2       string          // 请求参数(放在请求body内的 已经拼完的字符串)
    Cookies     []*http.Cookie  // cookie
    ContentType string          // contenttype
    Method      string          // 请求方法get/post
    Header      http.Header     // 请求头
}

// 新建请求结构体
func NewReqParam() *ReqParam {
    return &ReqParam{
        Addr        : "",
        Fulladdr    : "",
        Param       : make(url.Values),
        Body        : make(url.Values),
        Body2       : "",
        Cookies     : make([]*http.Cookie, 0),
        ContentType : ContentTypeUrlencoded,
        Method      : "",
        Header      : make(http.Header),
    }
}

// 发送get请求(非client模式 不带cookie)
func (r *ReqParam)SendGetRequest() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = GetFullUrl(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 发送请求
    data, err = doSendGet2(r.Fulladdr, r.Header)
    return data, err
}

// 发送get请求(https 不验证服务器端证书 client模式 不带cookie)
func (r *ReqParam)SendHttpsGetRequestSkipVerify() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = GetFullUrl(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 发送请求
    data, err = doSendGetHttpsSkipVerify(r.Fulladdr)
    return data, err
}

// 发送post请求(非client模式 不带cookie)
func (r *ReqParam)SendPostRequest() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = GetFullUrl(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 拼请求body参数
    r.setRequestBody2()
    // 发送请求
    data, err = doSendPost2(r.Body2, r.ContentType, r.Fulladdr, r.Header)
    return data, err
}

// 发送post请求(https 不验证服务器端证书 client模式 不带cookie)
func (r *ReqParam)SendHttpsPostRequestSkipVerify() ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    r.Fulladdr, err = GetFullUrl(r.Addr, r.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 拼请求body参数
    r.setRequestBody2()
    // 发送请求
    data, err = doSendPostHttpsSkipVerify(r.Body2, r.ContentType, r.Fulladdr)
    return data, err
}

// 设置请求body中的数据
func (r *ReqParam)setRequestBody2() {
    if nil != r.Body && len(r.Body) > 0 {
        r.Body2 = r.Body.Encode()
    }
}

// 设置请求的cookie
func (r *ReqParam)setCookie(key, value string) {
    cookie := new(http.Cookie)
    cookie.Name = key
    cookie.Value = value
    r.Cookies = append(r.Cookies, cookie)
}

// 发送get请求 apiurl 完整的请求地址
func doSendGet(apiurl string) ([]byte, error) {
    var rtn []byte
    // 发送get请求
    resp, err := http.Get(apiurl)
    if nil != err {
        log.Println("请求失败")
        return rtn, err
    }
    defer resp.Body.Close()
    // 读取返回数据
    rtn, err = ioutil.ReadAll(resp.Body)
    if nil != err {
        log.Println("返回数据读取失败")
        return rtn, err
    }
    return rtn, nil
}

// 发送post请求 p post请求body中的参数  apiurl 完整的请求地址 带自定义请求头
func doSendGet2(apiurl string, header http.Header) ([]byte, error) {
    var rtn []byte
    // 生成请求
    client := &http.Client{}
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

// 发送get请求https地址 不验证服务器端程序 
// apiurl 完整的请求地址
func doSendGetHttpsSkipVerify(apiurl string) ([]byte, error) {
    var rtn []byte
    // 指定不验证服务器端证书
    tr := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // 发送get请求
    resp, err := client.Get(apiurl)
    if nil != err {
        log.Println("请求失败")
        return rtn, err
    }
    defer resp.Body.Close()
    // 读取返回数据
    rtn, err = ioutil.ReadAll(resp.Body)
    if nil != err {
        log.Println("返回数据读取失败")
        return rtn, err
    }
    return rtn, nil
}

// 发送post请求 p post请求body中的参数  apiurl 完整的请求地址
func doSendPost(p, contentType, apiurl string) ([]byte, error) {
    var rtn []byte
    // 发送post请求
    resp, err := http.Post(apiurl, contentType, strings.NewReader(p))
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

// 发送post请求 p post请求body中的参数  apiurl 完整的请求地址 带自定义请求头
func doSendPost2(p, contentType, apiurl string, header http.Header) ([]byte, error) {
    var rtn []byte
    // 生成请求
    client := &http.Client{}
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

// 发送post请求https地址 不验证服务器端程序 
// p post请求body中的参数  apiurl 完整的请求地址
func doSendPostHttpsSkipVerify(p, contentType, apiurl string) ([]byte, error) {
    var rtn []byte
    // 指定不验证服务器端证书
    tr := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // 发送请求
    resp, err := client.Post(apiurl, contentType, strings.NewReader(p))
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

// 生成带参数的完整链接
func GetFullUrl(addr string, params url.Values) (string, error) {
    // 解析链接地址
    u, err := url.Parse(addr)
    if nil != err {
        log.Printf("getFullUrl 解析链接失败 链接：%s", addr)
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

// 生成带参数的完整链接
func GetFullUrl2(addr string, params map[string]string) (string, error) {
    // 解析链接地址
    u, err := url.Parse(addr)
    if nil != err {
        log.Printf("getFullUrl 解析链接失败 链接：%s", addr)
        return "", err
    }
    // 添加参数
    q := u.Query()
    if nil != params && len(params) > 0 {
        for k, v := range params {
            q.Add(k, v)
        }
    }
    // 参数拼到链接地址上
    u.RawQuery = q.Encode()
    return u.String(), nil
}