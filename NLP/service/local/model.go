package local

type request struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type response struct {
	Type    int    `json:"type"`
	Message string `json:"message"`
}
