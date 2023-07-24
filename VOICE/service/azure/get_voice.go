package azure

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/VOICE/config"
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetVoice(msg *sensitive.OutPut) {
	getMood(msg.Mood)
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.AzureCfg.Azure.EndPointForVoice, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("X-Microsoft-OutputFormat", "riff-24khz-16bit-mono-pcm")
	req.Header.Add("Content-Type", "application/ssml+xml")
	req.Header.Add("Host", config.AzureCfg.Azure.EndPointForVoice)
	req.Header.Add("Authorization", authentication)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit")

	requestBody := []byte(`<speak version="` + config.AzureCfg.Speak.Version + `" xmlns="` + config.AzureCfg.Speak.Xmlns + `" xmlns:mstts="` + config.AzureCfg.Speak.Xmlns_mstts + `" xml:lang="` + config.AzureCfg.Speak.Xml_lang + `">
								<voice name="` + config.AzureCfg.Voice.Name + `" effect="` + config.AzureCfg.Voice.Effect + `">
									<mstts:express-as style="` + style + `">
										<prosody rate="` + config.AzureCfg.Voice.Rate + `" volume="` + config.AzureCfg.Voice.Volume + `">` +
		msg.Msg +
		`</prosody>
									</mstts:express-as>
								</voice>
							</speak>`)
	req.ContentLength = int64(len(requestBody))
	req.Header.Set("Content-Length", strconv.FormatInt(int64(len(requestBody)), 10))
	req.Body = io.NopCloser(bytes.NewReader(requestBody))
	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	msg.Mu.Lock()
	msg.Voice = base64.StdEncoding.EncodeToString(body)
	msg.VType = 3
	msg.Mu.Unlock()
}
