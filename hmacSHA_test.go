package zUtil

import (
    "log"
    "testing"
)

func TestGetHMacSHA256(t *testing.T) {
    h := NewHMacSHA256(`{"token":"dd7d9f43-662e-44e2-bd16-3a0a8aa4efff"}`, "55ea54dac3f497c43344a9904f4aa1ae")
    log.Println(h.GetHMacSHA256())
    t.Log(h.GetHMacSHA256())
}

func TestGetHMacSHA256Base64Encoded(t *testing.T) {
    h := NewHMacSHA256(`{"token":"dd7d9f43-662e-44e2-bd16-3a0a8aa4efff"}`, "55ea54dac3f497c43344a9904f4aa1ae")
    log.Println(h.GetHMacSHA256Base64Encoded())
    t.Log(h.GetHMacSHA256Base64Encoded())
}