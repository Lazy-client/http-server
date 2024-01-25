package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
)

func main() {
	// 输入要发送的消息
	message := "Hello, World!"
	// 输入密钥
	secret := "secret"
	// 输入哈希算法
	algorithm := "SHA1"
	// 根据哈希算法选择相应的哈希函数
	var h func() hash.Hash
	switch algorithm {
	case "MD5":
		h = md5.New
	case "SHA1":
		h = sha1.New
	case "SHA256":
		h = sha256.New
	default:
		panic("Unsupported hash algorithm")
	}

	// 生成HMAC认证码
	mac := hmac.New(h, []byte(secret))
	mac.Write([]byte(message))
	signature := hex.EncodeToString(mac.Sum(nil))

	// 发送HTTP请求并带上HMAC认证码
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "HMAC "+signature)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 输出响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
