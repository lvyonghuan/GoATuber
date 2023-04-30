package dictMOOD

import (
	sensitive "GoTuber/MESSAGE/filter"
	"log"
	"strconv"
	"strings"
)

var searchMessage = make(chan string, 1)
var getWords = make(chan []string, 1)

// HandelMsg 处理消息
func HandelMsg(Msg *sensitive.OutPut) {
	msg := Msg.Msg
	search := sensitive.New()
	err := search.LoadWordDict("MOOD/dict/情感词汇本体.txt", 0) //捏麻的，别问我为什么要用两次AC自动机
	if err != nil {
		log.Printf("情感字典加载失败：%v", err)
		return
	}
	searchMessage <- msg
	moodWords := <-getWords
	var mood Mood
	for _, moodWord := range moodWords {
		moodWord = search.SearchLine(moodWord)
		words := strings.Split(moodWord, ",")
		for i, word := range words {
			if i == 1 || i == 3 {
				num, err := strconv.Atoi(words[i+1]) //获得情感值
				if err != nil {
					log.Print("MOOD,转化数字失败：", err)
				}
				if i == 3 {
					num = int(float32(num) * 0.5)
				}
				if word == "PA" || word == "PE" {
					mood.Happy += num
				} else if word == "PD" || word == "PH" || word == "PG" || word == "PB" || word == "PK" {
					mood.Health += num
				} else if word == "NA" {
					mood.Mad += num
				} else if word == "NB" || word == "NJ" || word == "NH" || word == "PF" {
					mood.Sad += num
				} else if word == "NI" || word == "NC" || word == "NG" {
					mood.Fear += num
				} else if word == "NE" || word == "ND" || word == "NN" || word == "NK" || word == "NL" {
					mood.Disgust += num
				} else if word == "PC" {
					mood.Surprise += num
				}
			}
		}
	}
	var moodNum [7]float64
	moodNum[0] = float64(mood.Happy)
	moodNum[1] = float64(mood.Mad)
	moodNum[2] = float64(mood.Sad)
	moodNum[3] = float64(mood.Disgust)
	moodNum[4] = float64(mood.Surprise)
	moodNum[5] = float64(mood.Fear)
	moodNum[6] = float64(mood.Health)
	maxMood := max(moodNum)
	//log.Println(maxMood, "happy:", moodNum[0], "mad:", moodNum[1], "sad:", moodNum[2], "disgust:", moodNum[3], "surprise:", moodNum[4], "fear:", moodNum[5], "health:", moodNum[6])
	Msg.Mu.Lock()
	Msg.Mood = maxMood
	Msg.Mu.Unlock()
}

func Search() {
	search := sensitive.New()
	err := search.LoadWordDict("MOOD/dict/情感词汇本体.txt", 1)
	if err != nil {
		log.Printf("情感字典加载失败：%v", err)
		return
	}
	for {
		select {
		case msg := <-searchMessage:
			words := search.FindAll(msg)
			getWords <- words
		}
	}
}

func max(moodNum [7]float64) string {
	var max float64
	var maxFlag int
	for i := 0; i < 7; i++ {
		if moodNum[i] > max {
			max = moodNum[i]
			maxFlag = i
		}
	}
	if maxFlag == 0 {
		return "happy"
	} else if maxFlag == 1 {
		return "mad"
	} else if maxFlag == 2 {
		return "sad"
	} else if maxFlag == 3 {
		return "disgust"
	} else if maxFlag == 4 {
		return "surprise"
	} else if maxFlag == 5 {
		return "fear"
	} else if maxFlag == 6 {
		return "health"
	}
	return ""
}
