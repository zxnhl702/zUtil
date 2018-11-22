// 发送mutipart/form-data格式请求和数据结构
package zUtil

import (
    "bytes"
    "crypto/tls"
    "io"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
    "net/url"
    "os"
)

// mutipart/form-data请求结构体
type MutiPartParam struct {
    Addr        string          // 请求地址
    Fulladdr    string          // 完整请求地址
    Param       url.Values      // 请求参数(放在请求url上的)
    Cookies     []*http.Cookie  // cookie
    ContentType string          // contenttype
    body        *bytes.Buffer   // mutipart/form-data中的请求数据 FileData+FormData
    FileData    []*Files        // mutipart/form-data中需要上传的文件
    FormData    url.Values      // mutipart/form-data中其他需要添加的请求数据
}

// 通过mutipart/form-data上传的文件
type Files struct {
    Key         string          // mutipart/form-data 中的name
    Filename    string          // mutipart/form-data 中的filename
    Filepath    string          // 文件的绝对路径 用于读取文件内容
    Data        *os.File        // 文件内容
}

// 新建请求结构体
func NewMutiPartParam() *MutiPartParam {
    return &MutiPartParam {
        Addr        : "",
        Fulladdr    : "",
        Param       : make(url.Values),
        Cookies     : make([]*http.Cookie, 0),
        ContentType : "",
        body        : new(bytes.Buffer),
        FileData    : make([]*Files, 0),
        FormData    : make(url.Values),
    }
}

// 读取全部文件的内容来生成mutipart的数据
func (m *MutiPartParam) readAllFile() error {
    var err error
    // 遍历全部文件
    for _, f := range m.FileData {
        // 读取文件
        f.Data, err = os.Open(f.Filepath)
        if nil != err {
            log.Println("读取文件失败")
            log.Println(err)
            return err
        }
    }
    return nil
}

// 读取全部文件和普通参数到请求body中
func  (m *MutiPartParam) setRequestData() error {
    // 初始化mutipart/form-data的body
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    // 设置文件到body
    for _, f := range m.FileData {
        // 创建form-data中的formfile
        part, err := writer.CreateFormFile(f.Key, f.Filename)
        if nil != err {
            log.Println("创建formfile失败")
            log.Println(err)
            return err
        }
        // 文件数据复制到formfile中
        io.Copy(part, f.Data)
    }
    // 设置普通参数到body中
    for k, v := range m.FormData {
        for _, vv := range v {
            writer.WriteField(k, vv)
        }
    }
    // 设置content type
    m.ContentType = writer.FormDataContentType() + "; " + m.ContentType
    // 给body增加eof
    err := writer.Close()
    if nil != err {
        log.Println("writer关闭失败")
        log.Println(err)
        return err
    }
    m.body = body
    return nil
}

// 发送请求
func (m *MutiPartParam) Send(ishttps bool) ([]byte, error) {
    var data []byte
    var err error
    // 完整的请求地址
    m.Fulladdr, err = GetFullUrl(m.Addr, m.Param)
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 读取文件
    if len(m.FileData) > 0 {
        err = m.readAllFile()
        if nil != err {
            log.Println(err)
            return data, err
        }
    }
    // 设置请求数据
    err = m.setRequestData()
    if nil != err {
        log.Println(err)
        return data, err
    }
    // 发送请求
    if !ishttps {
        data, err = m.doSendPost()
    } else {
        data, err = m.doSendPostHttpsSkipVerify()
    }
    return data, err
}

// 发送post请求
func (m *MutiPartParam) doSendPost() ([]byte, error) {
    var data []byte
    var err error
    // 发送post请求
    resp, err := http.Post(m.Fulladdr, m.ContentType, m.body)
    if nil != err {
        log.Println("发送请求失败")
        return data, err
    }
    defer resp.Body.Close()
    // 读取返回参数
    data, err = ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if nil != err {
        log.Println("返回数据读取失败")
        return data, err
    }
    return data, nil
}

// 发送post请求https地址 不验证服务器端程序 
// apiurl 完整的请求地址
func (m *MutiPartParam) doSendPostHttpsSkipVerify() ([]byte, error) {
    var data []byte
    var err error
    // 指定不验证服务器端证书
    tr := &http.Transport{
        TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // 发送请求
    resp, err := client.Post(m.Fulladdr, m.ContentType, m.body)
    if nil != err {
        log.Println("发送请求失败")
        return data, err
    }
    defer resp.Body.Close()
    // 读取返回参数
    data, err = ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if nil != err {
        log.Println("返回数据读取失败")
        return data, err
    }
    return data, nil
}