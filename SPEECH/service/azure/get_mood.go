package azure

var style string

// 获取情感
func getMood(mood string) {
	if mood == "happy" {
		style = "friendly"
	} else if mood == "mad" {
		style = "angry"
	} else if mood == "sad" {
		style = "sad"
	} else if mood == "disgust" {
		style = "disgruntled"
	} else if mood == "surprise" {
		style = "excited"
	} else if mood == "fear" {
		style = "terrified"
	} else if mood == "health" {
		style = "chat"
	}
}
