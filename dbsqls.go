// Package zUtil sql生成函数
package zUtil

import (
    "log"
    "regexp"
    "strings"
    "text/template"

    "github.com/flosch/pongo2"
)

// 请求参数数据全部转换成pongo2的数据类型
func (q QueryParam)ToPongoContext() *pongo2.Context {
    // 检索参数格式转换
    p2 := make(pongo2.Context)
    for k, v := range q {
        p2[k] = v
    }
    return &p2
}

// GetDynamicSQL 根据模板生成动态sql并去掉多余空格和换行
func GetDynamicSQL(p *pongo2.Context, sqltpl ...string) ([]string, error) {
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

// GetDynamicSQLWithGoTemplate 使用golang text/template模板库生成动态sql并去掉或与空格和换行 TODO Untest
func (q QueryParam)GetDynamicSQLWithGoTemplate(sqltpl ...string) ([]string, error) {
    var sqls []string
    // 多余空格和换行的正则表达式
    exp1 := regexp.MustCompile(`\s{2,}|\n`)
    for _, v := range sqltpl {
        tmpstr := &strings.Builder{}
        // 初始化模板
        t, err := template.New("tmp").Parse(v)
        if nil != err {
            log.Println(err)
            return sqls, err
        }
        // 参数写入模板
        err = t.Execute(tmpstr, q)
        if nil != err {
            log.Println(err)
            return sqls, err
        }
        // 去掉多余空格和换行 并写入数组中
        sqls = append(sqls, exp1.ReplaceAllString(tmpstr.String(), " "))
    }
    return sqls, nil
}