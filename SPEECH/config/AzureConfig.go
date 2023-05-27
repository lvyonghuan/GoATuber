package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
	"time"
)

type AzureConfig struct {
	Azure struct {
		Key              string `mapstructure:"key"`
		EndPoint         string `mapstructure:"end_point"`
		EndPointForVoice string
	}
	Speak struct {
		Version     string `mapstructure:"version"`
		Xml_lang    string `mapstructure:"xml_lang"`
		Xmlns_mstts string `mapstructure:"xmlns_mstts"`
		Xmlns       string `mapstructure:"xmlns"`
	}
	Voice struct {
		Name   string `mapstructure:"name"`
		Effect string `mapstructure:"effect"`
		Rate   string `mapstructure:"rate"`
		Volume string `mapstructure:"volume"`
	}
}

var AzureCfg AzureConfig

func InitAzureConfig() {
	if _, err := os.Stat("config/SPEECH/AzureConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/SPEECH/AzureConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# azure配置\n[azure]\n" +
			"# key,两个填一个就行了\n" +
			"key = \"xxxxxxx\"\n" +
			"# voice endpoint,语音终结点\n" +
			"end_point = \"xxxxxx\"\n\n" +
			"# speak根元素\n[speak]\n" +
			"# version 指示用于解释文档标记的 SSML 规范的版本。 当前版本为“1.0”。\n" +
			"version = \"1.0\"\n" +
			"# xml:lang 根文档的语言。 该值可以包含语言代码，例如 en（英语），也可以包含区域设置，例如 en-US（美国英语）。\n" +
			"xml_lang = \"cn\"\n" +
			"# xmlns:mstts\n" +
			"xmlns_mstts = \"https://www.w3.org/2001/mstts\"\n" +
			"# xmlns 用于定义 SSML 文档的标记词汇（元素类型和属性名称）的文档的 URI。 当前 URI 为 \"http://www.w3.org/2001/10/synthesis\"。\n" +
			"xmlns = \"http://www.w3.org/2001/10/synthesis\"\n" +
			"# voice元素\n[voice]\n" +
			"# name 用于文本转语音输出的语音。 有关受支持的预生成语音的完整列表，请参阅语言支持(https://learn.microsoft.com/zh-cn/azure/cognitive-services/speech-service/language-support?tabs=tts)。\n" +
			"name = \"zh-CN-XiaoyiNeural\"\n" +
			"# effect 音频效果处理器，用于在设备上针对特定方案优化合成语音输出的质量。 可选。\n" +
			"effect = \"\"\n" +
			"# rate 语速，填写比默认值高或低的百分比\n" +
			"rate = \"+0.00%\"\n" +
			"# volume 音量，基准值为默认值\n" +
			"volume = \"+0.00%\""))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("AzureConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/SPEECH") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&AzureCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
	re := regexp.MustCompile(`https:\/\/(\w+)\.`)
	match := re.FindStringSubmatch(AzureCfg.Azure.EndPoint)
	if match == nil {
		log.Fatalf("azure调取模块，config设置：endpoint获取错误")
	} else {
		AzureCfg.Azure.EndPointForVoice = match[0] + "tts.speech.microsoft.com/cognitiveservices/v1"
	}
}
