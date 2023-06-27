package local

type request struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type response struct {
	Type    int    `json:"type"` //type为0时代表生成错误信息。1代表成功生成。
	Message string `json:"message"`
}
