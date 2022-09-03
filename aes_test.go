// Package zUtil 数据加密解密的公共函数
// AES的测试文件
package zUtil

import (
	"log"
	"testing"
)

// 测试AES加密 cbc模式 base64编码
func TestAESEncryptCBCBase64(t *testing.T) {
    a := NewAES("zsdy2020zsgd2020", "zsdy2020zsgd2020")
    msg := "Hello World 123123 !@#"
    encryptedMsg, _ := a.AESEncryptCBCWithBase64([]byte(msg))
    log.Println(encryptedMsg)
    t.Log(encryptedMsg)
}

// 测试AES解密 cbc模式 base64编码
func TestAESDecryptCBCBase64(t *testing.T) {
    a := NewAES("zsdy2020zsgd2020", "zsdy2020zsgd2020")
    data := "H1WCYadzoggE8xRtYxjG81eaSrnGeinZOnhmFCDc0go="
    decrpytedMsg, _ := a.AESDecryptCBCWithBase64([]byte(data))
    log.Println(decrpytedMsg)
    t.Log(decrpytedMsg)
}

// 测试AES加密 cbc模式 16进制字符串
func TestAESEncryptCBCHex(t *testing.T) {
    a := NewAES("zsdy2020zsgd2020", "zsdy2020zsgd2020")
    msg := "Hello World 123123 !@#"
    encryptedMsg, _ := a.AESEncryptCBCWithHex([]byte(msg))
    log.Println(encryptedMsg)
    t.Log(encryptedMsg)
}

// 测试AES解密 cbc模式 16进制字符串
func TestAESDecryptCBCHex(t *testing.T) {
    a := NewAES("zsdy2020zsgd2020", "zsdy2020zsgd2020")
    data := "1f558261a773a20804f3146d6318c6f3579a4ab9c67a29d93a78661420dcd20a"
    decrpytedMsg, _ := a.AESDecryptCBCWithHex([]byte(data))
    log.Println(decrpytedMsg)
    t.Log(decrpytedMsg)
}