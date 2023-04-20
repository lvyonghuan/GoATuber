package proxy

import (
	"GoTuber/NLP/config"
	"log"
	"net/http"
	"net/url"
)

func Client() (http.Client, error) {
	if config.GPTCfg.Proxy.UseProxy == false {
		return http.Client{}, nil
	}
	// 设置clash代理
	uri, err := url.Parse(config.GPTCfg.Proxy.ProxyUrl)
	if err != nil {
		log.Fatal(err)
		return http.Client{}, nil
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	return client, nil
}
