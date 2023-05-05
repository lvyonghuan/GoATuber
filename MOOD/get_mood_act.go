package MOOD

//这个文件是为了更好地适配任意模型而设置的。

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mood struct {
	Type int    //类型，1代表动作，2代表表情
	Act  string //行为名称
	Mood string //对应情绪
}

var MoodAct []Mood

func readMoodAct() {
	file, err := os.Open("./MOOD/mood.cfg")
	if err != nil {
		log.Fatalf("open config file failed: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for i := 1; ; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if i >= 2 {
			act := strings.Split(line, ",")
			typ, err := strconv.Atoi(act[0])
			if err != nil {
				log.Fatal("获取情感-动作、表情映射类型失败，错误原因：", err)
				return
			}
			mood := &Mood{
				Type: typ,
				Act:  act[1],
				Mood: act[2],
			}
			MoodAct = append(MoodAct, *mood)
		}
	}
}
