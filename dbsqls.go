// sql生成函数
package zUtil

import (
    "log"
    "regexp"

    "github.com/flosch/pongo2"
)

// 请求参数数据全部转换成pongo2的数据类型
func (q QueryParam)toPongoContext() *pongo2.Context {
    // 检索参数格式转换
    p2 := make(pongo2.Context)
    for k, v := range q {
        p2[k] = v
    }
    return &p2
}

// 根据模板生成动态sql并去掉多余空格和换行
func GetDynamicSql(p *pongo2.Context, sqltpl ...string) ([]string, error) {
    var sqls []string
    // 多余空格和换行的正则表达式
    exp1 := regexp.MustCompile(`\s{2,}|\n`)
    for _, v := range sqltpl {
        // 初始化模板
        t, err := pongo2.FromString(v)
        if nil != err {
            log.Println(err)
            return sqls, err
        }
        // 参数写入模板
        s, err := t.Execute(*p)
        if nil != err {
            log.Println(err)
            return sqls, err
        }
        // 去掉多余空格和换行 并写入数组中
        sqls = append(sqls, exp1.ReplaceAllString(s, " "))
    }
    return sqls, nil
}