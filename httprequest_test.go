package zUtil

import (
    "testing"
)

func TestNewReqParam(t *testing.T) {
    p := NewReqParam()
    p.Body.Add("tset1", "test2")
    t.Log(p.Body.Get("tset1"))
}