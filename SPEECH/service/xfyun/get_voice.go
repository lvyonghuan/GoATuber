package xfyun

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

// 鉴权url示例，取自讯飞官方文档
func assembleAuthUrl(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		log.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	//签名结果
	sha, err := HmacWithShaTobase64("hmac-sha256", []byte(sgin), apiSecret)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("api_key=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))
	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

// HmacWithShaTobase64 搓的
func HmacWithShaTobase64(key string, data []byte, secret string) (string, error) {
	sgin := []byte(key + string(data))
	apiSecret := []byte(secret)
	h := hmac.New(sha256.New, apiSecret)
	_, err := h.Write(sgin)
	if err != nil {
		return "", err
	}
	mac := h.Sum(nil)
	result := base64.StdEncoding.EncodeToString(mac)
	return result, nil
}
