// Package zUtil 微信支付签名
package zUtil

import (
    "crypto/md5"
    "encoding/hex"
    "reflect"
    "sort"
    "strings"
    "log"
)

// Params map形式的请求参数
type Params map[string]string

// 结构体转换成map
func struct2Map(data interface{}) Params {
    // 反射得到结构体中的数据类型
    t := reflect.TypeOf(data)
    // 反射得到结构体中的值
    v := reflect.ValueOf(data)
    params := make(Params)
    // 遍历结构体中的元素
    for i:=0; i<t.NumField(); i++ {
        // 获取结构体中xml tag的设置
        keys := strings.Split(t.Field(i).Tag.Get("xml"), ",")
        // 获取结构体元素的值
        values := InterfaceToString(v.Field(i).Interface())
        // 如果tag中含有omitempty并且没有值 或者 xml.Name类型的 或者 元素名称是sign的 舍弃
        if (tagWithOmitempty(keys) && "" == values) || t.Field(i).Type.Name() == "Name" || 
            "" == values || "sign" == keys[0] {
            continue
        } else {
            // key填写的时候
            if "" != keys[0] {
                // 将tag作为key 反射得到的值作为值 塞入map中
                params[keys[0]] = values
            } else {
                // key没有填写的时候 反射获取元素名作为key 
                params[t.Field(i).Name] = values
            }
        }
        
    }
    return params
}

// tag信息中是否含有omitempty
func tagWithOmitempty(keys []string) bool {
    for _, v := range keys {
        if v == "omitempty" {
            return true
        }
    }
    return false
}

// GetMD5Sign 计算签名
func (param Params)GetMD5Sign(paykey string) string {
    // 获取算法所需的stringA
    stringA := param.SortAndConcat()
    // 拼商户api密钥到字符串
    stringSignTemp := stringA + "&key=" + paykey
    log.Println(stringSignTemp)
    // MD5签名
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(stringSignTemp))
    return strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil)))
}

// GetMD5SignWithoutPaykey 计算签名 不带微信商家的peykey
func (param Params)GetMD5SignWithoutPaykey() string {
    // 获取算法所需的stringA
    stringA := param.SortAndConcat()
    // MD5签名
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(stringA))
    return strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil)))
}

// SortAndConcat 按字母顺序排序参数并拼接为key1=value1&key2&value2的形式
func (param Params)SortAndConcat() string {
    // 获取全部参数名称
    var keys []string
    for k := range param {
        keys = append(keys, k)
    }
    // 参数名称排序
    var sortedParam []string
    sort.Strings(keys)
    // 连接参数与参数值
    for _, k := range keys {
        sortedParam = append(sortedParam, k+"="+param[k])
    }
    // 拼接全部参数
    return strings.Join(sortedParam, "&")
}