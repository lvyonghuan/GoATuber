package proxy

import (
	"log"
	"net/http"
	"net/url"
)

func Client() (http.Client, error) {
	if Cfg.Proxy.UseProxy == false {
		return http.Client{}, nil
	}
	// 设置clash代理
	uri, err := url.Parse(Cfg.Proxy.ProxyUrl)
	if err != nil {
		log.Println(err)
		return http.Client{}, nil
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	return client, nil
}
