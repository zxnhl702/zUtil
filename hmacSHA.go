// Package zUtil 计算hmacsha256的函数
package zUtil

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
)

// HMacSHA256 的结构体
type HMacSHA256 struct {
    // 计算签名的消息
    Msg     []byte
    // 加密密钥
    Key     []byte
}

// 新建hamcsha256结构体
// 参数 
//   msg：需要计算的消息
//   key：加密密钥
func NewHMacSHA256(msg, key string) *HMacSHA256 {
    return &HMacSHA256{
        Msg: []byte(msg),
        Key: []byte(key),
    }
}

// 获取hmacsha256计算后的签名
func (v *HMacSHA256)getHMacSHA256() string {
    // 新建hmac哈希 使用sha256算法
    h := hmac.New(sha256.New, v.Key)
    // 需要计算的消息写入哈希中
    h.Write(v.Msg)
    // 计算前面 并转成字符串
    sha := hex.EncodeToString(h.Sum(nil))
    return sha
}

// 获取hmacsha256计算后的签名并用base64编码
func (v *HMacSHA256)getHMacSHA256Base64Encoded() string {
    // 先计算hmacsha256计算后的签名
    sha := v.getHMacSHA256()
    // base64编码
    return base64.StdEncoding.EncodeToString([]byte(sha))
}