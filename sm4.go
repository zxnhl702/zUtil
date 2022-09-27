// Package zUtil 数据加密解密的公共函数
// 国密4(SM4)算法加密
package zUtil

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"

	"github.com/tjfoc/gmsm/sm4"
)

// SM4加密
type SM4 struct {
    Key             []byte  // 密钥
    IV              []byte  // 偏移量
    Mode            int     // 加解密模式
}

// 新建 SM4加密对象
func NewSM4(key, iv string, mode int) *SM4 {
    return &SM4{Key : []byte(key), IV: []byte(iv), Mode: mode}
}

// SM4加密
func (s *SM4)SM4Encrypt(msg []byte) ([]byte, error) {
    // 判断加解密模式
    if CBC == s.Mode {
        // SM4加密 CBC模式
        return s.sm4EncryptCBC(msg)
    } else if ECB == s.Mode {
        // SM4加密 ECB模式
        return s.sm4EncryptECB(msg)
    } else if CFB == s.Mode {
        // SM4加密 CFB模式
        return s.sm4EncryptCFB(msg)
    } else if OFB == s.Mode {
        // SM4加密 OFB模式
        return s.sm4EncryptOFB(msg)
    } else {
        return []byte(""), errors.New("未知的加解密模式")
    }
}

// SM4解密
func (s *SM4)SM4Decrypt(msg []byte) ([]byte, error) {
    // 判断加解密模式
    if CBC == s.Mode {
        // SM4解密 CBC模式
        return s.sm4DecryptCBC(msg)
    } else if ECB == s.Mode {
        // SM4解密 ECB模式
        return s.sm4DecryptEBC(msg)
    } else if CFB == s.Mode {
        // SM4解密 CFB模式
        return s.sm4DecryptCFB(msg)
    } else if OFB == s.Mode {
        // SM4解密 OFB模式
        return s.sm4DecryptOFB(msg)
    } else {
        return []byte(""), errors.New("未知的加解密模式")
    }
}

// SM4加密 CBC模式 pkcs7
func (s *SM4)sm4EncryptCBC(msg []byte) ([]byte, error) {
    var err error
    var sm4Data []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return sm4Data, err
    }
    // 数据加密
    return sm4.Sm4Cbc(s.Key, msg, true)
}

// SM4加密 ECB模式 pkcs7
func (s *SM4)sm4EncryptECB(msg []byte) ([]byte, error) {
    var err error
    var sm4Data []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return sm4Data, err
    }
    // 数据加密
    return sm4.Sm4Ecb(s.Key, msg, true)
}

// SM4加密 CFB模式 pkcs7
func (s *SM4)sm4EncryptCFB(msg []byte) ([]byte, error) {
    var err error
    var sm4Data []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return sm4Data, err
    }
    // 数据加密
    return sm4.Sm4CFB(s.Key, msg, true)
}

// SM4加密 OFB模式 pkcs7
func (s *SM4)sm4EncryptOFB(msg []byte) ([]byte, error) {
    var err error
    var sm4Data []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return sm4Data, err
    }
    // 数据加密
    return sm4.Sm4OFB(s.Key, msg, true)
}

// SM4解密 ECB模式 pkcs7
func (s *SM4)sm4DecryptEBC(msg []byte) ([]byte, error) {
    var err error
    var decrypted []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return decrypted, err
    }
    // 数据解密
    return sm4.Sm4Ecb(s.Key, msg, false)
}

// SM4解密 CBC模式 pkcs7
func (s *SM4)sm4DecryptCBC(msg []byte) ([]byte, error) {
    var err error
    var decrypted []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return decrypted, err
    }
    // 数据解密
    return sm4.Sm4Cbc(s.Key, msg, false)
}

// SM4解密 CFB模式 pkcs7
func (s *SM4)sm4DecryptCFB(msg []byte) ([]byte, error) {
    var err error
    var decrypted []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return decrypted, err
    }
    // 数据解密
    return sm4.Sm4CFB(s.Key, msg, false)
}

// SM4解密 OFB模式 pkcs7
func (s *SM4)sm4DecryptOFB(msg []byte) ([]byte, error) {
    var err error
    var decrypted []byte
    // 设置偏移量
    err = sm4.SetIV(s.IV)
    if nil != err {
        log.Printf("SM4设置偏移量失败, 错误信息：%v", err)
        return decrypted, err
    }
    // 数据解密
    return sm4.Sm4OFB(s.Key, msg, false)
}

// SM4加密并转换成base64 pkcs7
func (s *SM4)SM4EncryptWithBase64(msg []byte) (string, error) {
    var err error
    // SM4加密
    sm4Data, err := s.SM4Encrypt(msg)
    if nil != err {
        log.Printf("SM4加密失败, 错误信息:%v", err)
        return "", err
    }
    // base64编码
    base64Data := base64.StdEncoding.EncodeToString(sm4Data)
    return base64Data, nil
}

// SM4加密并转换成16进制 pkcs7
func (s *SM4)SM4EncryptWithHex(msg []byte) (string, error) {
    var err error
    // SM4加密
    sm4Data, err := s.SM4Encrypt(msg)
    if nil != err {
        log.Printf("SM4加密失败, 错误信息:%v", err)
        return "", err
    }
    // hex
    hexData := hex.EncodeToString(sm4Data)
    return hexData, nil
}

// SM4解密 base64解码后 再解密
func (s *SM4)SM4DecryptWithBase64(msg []byte) (string, error) {
    var err error
    // 先base64解码
    base64Data, err := base64.StdEncoding.DecodeString(string(msg))
    if nil != err {
        log.Printf("base64解码失败, 错误信息:%v", err)
        return string(base64Data), err
    }
    // SM4解密
    decrypted, err := s.SM4Decrypt(base64Data)
    return string(decrypted), err
}

// SM4解密 16进制转换后 再解密
func (s *SM4)SM4DecryptWithHex(msg []byte) (string, error) {
    var err error
    // 先16进制转换
    hexData, err := hex.DecodeString(string(msg))
    if nil != err {
        log.Printf("16进制转换失败, 错误信息:%v", err)
        return string(hexData), err
    }
    // SM4解密
    decrypted, err := s.SM4Decrypt(hexData)
    return string(decrypted), err
}