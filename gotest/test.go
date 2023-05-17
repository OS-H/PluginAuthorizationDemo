package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	reqbody := strings.NewReader(`{"pluginid":"11111"}`)
	// 创建一个 GET 请求
	req, err := http.NewRequest("GET", "http://127.0.0.1:29090/authorization", reqbody)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	privateKeyString := "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAztIM292Rj2YiRSjCSbrOAyO0fW9w6gfw8eUyfA+nRYahKmhV\nBMhj5qw1n6GXIc3J6zkjk6FpKsYgMnFN824FexxSwpISXpNIfD7SESTP5Jq/isYP\noKGm/Q42rhQ9z+JKWsPQY5SjU3blaL74zIpuSDCeU2G9eZWz4AZHMwGnVOdbFGGm\nzJhWCGu1I2diiZUu9NV+ZEHE3nTNGyAgdVz5kxQVmRgxNF+GWueOFcW23kIGKRYv\nFJpSB5hpUV64slUJPzVv4O4J+TQoLCNq3nEX9PSNRRFAQxQRYe2DR9v2RJ03bQE9\nc6zPDrpGmqtLFXawzeVkYqecvYcxe32VKZ1bLwIDAQABAoIBACPf+6sHvAALz1X/\nw/PWG3Yf82butb9isUDEaQVsEa/Vso8Qme7Cc9HHfWW7OeP7NlM/DhTFouBwjZUy\nYjsfkoPQXeqyO8177s7edsHSiN02mpMP2BYc9EJg/MslZ7NvpUYpQTSEy+/mZ9TL\ni4yvVoHfLRd5lMxKU3FApYkLeGMZjrF6dlZ1NgqgPOuoxaXjGEhExHYR/5tSQQa8\nyx7rDvIFlOTvkixDMHld9LgW4WMVIKQP879BTsGca5Y8/G7ABaykr8G3u2tPXDIT\n3J/aks4COX8Yw+8mLYJrMIIfUCQ3LXTBswJxbCQufNDyVXs7xXF4lUwusgxBIeEz\nz6UiqcECgYEA4LooCSqa2b2EXT2ihJKb6H8ZRHqF7U2wvgYszYIvVOJn3kDkhXZ4\nzsWUB9tjrLBUS0ko87RpEnKrEApIRRqKXcOpcvf8JaTdWp4kLb4Fjnytp5jLet3b\nXlzoWRRx6YTNYNChs38OxVPIR2PCnbHz48QMNxqeLy3k0Y+tBbw9ZVsCgYEA65n5\nTfvOuV5mLrwGafmoyYlwRkehDSG/oyWS5jukO7nCODBJfVsOPpjI2SSu7sBtjlmE\nkPeWuGZgK6N+hVgYweWa13lCqm5mIJLcI0fkXGsDK4sSuBDciQe07hb0LXU+PE9J\nNHgH0h3KhUMttkVlmcGmlMNaW3qqKdJ3ARVoRb0CgYEAw36uHWtG0myfnU1k99di\ncds/a+b6YvnW6zgL+atq6XkbyqjBI6lwZtBSepNMHoo2ilfWnEsxrK68SXPoctUn\n0XHJEw7P9x94wMAZ0QEhbFbh6o5tVTFzCJ/iMLwsbGzvDW3xfWjmvJqp/BC42N5Z\nwKZnyfgJ7BkMmZFXf0nGT0kCgYAk7V1F+9HK/CDH8nCO67Ko5AHVAiUcCc4fpCQC\nMhbrxZHLfMYH/92bshbI8hb5FPAW/7Dnh+b3wBQSwu1xuP0oZvR+EWOBkwwuztXy\nMbJ5ScyVZpbogrwOPkb9ilt7RIUcrtCqiKWxKTo06PKhPv9NuiyB5Jyk+fTx2SsN\n4G0XgQKBgDabX6LFr7wBkkRsFClxYH+Bl3e5/lcxAMwkXTavFPfEjBDMQbjPkEDR\nDa5QWK55g3DtM6ACjeNDJvQrQODOIBUztx3735UqoesCCGennB9/HvyEW6GGOjmM\nJWqj+nT1SfUGxh7xXwYUu+DJw9Xr9ENn3iG/VC8imL3iKkypA4I+\n-----END RSA PRIVATE KEY-----\n"

	// 解码私钥字符串
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyString))
	if privateKeyBlock == nil {
		fmt.Println("Failed to decode private key")
		return
	}
	// 解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Failed to parse private key:" + err.Error())
		return
	}
	// 解密数据
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, body)
	if err != nil {
		fmt.Println("Failed to decrypt:" + err.Error())
	}
	fmt.Println(string(decryptedText))

}
