package zUtil

import (
	"testing"
)

// 测试SM4加密 ECB模式 base64编码
func TestSM4EncryptECBBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", ECB)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 ECB模式 base64编码
func TestSM4DecryptECBBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", ECB)
    msg := "ErhUB1LWY0NZWU6ruNyz2juh5ZU8FoxWY2gCK1skmHA="
    data, _ := s.SM4DecryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4加密 ECB模式 16进制编码
func TestSM4EncryptEBCHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", ECB)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithHex([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 ECB模式 16进制编码
func TestSM4DecryptEBCHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", ECB)
    msg := "12b8540752d6634359594eabb8dcb3da3ba1e5953c168c566368022b5b249870"
    data, _ := s.SM4DecryptWithHex([]byte(msg))
    t.Log(data)
}

// 测试SM4加密 CBC模式 base64编码
func TestSM4EncryptCBCBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CBC)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 CBC模式 base64编码
func TestSM4DecryptCBCBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CBC)
    msg := "E5fjQfTTpDUte0XulTM6t0/qxUKR6TYy9/Jhl+R+/Y8="
    data, _ := s.SM4DecryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4加密 CBC模式 16进制编码
func TestSM4EncryptCBCHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CBC)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithHex([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 CBC模式 16进制编码
func TestSM4DecryptCBCHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CBC)
    msg := "1397e341f4d3a4352d7b45ee95333ab74feac54291e93632f7f26197e47efd8f"
    data, _ := s.SM4DecryptWithHex([]byte(msg))
    t.Log(data)
}

// 测试SM4加密 CFB模式 base64编码
func TestSM4EncryptCFBBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CFB)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 CFB模式 base64编码
func TestSM4DecryptCFBBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CFB)
    msg := "9QLt5lucN+9c0kZ3g8hiGlOgehhf5MOdMGVYJ4IyGww="
    data, _ := s.SM4DecryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4加密 CFB模式 16进制编码
func TestSM4EncryptCFBHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CFB)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithHex([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 CFB模式 16进制编码
func TestSM4DecryptCFBHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", CFB)
    msg := "f502ede65b9c37ef5cd2467783c8621a53a07a185fe4c39d3065582782321b0c"
    data, _ := s.SM4DecryptWithHex([]byte(msg))
    t.Log(data)
}

// 测试SM4加密 OFB模式 base64编码
func TestSM4EncryptOFBBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", OFB)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 OFB模式 base64编码
func TestSM4DecryptOFBBase64(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", OFB)
    msg := "9QLt5lucN+9c0kZ3g8hiGvDNbAww5Zw44WTHvASCdxs="
    data, _ := s.SM4DecryptWithBase64([]byte(msg))
    t.Log(data)
}

// 测试SM4加密 OFB模式 16进制编码
func TestSM4EncryptOFBHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", OFB)
    msg := "Hello World 123123 !@#"
    data, _ := s.SM4EncryptWithHex([]byte(msg))
    t.Log(data)
}

// 测试SM4解密 OFB模式 16进制编码
func TestSM4DecryptOFBHex(t *testing.T) {
    s := NewSM4("zsdy2020zsgd2020", "zsdy2020zsgd2020", OFB)
    msg := "f502ede65b9c37ef5cd2467783c8621af0cd6c0c30e59c38e164c7bc0482771b"
    data, _ := s.SM4DecryptWithHex([]byte(msg))
    t.Log(data)
}