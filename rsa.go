package zUtil

import (
    "bytes"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "log"
)

// 读取RSA密钥文件
//   keypath 密钥文件的路径
func ReadRSAKeyFile(keypath string) (*pem.Block, string, error) {
    // 读取密钥文件
    keyfile, err := ioutil.ReadFile(keypath)
    if nil != err {
        log.Println(err)
        return nil, string(keyfile), err
    }
    // pem格式的密钥解码
    block, _ := pem.Decode(keyfile)
    if nil == block {
        return nil, string(keyfile), err
    }
    return block, string(keyfile), nil
}

// GetRSAPublicKey 生成RSA加密密钥
func GetRSAPublicKey(filepath string) (*rsa.PublicKey, string, error) {
    // 读取加密密钥
    block, keystring, err := ReadRSAKeyFile(filepath)
    if nil != err {
        log.Println(err)
        return nil, keystring, err
    }
    // 生成加密密钥
    publicinteface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if nil != err {
        log.Println(err)
        return nil, keystring, err
    }
    pub := publicinteface.(*rsa.PublicKey)
    return pub, keystring, nil
}

// GetRSAPrivateKeyPKCS8 生成RSA解密密钥 密钥文件PKCS8格式
func GetRSAPrivateKeyPKCS8(filepath string) (*rsa.PrivateKey, string, error) {
    // 读取加密密钥
    block, keystring, err := ReadRSAKeyFile(filepath)
    if nil != err {
        log.Println(err)
        return nil, keystring, err
    }
    // 生成解密密钥
    privinterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
    if nil != err {
        log.Println(err)
        return nil, keystring, err
    }
    priv := privinterface.(*rsa.PrivateKey)
    return priv, keystring, nil
}

// GetRSAPrivateKeyPKCS1 生成RSA解密密钥 密钥文件PKCS1格式
func GetRSAPrivateKeyPKCS1(filepath string) (*rsa.PrivateKey, string, error) {
    // 读取加密密钥
    block, keystring, err := ReadRSAKeyFile(filepath)
    if nil != err {
        log.Println(err)
        return nil, keystring, err
    }
    // 生成解密密钥
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if nil != err {
        log.Println(err)
        return nil, keystring, err
    }
    return priv, keystring, nil
}

// RSAEncrypt RSA加密 公钥文件路径作为参数
func RSAEncrypt(keypath string, msg []byte) ([]byte, error) {
    var encryptedMsg []byte
    // 读取公钥(加密密钥)
    block, _, err := ReadRSAKeyFile(keypath)
    if nil != err {
        log.Println(err)
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

// RSADecrypt RSA解密 私钥文件路径作为参数 PKCS8格式
func RSADecrypt(keypath string, msg []byte) ([]byte, error) {
    var decryptedMsg []byte
    // 读取私钥(解密密钥)
    block, _, err := ReadRSAKeyFile(keypath)
    if nil != err {
        log.Println(err)
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

// RsaEncrypt RSA加密 公钥对象作为参数
func RsaEncrypt(pub *rsa.PublicKey, msg []byte) ([]byte, error) {
    var encryptedMsg []byte
    var err error
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

// RsaDecrypt RSA解密 私钥对象作为参数
func RsaDecrypt(priv *rsa.PrivateKey, msg []byte) ([]byte, error) {
    var decryptedMsg []byte
    var err error
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

// RsaSign RSA私钥签名
func RsaSign(priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error) {
    return rsa.SignPKCS1v15(rand.Reader, priv, hash, hashed)
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