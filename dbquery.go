package zUtil

import (
    "database/sql"
    "encoding/json"
    "log"
    "reflect"
    "strings"
    "strconv"
    "time"
)

// 获取单行数据
func GetSingleRowData(sqls string, db *sql.DB) (*map[string]interface{}, error) {
    // 调用获取多行数据的函数
    list, err := GetMultyRowsData(sqls, db)
    if nil != err {
        log.Println(err)
        return nil, err
    }
    // 获取的数据超过1行报错 否则返回获取的数据
    if len(list) > 1 {
        return nil, sql.ErrNoRows
    } else {
        return list[0], nil
    }
}

// 获取多行数据
func GetMultyRowsData(sqls string, db *sql.DB) ([]*map[string]interface{}, error) {
    // 返回数据
    list := make([]*map[string]interface{}, 0)
    // 执行sql
    rows, err := db.Query(sqls)
    if nil != err {
        log.Println(err)
        return list, err
    }
    defer rows.Close()
    
    // 每一列的列名
    columns, _ := rows.Columns()
    // 字段数量
    size := len(columns)

    for rows.Next() {
        // 行数据的指针 用来调用scan
        pts := make([]interface{}, size)
        // 行数据 用来存数据
        container := make([]interface{}, size)
        for i := range pts {
            pts[i] = &container[i]
        }
        // 扫描行数据
        rows.Scan(pts...)

        // 行数据 map形式存储
        rowData := make(map[string]interface{})
        // 循环取出来的每一列数据 根据不同的数据类型进行转换 然后存入map中
        for i, v := range container {
            log.Println(i, reflect.TypeOf(v).Name(), reflect.TypeOf(v).Kind().String())
            // 处理列名 确定map中的key
            col := getColumnName(i, columns[i])
            // 根据每个字段数据类型的不同 转换成string
            switch vv := v.(type) {
            case int64:
                // 整型 直接放到map中
                log.Println("int64",  columns[i], col, vv)
                rowData[col] = vv
            case float64:
                // 浮点型 直接放到map中
                log.Println("float64",  columns[i], col, vv)
                rowData[col] = vv
            case bool:
                // 布尔型 直接放到map中
                log.Println("float64",  columns[i], col, vv)
                rowData[col] = vv
            case time.Time:
                // 时间类型 转换成string 放到map中
                log.Println("time", columns[i], col, vv.Format("2006-01-02 15:04:05"))
                rowData[col] = vv.Format("2006-01-02 15:04:05")
            default:
                // 默认 先断言[]byte然后转换成string 放倒map中
                vvv, _ :=  v.([]byte)
                log.Println("default",  columns[i], col, string(vvv))
                rowData[col] = string(vvv)
            }
        }
        list = append(list, &rowData)
    }
    return list, nil
}

// 执行sql insert update delete
func ExecSql(sqls string, db *sql.DB) (int64, error) {
    var rowid int64
    var err error
    result, err := db.Exec(sqls)
    if nil != err {
        log.Println(err)
        return rowid, err
    }
    if strings.Index(sqls, "insert") >= 0 {
        // insert
        rowid, err = result.LastInsertId()
    } else { 
        // update or detete
        rowid, err = result.RowsAffected()
    }
    return rowid, err
}

// 处理列名
func getColumnName(index int, colname string) string {
    if strings.Index(colname, "(") == -1 {
        return colname
    } else {
        return "col" + strconv.Itoa(index+1)
    }
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

func Exec(sqlCmd string, db *sql.DB) (string, interface{}) {
    res, err := db.Exec(sqlCmd)
    Perror(err, "无法执行写操作")
    var ret int64
    if strings.Index(sqlCmd, "insert") >= 0 {
        ret, _ = res.LastInsertId()
    } else { // update or detete
        ret, _ = res.RowsAffected()
    }
    return "写操作成功", ret
}