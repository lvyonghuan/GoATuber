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
	Event string //触发act的事件
	Act   string //行为名称
}

var MoodAct = make(map[string][]Mood)

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
				log.Println("获取情感-动作、表情映射类型失败，错误原因：", err)
				return
			}
			mood := &Mood{
				Event: act[1],
				Act:   act[2],
			}
			str := strings.Replace(act[3], "\r", strconv.Itoa(typ), 1)
			str = strings.Replace(str, "\n", "", 1)
			if typ == 1 {
				MoodAct[str] = append(MoodAct[str], *mood)
			} else if typ == 2 {
				MoodAct[str] = append(MoodAct[str], *mood)
			} else {
				log.Println("非法情感行为类型")
			}
		}
	}
}
