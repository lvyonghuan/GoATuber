package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type GptConfig struct {
	// openai相关配置
	OpenAi struct {
		ApiKey           string             `mapstructure:"api_key"` //api-key
		Model            string             //使用的模型
		Temperature      float32            //对话温度
		TopP             float32            `mapstructure:"top_p"`             //代替温度采样的方法，称为核采样
		MaxTokens        int                `mapstructure:"max_tokens"`        //限制生成token使用
		Stop             []string           `mapstructure:"stop"`              //应该是生成停止标志？
		PresencePenalty  float32            `mapstructure:"presence_penalty"`  //-2.0和2.0之间的数字。正值会根据到目前为止是否出现在文本中来惩罚新标记，从而增加模型谈论新主题的可能性。 by google
		FrequencyPenalty float32            `mapstructure:"frequency_penalty"` //-2.0和2.0之间的数字。正值会根据新标记在文本中的现有频率对其进行惩罚，从而降低模型逐字重复同一行的可能性。 by google
		LogitBias        map[string]float32 `mapstructure:"logit_bias"`        //不懂，默认为nil
	}
	//代理配置
	Proxy struct {
		UseProxy bool   `mapstructure:"use_proxy"`
		ProxyUrl string `mapstructure:"proxy_url"`
	}
}

var GPTCfg GptConfig

func InitGPTConfig() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if _, err := os.Stat("nlp_model/gpt_config.example.cfg"); os.IsNotExist(err) {
		f, err := os.Create("nlp_model/gpt_config.example.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# config.toml 配置文件\n\n" +
			"# openai配置\n[openai]\n" +
			"# 你的 OpenAI API Key, 可以在 https://beta.openai.com/account/api-keys 获取\n" +
			"api_key = \"sk-xxxxxx\"\n" +
			"# 使用的模型，默认是 gpt-3.5-turbo\n" +
			"model = \"gpt-3.5-turbo\"\n" +
			"# 对话温度，越大越随机 参照https://algowriting.medium.com/gpt-3-temperature-setting-101-41200ff0d0be\n" +
			"temperature = 0.3\n" +
			"代替温度采样的方法，称为核采样。其中模型考虑具有top_p概率质量的标记的结果。所以0.1意味着只考虑构成前10%概率质量的标记。我们通常建议更改此值或对话温度，但不要同时更改两者。默认为1.\n" +
			"top_p=1" +
			"# 每次对话最大生成使用token数量\n" +
			"max_tokens = 1000\n" +
			"# stop,不太明白\n" +
			"stop=nil\n" +
			"# -2.0和2.0之间的数字。正值会根据到目前为止是否出现在文本中来惩罚新标记，从而增加模型谈论新主题的可能性。默认为0。\n" +
			"presence_penalty = 0\n" +
			"# -2.0和2.0之间的数字。正值会根据新标记在文本中的现有频率对其进行惩罚，从而降低模型逐字重复同一行的可能性。默认为0。\n" +
			"frequency_penalty = 0\n" +
			"# 不懂，默认为nil\n" +
			"logit_bias = nil\n\n" +
			"# 代理设置\n[proxy]\n" +
			"# openai是否走代理，默认关闭\n" +
			"use_proxy = false\n" +
			"# 代理地址\n" +
			"proxy_url = \"http://127.0.0.1:7890\"\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("gpt_config.example.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./nlp_model") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	err = viper.Unmarshal(&GPTCfg)
	if err != nil {
		log.Fatalf("unmarshal config failed: %v", err)
	}
}
