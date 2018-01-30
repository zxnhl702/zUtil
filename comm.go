// 为分类的其他使用的公共函数
package zUtil

import (
    "bytes"
    "crypto/md5"
    "database/sql"
    "encoding/hex"
    "encoding/json"
    "encoding/xml"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "os"
    "time"
)

// 标准返回数据的结构体
type Ret struct {
    Success bool        `json:"success" xml:"success"`
    ErrMsg  string      `json:"errMsg" xml:"errMsg"`
    Data    interface{} `json:"data" xml:"data"`
}

// 移动/重命名文件
func MoveFile(fileRoot, filename, newfilename string) error {
    oldPath := fileRoot + "/" + filename
    newPath := fileRoot + "/" + newfilename
    log.Printf("old: %s|||new:%s", oldPath, newPath)
    time.Now()
    return os.Rename(oldPath, newPath)
}

// 获取随机MD5字符串
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

// 获取GET的参数
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

// 打印并抛出异常
func Perr(e error, errMsg string, tx *sql.Tx) {
    // 有tx时 抛出异常时回滚事务
    if nil != tx {
        PerrorWithRollBack(e, errMsg, tx)
    } else {
        Perror(e, errMsg)
    }
}

// 打印并抛出异常
func Perror(e error, errMsg string) {
    if e != nil {
        log.Println(e)
        panic(errMsg)
    }
}

// 打印并抛出异常
func PerrorWithRollBack(e error, errMsg string, tx *sql.Tx) {
    if e != nil {
        tx.Rollback()
        log.Println(e)
        panic(errMsg)
    }
}

// 对象生成json字符串
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

// 设置返回的json数据
func GetJsonResult(r *http.Request, rt *Ret) []byte {
    bs, err := json.Marshal(rt)
    if err != nil {
        panic(err)
    }
    // 没有callback时 返回json字符串
    if "" == GetParameter(r, "callback") {
        return bs
    } else {
        // 有callback时  返回jsonp
        return []byte(GetParameter(r, "callback") + "(" + string(bs) + ")")
    }
}

// 设置返回的xml数据
func GenXMLResult(r *http.Request, rt *Ret) []byte {
    bs, err := xml.Marshal(rt)
    if nil != err {
        panic(err)
    }
    return bs
}

// uri 编码转换
func Uriencode(s string) string {
    return url.QueryEscape(s)
}

// uri 解码转换
func Uridecode(s string ) (string, error) {
    return url.QueryUnescape(s)
}

// 执行select语句并返回数据
func Fetch(sqlCmd string, valstr string, db *sql.DB) (string, interface{}) {
    vals := make([]interface{}, 20)
    json.Unmarshal([]byte(valstr), &vals)

    rows, err := db.Query(sqlCmd, vals...)
    Perror(err, "无法执行sql")
    columns, err := rows.Columns()
    Perror(err, "无法获取列信息")
    sqlLen := len(columns)

    var ret []interface{}
    for rows.Next() {
        vals := make([]sql.RawBytes, sqlLen)
        scanArgs := make([]interface{}, len(vals))
        for i := range vals {
            scanArgs[i] = &vals[i]
        }
        rows.Scan(scanArgs...)

        s_vals := make(map[string]string, 0)
        for i, col := range vals {
            s_vals[columns[i]] = string(col)
        }
        ret = append(ret, s_vals)
    }

    return "sql执行成功", ret
}

// 执行select语句并返回数据数组
func FetchWithArray(sqlCmd string, valstr string, db *sql.DB) (string, []interface{}) {
    vals := make([]interface{}, 20)
    json.Unmarshal([]byte(valstr), &vals)

    rows, err := db.Query(sqlCmd, vals...)
    Perror(err, "无法执行sql")
    columns, err := rows.Columns()
    Perror(err, "无法获取列信息")
    sqlLen := len(columns)

    var ret []interface{}
    for rows.Next() {
        vals := make([]sql.RawBytes, sqlLen)
        scanArgs := make([]interface{}, len(vals))
        for i := range vals {
            scanArgs[i] = &vals[i]
        }
        rows.Scan(scanArgs...)

        s_vals := make(map[string]string, 0)
        for i, col := range vals {
            s_vals[columns[i]] = string(col)
        }
        ret = append(ret, s_vals)
    }

    return "sql执行成功", ret
}

// 随机字符串的类型
const (
    RandomStringNumberOnly int = iota    // 只有数字
    RandomStringLowerOnly                       // 只有小写字母
    RandomStringUpperOnly                       // 只有大写字母
    RandomStringAll                             // 大小写字母+数字
)
// 产生随机字符串
func RandomString(size int, kind int) []byte {
    ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
    is_all := kind > 2 || kind < 0
    rand.Seed(time.Now().UnixNano())
    for i :=0; i < size; i++ {
        if is_all { // random ikind
            ikind = rand.Intn(3)
        }
        scope, base := kinds[ikind][0], kinds[ikind][1]
        result[i] = uint8(base+rand.Intn(scope))
    }
    return result
}