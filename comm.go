// Package zUtil 为分类的其他使用的公共函数
package zUtil

import (
    "bytes"
    "crypto/md5"
    crand "crypto/rand"
    "database/sql"
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
    "encoding/xml"
    "io"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "os"
    "time"
)

// Ret 标准返回数据的结构体
type Ret struct {
    Success bool        `json:"success" xml:"success"`
    ErrMsg  string      `json:"errMsg" xml:"errMsg"`
    Data    interface{} `json:"data" xml:"data"`
}

// MoveFile 移动/重命名文件
func MoveFile(fileRoot, filename, newfilename string) error {
    oldPath := fileRoot + "/" + filename
    newPath := fileRoot + "/" + newfilename
    log.Printf("old: %s|||new:%s", oldPath, newPath)
    time.Now()
    return os.Rename(oldPath, newPath)
}

// GetMd5String 获取随机MD5字符串
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

// GetRandomFilename 获取随机文件名
func GetRandomFilename() string {
    b := make([]byte, 48)
    if _, err := io.ReadFull(crand.Reader, b); err != nil {
        panic("创建文件名出错")
    }
    return time.Now().Format("20060102150405") + GetMd5String(base64.URLEncoding.EncodeToString(b))
}

// GetParameter 获取GET的参数
func GetParameter(r *http.Request, key string) string {
    // get请求 获取url中的参数
    if "GET" == r.Method {
        return r.URL.Query().Get(key)
    }
    // post请求 优先获取body中的参数
    if "POST" == r.Method  && "" != r.FormValue(key) {
        return r.FormValue(key)
    } else if "POST" == r.Method  && "" != r.URL.Query().Get(key) {
        // body中没有 再获取url中的参数
        return r.URL.Query().Get(key)
    }
    // 全都没有返回空字符串
    return ""
}

// Perr 打印并抛出异常
func Perr(e error, errMsg string, tx *sql.Tx) {
    // 有tx时 抛出异常时回滚事务
    if nil != tx {
        PerrorWithRollBack(e, errMsg, tx)
    } else {
        Perror(e, errMsg)
    }
}

// Perror 打印并抛出异常
func Perror(e error, errMsg string) {
    if e != nil {
        log.Println(e)
        panic(errMsg)
    }
}

// PerrorWithRollBack 打印并抛出异常
func PerrorWithRollBack(e error, errMsg string, tx *sql.Tx) {
    if e != nil {
        tx.Rollback()
        log.Println(e)
        panic(errMsg)
    }
}

// JSONMarshal 对象生成json字符串
// safeEncoding=true 将<>&3个符号从unicode替换成符号
// safeEncoding=false 不做替换
func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
    b, err := json.Marshal(v)

    if safeEncoding {
        b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
        b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
        b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
    }
    return b, err
}

// GetJSONResult 设置返回的json数据
func GetJSONResult(r *http.Request, rt *Ret) []byte {
    bs, err := json.Marshal(rt)
    if err != nil {
        panic(err)
    }
    // 没有callback时 返回json字符串
    if "" == GetParameter(r, "callback") {
        return bs
    }
    // 有callback时  返回jsonp
    return []byte(GetParameter(r, "callback") + "(" + string(bs) + ")")
}

// GenXMLResult 设置返回的xml数据
func GenXMLResult(r *http.Request, rt *Ret) []byte {
    bs, err := xml.Marshal(rt)
    if nil != err {
        panic(err)
    }
    return bs
}

// Uriencode uri 编码转换
func Uriencode(s string) string {
    return url.QueryEscape(s)
}

// Uridecode uri 解码转换
func Uridecode(s string ) (string, error) {
    return url.QueryUnescape(s)
}

// 随机字符串的类型
const (
    RandomStringNumberOnly int = iota    // 只有数字
    RandomStringLowerOnly                       // 只有小写字母
    RandomStringUpperOnly                       // 只有大写字母
    RandomStringAll                             // 大小写字母+数字
)
// RandomString 产生随机字符串
func RandomString(size int, kind int) []byte {
    ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
    isALL := kind > 2 || kind < 0
    rand.Seed(time.Now().UnixNano())
    for i :=0; i < size; i++ {
        if isALL { // random ikind
            ikind = rand.Intn(3)
        }
        scope, base := kinds[ikind][0], kinds[ikind][1]
        result[i] = uint8(base+rand.Intn(scope))
    }
    return result
}