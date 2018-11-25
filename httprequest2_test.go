package zUtil

import (
    "testing"
)

func TestSendGetRequestToGetDataAndCookie(t *testing.T) {
    r := NewReqParam2()
    r.Addr = "http://aaaa.bbb.gov.cn/forsbcxxxway.do"

    r.Param.Add("aac002", "3309XXXXXXXXXXXXX13")
    r.Param.Add("aga001", "14121YYY")
    r.Param.Add("aga001cy", "14121ZZZ")

    data, cookies, err := r.SendGetRequestToGetDataAndCookie()

    t.Log(err)
    t.Log(string(data))
    for _, c := range cookies {
        t.Log(c.Name, c.Value, c.Path, c.HttpOnly)
    } 
}