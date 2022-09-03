// Package zUtil 数据加密解密的公共函数
// AES算法加密
package zUtil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"log"
)

// AES加密
type AES struct {
    Key             []byte  // 密钥
    IV              []byte  // 偏移量
}

// 新建 AES加密对象
func NewAES(key, iv string) *AES {
    return &AES{Key : []byte(key), IV: []byte(iv)}
}

// 填充数据
func padding(src []byte, blockSize int) []byte {
    padNum := blockSize - len(src) % blockSize
    pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
    return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
    n := len(src)
    unPadNum := int(src[n-1])
    return src[:n-unPadNum]
}

// AES加密 CBC模式
func (r *AES)AESEncryptCBC(src []byte) ([]byte, error) {
    block, err := aes.NewCipher(r.Key)
    if err != nil {
        return nil, err
    }
    src = padding(src, block.BlockSize())
    blockMode := cipher.NewCBCEncrypter(block, r.IV)
    blockMode.CryptBlocks(src, src)
    return src, nil
}

// AES解密 CBC模式
func (r *AES)AESDecryptCBC(src []byte) ([]byte, error) {
    block, err := aes.NewCipher(r.Key)
    if err != nil {
        return nil, err
    }
    blockMode := cipher.NewCBCDecrypter(block, r.IV)
    blockMode.CryptBlocks(src, src)
    src = unpadding(src)
    return src, nil
}

// AES加密并且base64编码 CBC模式
func (r *AES)AESEncryptCBCWithBase64(src []byte) (string, error) {
    var base64Data string
    // 先AES加密
    aesData, err := r.AESEncryptCBC(src)
    if nil != err {
        log.Println(err)
        return base64Data, err
    }
    // 再base64编码
    base64Data = base64.StdEncoding.EncodeToString(aesData)
    return base64Data, nil
}

// AES解密并且base64解码 CBC模式
func (r *AES)AESDecryptCBCWithBase64(src []byte) (string, error) {
    // 先base64解码
    base64Data, err := base64.StdEncoding.DecodeString(string(src))
    if nil != err {
        log.Println(err)
        return string(base64Data), err
    }
    // 再AES解密
    aesData, err := r.AESDecryptCBC(base64Data)
    if nil != err {
        log.Println(err)
        return string(aesData), err
    }
    return string(aesData), nil
}

// AES加密并且返回16进制字符串 CBC模式
func (r *AES)AESEncryptCBCWithHex(src []byte) (string, error) {
    var hexData string
    // 先AES加密
    aesData, err := r.AESEncryptCBC(src)
    if nil != err {
        log.Println(err)
        return hexData, err
    }
    // 再base64编码
    hexData = hex.EncodeToString(aesData)
    return hexData, nil
}

// AES解密 16进制字符串 CBC模式
func (r *AES)AESDecryptCBCWithHex(src []byte) (string, error) {
    // 先16进制字符串转回字节数组
    hexData, err := hex.DecodeString(string(src))
    if nil != err {
        log.Println(err)
        return string(hexData), err
    }
    // 再AES解密
    aesData, err := r.AESDecryptCBC(hexData)
    if nil != err {
        log.Println(err)
        return string(aesData), err
    }
    return string(aesData), nil
}