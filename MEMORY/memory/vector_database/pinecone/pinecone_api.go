package pinecone

import (
	"GoTuber/proxy"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

// QueryResp 请求返回
type QueryResp struct {
	Matches   []Matches `json:"matches"`
	Namespace string    `json:"namespace"`
}

type Matches struct {
	Id       string                 `json:"id"`
	Score    float64                `json:"score"`
	Metadata map[string]interface{} `json:"metadata"`
}

// StoreResp 储存返回
type StoreResp struct {
	UpsertedCount int64 `json:"upsertedCount"`
}

// QueryPayload Query请求
type QueryPayload struct {
	Filter          map[string]any `json:"filter"`
	IncludeValues   bool           `json:"includeValues"`
	IncludeMetadata bool           `json:"includeMetadata"`
	Vector          []float32      `json:"vector"`
	TopK            int            `json:"topK"`
	Namespace       string         `json:"namespace"`
}

// VectorData 存储请求
type VectorData struct {
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace"`
}

type Vector struct {
	ID       string         `json:"id"`
	Values   []float32      `json:"values"`
	Metadata map[string]any `json:"metadata"`
}

// PineconeQuery 请求相似数据
func PineconeQuery(filter, namespace string, vector []float32) []string {
	url := "https://" + pineconeCfg.Url + "/query"
	data := QueryPayload{
		Filter: map[string]any{
			"Type": filter,
		},
		IncludeValues:   false,
		IncludeMetadata: true,
		Vector:          vector,
		TopK:            1,
		Namespace:       namespace,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println("json格式化错误,", err)
		return nil
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", pineconeCfg.ApiKey)

	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var resp QueryResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("pinecone query错误:", err)
		return nil
	}
	if len(resp.Matches) == 0 {
		log.Println(string(body))
		return nil
	}
	return convertMapToStringSlice(resp.Matches[0].Metadata)
}

// PineconeStore 存储数据
func PineconeStore(typ, text, user, namespace string, vector []float32) {
	url := "https://" + pineconeCfg.Url + "/vectors/upsert"

	id := generateRandomString(10)
	if id == "" {
		return
	}
	data := VectorData{
		Vectors: []Vector{
			{
				ID:     id,
				Values: vector,
				Metadata: map[string]any{
					"Type": typ,
					"Text": text,
					"User": user,
				},
			},
		},
		Namespace: namespace,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println("json格式化错误,", err)
		return
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", pineconeCfg.ApiKey)

	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var resp StoreResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("pinecone upsert错误:", err)
	} else if resp.UpsertedCount < 1 {
		log.Println(string(body))
		log.Println("pinecone upsert错误:插入向量数量错误")
	}
}

// 随机生成10位的字符串用作id
func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println("生成vector id错误：", err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}

// 空接口转字符串切片
func convertMapToStringSlice(data map[string]interface{}) []string {
	result := make([]string, 0, len(data))
	for _, value := range data {
		strValue := fmt.Sprintf("%v", value)
		result = append(result, strValue)
	}
	return result
}
