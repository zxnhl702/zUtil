package zUtil

import (
    "bytes"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "log"
)

// RSAEncrypt RSA加密
func RSAEncrypt(keypath string, msg []byte) ([]byte, error) {
    var encryptedMsg []byte
    // 读取公钥(加密密钥)
    pk, err := ioutil.ReadFile(keypath)
    if nil != err {
        log.Println(err)
        return encryptedMsg, err
    }
    // pem格式的公钥解码
    block, _ := pem.Decode(pk)
    if nil == block {
        return encryptedMsg, err
    }
    // 生成加密密钥
    publicinteface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if nil != err {
        log.Println(err)
        return encryptedMsg, err
    }
    pub := publicinteface.(*rsa.PublicKey)
    // log.Println("key length", pub.N.BitLen())
    // 分段长度
    partLen := pub.N.BitLen() / 8 - 11
    // 消息分段
    chunks := splitMsg(msg, partLen)
    buffers := bytes.NewBufferString("")
    // 分段加密
    for _, c := range chunks {
        b, err := rsa.EncryptPKCS1v15(rand.Reader, pub, c)
        if nil != err {
            log.Println(err)
            return encryptedMsg, err
        }
        buffers.Write(b)
    }
    encryptedMsg = buffers.Bytes()
    return encryptedMsg, err
}

// RSADecrypt RSA解密
func RSADecrypt(keypath string, msg []byte) ([]byte, error) {
    var decryptedMsg []byte
    // 读取私钥(解密密钥)
    sk, err := ioutil.ReadFile(keypath)
    if nil != err {
        log.Println(err)
        return decryptedMsg, err
    }
    // pem格式的私钥解码
    block, _ := pem.Decode(sk)
    if nil == block {
        return decryptedMsg, err
    }
    // 生成解密密钥
    privinterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
    if nil != err {
        log.Println(err)
        return decryptedMsg, err
    }
    priv := privinterface.(*rsa.PrivateKey)
    // log.Println("key length", priv.N.BitLen())
    // 分段长度
    partLen := priv.N.BitLen() / 8
    // 消息分段
    chunks := splitMsg(msg, partLen)
    buffers := bytes.NewBufferString("")
    // 分段解密
    for _, c := range chunks {
        b, err := rsa.DecryptPKCS1v15(rand.Reader, priv, c)
        if nil != err {
            log.Println(err)
            return decryptedMsg, err
        }
        buffers.Write(b)
    }
    decryptedMsg = buffers.Bytes()
    return decryptedMsg, err
}

// 原文分段
func splitMsg(msg []byte, limit int) [][]byte {
    var chunk []byte
    chunks := make([][]byte, 0, len(msg)/limit +1)

    // 分段
    for len(msg) >= limit {
        chunk, msg = msg[:limit], msg[limit:]
        chunks = append(chunks, chunk)
    }
    if(len(msg) > 0) {
        chunks = append(chunks, msg[:len(msg)])
    }
    return chunks
}