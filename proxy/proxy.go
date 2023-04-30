package proxy

import (
	"log"
	"net/http"
	"net/url"
)

func Client() (http.Client, error) {
	if Cfg.proxy.UseProxy == false {
		return http.Client{}, nil
	}
	// 设置clash代理
	uri, err := url.Parse(Cfg.proxy.ProxyUrl)
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
