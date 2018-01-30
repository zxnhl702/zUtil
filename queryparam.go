// 接受数据的泛型结构以及数据类型转换
package zUtil

import (
    "log"
    "strconv"
)

// 查询参数
type QueryParam map[string]interface{}

// 请求参数转换成字符串
func (q QueryParam)toString(key string) string {
    return InterfaceToString(q[key])
}

// 请求参数转换成整数
func (q QueryParam)toInt(key string) int {
    return InterfaceToInt(q[key])
}

// 请求参数转换成浮点数
func (q QueryParam)toFloat64(key string) float64 {
    return InterfaceToFloat64(q[key])
}

// 请求参数转换成字符串数据
func (q QueryParam)toStringArray(key string) []string {
    return InterfaceToStringArray(q[key])
}

// 将泛型转换成字符串
func InterfaceToString(value interface{}) string {
    // 断言各种类型 进行格式转换成string
    switch v := value.(type) {
    case string:
        return v
    case int:
        return strconv.Itoa(v)
    case float32:
        return strconv.FormatFloat(float64(v), 'f', -1, 32)
    case float64:
        return strconv.FormatFloat(v, 'f', -1, 64)
    case bool:
        return strconv.FormatBool(v)
    default:
        return ""
    }
}

// 将泛型转换成整型 转换失败输出-99999
func InterfaceToInt(value interface{}) int {
    switch v := value.(type) {
    case string:
        log.Println("string", v)
        tmp, err := strconv.Atoi(v)
        if nil != err {
            return -99998
        } else {
            return tmp
        }
    case int:
        log.Println("int", v)
        return v
    case float32:
        log.Println("float32", v)
        return int(v)
    case float64:
        log.Println("float64", v)
        return int(v)
    case bool:
        if v {
            return 1
        } else {
            return 0
        }
    default:
        log.Println("default")
        return -99999
    }
}

// 将泛型转换成浮点型float64 转换失败输出-99999
func InterfaceToFloat64(value interface{}) float64 {
    switch v := value.(type) {
    case string:
        log.Println("string", v)
        tmp, err := strconv.ParseFloat(v, 64)
        if nil != err {
            return -99998.0
        } else {
            return tmp
        }
    case int:
        log.Println("int", v)
        return float64(v)
    case float32:
        log.Println("float32", v)
        return float64(v)
    case float64:
        log.Println("float64", v)
        return v
    case bool:
        if v {
            return 1
        } else {
            return 0
        }
    default:
        log.Println("default")
        return -99999
    }
}

// 泛型转换成字符串数组
func InterfaceToStringArray(value interface{}) []string {
    var arr []string
    switch v := value.(type) {
    case []string:
        return v
    case string:
        arr = make([]string, 1)
        arr[0] = v
        return arr
    case int:
        arr = make([]string, 1)
        arr[0] = strconv.Itoa(v)
        return arr
    case []int:
        for _, i := range v {
            arr = append(arr, strconv.Itoa(i))
        }
        return arr
    case float32:
        arr = make([]string, 1)
        arr[0] = strconv.FormatFloat(float64(v), 'f', -1, 32)
        return arr
    case []float32:
        for _, f := range v {
            arr = append(arr, strconv.FormatFloat(float64(f), 'f', -1, 32))
        }
        return arr
    case float64:
        arr = make([]string, 1)
        arr[0] = strconv.FormatFloat(v, 'f', -1, 64)
        return arr
    case []float64:
        for _, f := range v {
            arr = append(arr, strconv.FormatFloat(f, 'f', -1, 64))
        }
        return arr
    case bool:
        arr = make([]string, 1)
        arr[0] = strconv.FormatBool(v)
        return arr
    case []interface{}:
        for _, i := range v {
            arr = append(arr, InterfaceToString(i))
        }
        return arr
    default:
        log.Println("cannot convert to array")
        return arr
    }
}