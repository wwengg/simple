package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

/*
GenerateAESKey 生成指定字节长度的随机密钥
选择密钥长度：16 (128-bit), 24 (192-bit), 32 (256-bit)
*/
func GenerateAESKey(keySize int) (string, error) {
	key := make([]byte, keySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

/*
GenerateRandomString 生成随机字符串
length为字节长度，编码后长度会不同
*/
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// 使用URL安全的Base64编码（无填充）
	return base64.URLEncoding.EncodeToString(b), nil
}
