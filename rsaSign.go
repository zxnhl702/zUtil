// RSA的签名函数
package zUtil

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

// RSA签名对象
type RSASign struct {
    Priv        *rsa.PrivateKey     // RSA私钥对象
}

// 通过私钥对象新建RSA签名对象
func NewRSASignWithPrivateKey(priv *rsa.PrivateKey) *RSASign {
    return &RSASign{Priv : priv}
}

// 通过PKCS1格式私钥文件路径新建RSA签名对象
func NewRSASignWithPKCS1Filepath(pkcs1Filepath string) (*RSASign, error) {
    // 生成RSA解密密钥 密钥文件PKCS1格式
    priv, _, err := GetRSAPrivateKeyPKCS1(pkcs1Filepath)
    // 调用通过私钥对象生成RSA签名对象的函数
    return NewRSASignWithPrivateKey(priv), err
}

// 通过PKCS8格式私钥文件路径新建RSA签名对象
func NewRSASignWithPKCS8Filepath(pkcs8Filepath string) (*RSASign, error) {
    // 生成RSA解密密钥 密钥文件PKCS8格式
    priv, _, err := GetRSAPrivateKeyPKCS8(pkcs8Filepath)
    // 调用通过私钥对象生成RSA签名对象的函数
    return NewRSASignWithPrivateKey(priv), err
}

// SignMD5withRSA MD5摘要 RSA签名 签名结果进行base64编码
//  msg 待签名的原始数据
func (a *RSASign)SignMD5withRSA(msg []byte) (string, error) {
    var signBase64 string
    // 计算MD5的hash值
    hashData := md5.Sum(msg)
    // 使用RSA的私钥签名
    signBytes, err := rsa.SignPKCS1v15(rand.Reader, a.Priv, crypto.MD5, hashData[ : ])
    if nil != err {
        log.Println(err)
        return signBase64, err
    }
    // base64编码
    signBase64 = base64.URLEncoding.EncodeToString(signBytes)
    return signBase64, nil
}

// SignSHA1withRSA SHA1摘要 RSA签名 签名结果进行base64编码
//  msg 待签名的原始数据
func (a *RSASign)SignSHA1withRSA(msg []byte) (string, error) {
    var signBase64 string
    // 计算SHA1的hash值
    hashData := sha1.Sum(msg)
    // 使用RSA的私钥签名
    signBytes, err := rsa.SignPKCS1v15(rand.Reader, a.Priv, crypto.SHA1, hashData[ : ])
    if nil != err {
        log.Println(err)
        return signBase64, err
    }
    // base64编码
    signBase64 = base64.URLEncoding.EncodeToString(signBytes)
    return signBase64, nil
}

// SignSHA256withRSA SHA256摘要 RSA签名 签名结果进行base64编码
//  msg 待签名的原始数据
func (a *RSASign)SignSHA256withRSA(msg []byte) (string, error) {
    var signBase64 string
    // 计算SHA256的hash值
    hashData := sha256.Sum256(msg)
    // 使用RSA的私钥签名
    signBytes, err := rsa.SignPKCS1v15(rand.Reader, a.Priv, crypto.SHA256, hashData[ : ])
    if nil != err {
        log.Println(err)
        return signBase64, err
    }
    // base64编码
    signBase64 = base64.URLEncoding.EncodeToString(signBytes)
    return signBase64, nil
}

// RSA验签对象
type RSAVerifySign struct {
    Pub         *rsa.PublicKey      // RSA公钥对象
}

// 通过私钥对象新建RSA验签对象
func NewRSAVerifySignWithPublishKey(pk *rsa.PublicKey) *RSAVerifySign {
    return &RSAVerifySign{Pub : pk}
}

// 通过PKCS1格式私钥文件路径新建RSA验签对象
func NewRSAVerifySignWithPKFilepath(pkFilePath string) (*RSAVerifySign, error) {
    // 生成RSA公钥
    pk, _, err := GetRSAPublicKey(pkFilePath)
    // 调用通过私钥对象生成RSA签名对象的函数
    return NewRSAVerifySignWithPublishKey(pk), err
}

// VerifySignMD5withRSA 使用RSA公钥验签 MD5摘要
// msg 待签名的原始数据
// sign 签名
func (a *RSAVerifySign)VerifySignMD5withRSA(msg, sign []byte) error {
    // 计算消息hash值 MD5算法
    hashData := md5.Sum(msg)
    // 使用RSA公钥验签
    err := rsa.VerifyPKCS1v15(a.Pub, crypto.MD5, hashData[ : ], sign)
    if nil != err {
        log.Printf("验签失败, 错误信息:%v", err)
    }
    return err
}

// VerifySignMD5withRSA 使用RSA公钥验签 SHA1摘要
// msg 待签名的原始数据
// sign 签名
func (a *RSAVerifySign)VerifySignSHA1withRSA(msg, sign []byte) error {
    // 计算消息hash值 MD5算法
    hashData := sha1.Sum(msg)
    // 使用RSA公钥验签
    err := rsa.VerifyPKCS1v15(a.Pub, crypto.SHA1, hashData[ : ], sign)
    if nil != err {
        log.Printf("验签失败, 错误信息:%v", err)
    }
    return err
}

// VerifySignMD5withRSA 使用RSA公钥验签 SHA256摘要
// msg 待签名的原始数据
// sign 签名
func (a *RSAVerifySign)VerifySignSHA256withRSA(msg, sign []byte) error {
    // 计算消息hash值 MD5算法
    hashData := md5.Sum(msg)
    // 使用RSA公钥验签
    err := rsa.VerifyPKCS1v15(a.Pub, crypto.SHA256, hashData[ : ], sign)
    if nil != err {
        log.Printf("验签失败, 错误信息:%v", err)
    }
    return err
}