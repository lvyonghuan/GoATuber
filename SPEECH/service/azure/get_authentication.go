package azure

import (
	"GoTuber/SPEECH/config"
	"io"
	"log"
	"net/http"
)

var authentication string //authentication,9分钟更新一次，下面是获取的函数,九分钟重复一次

func GetAuthentication() {
	for {
		fetchTokenURL := config.AzureCfg.Azure.EndPoint
		req, _ := http.NewRequest("POST", fetchTokenURL, nil)
		req.Header.Set("Ocp-Apim-Subscription-Key", config.AzureCfg.Azure.Key)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()

		bodyBytes, _ := io.ReadAll(resp.Body)
		accessToken := string(bodyBytes)
		authentication = "Bearer " + accessToken
	}
}
