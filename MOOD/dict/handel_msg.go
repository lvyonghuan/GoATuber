package dictMOOD

import (
	sensitive "GoTuber/MESSAGE/filter"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var searchMessage = make(chan string, 1)
var getWords = make(chan []string, 1)

// HandelMsg 处理消息
func HandelMsg(Msg sensitive.OutPut) {
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
		_, moodWord = search.Validate(moodWord)

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
	sum := mood.Happy + mood.Mad + mood.Sad + mood.Disgust + mood.Surprise + mood.Fear + mood.Health
	happy := mood.Happy / sum
	mad := mood.Mad / sum
	disgust := mood.Disgust / sum
	surprise := mood.Surprise / sum
	fear := mood.Fear / sum
	health := mood.Health / sum
	maxMood := max(happy, mad, disgust, surprise, fear, health)
	log.Println(maxMood, "happy:", happy, "mad:", mad, "disgust:", disgust, "surprise:", surprise, "fear:", fear, "health:", health)
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

func max(vars ...interface{}) string {
	max := vars[0]
	varName := ""
	for _, v := range vars {
		if v.(float64) > max.(float64) {
			max = v
			varName = fmt.Sprintf("%v", reflect.TypeOf(v))
		}
	}
	return varName
}
