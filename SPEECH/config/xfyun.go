package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Xfyun struct {
	Xfyun struct {
		ApiKey    string `mapstructure:"api_key"`
		ApiSecret string `mapstructure:"api_secret"`
		AppID     string `mapstructure:"app_id"`
	}
	XfyunVoice struct {
		Aue    string `mapstructure:"aue"`    //音频编码
		Sfl    int    `mapstructure:"sfl"`    //流式返回
		Auf    string `mapstructure:"auf"`    //音频采样率
		Vcn    string `mapstructure:"vcn"`    //发音人
		Speed  int    `mapstructure:"speed"`  //语速
		Volume int    `mapstructure:"volume"` //音量
		Pitch  int    `mapstructure:"pitch"`  //音高
		Reg    string `mapstructure:"reg"`    //英文发音方式
		Rdn    string `mapstructure:"rdn"`    //数字发音方式
	}
}

var XFCfg Xfyun

func InitXFConfig() {
	if _, err := os.Stat("config/SPEECH/XFConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/SPEECH/XFConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# 讯飞用户配置（待分离）\n[xfyun]\n" +
			"app_id = \"xxxxxx\"\n" +
			"api_secret = \"xxxxxx\"\n" +
			"api_key = \"xxxxxx\"\n" +
			"# 讯飞语音接口配置\n[xfyunVoice]\n" +
			"# 音频编码（\n" +
			"aue = \"lame\" \n" +
			"# 是否开启流式返回（1开启，0关闭）（配合aue=lame使用）\n" +
			"sfl = 1\n" +
			"# 音频采样率(8k或者16k)\n" +
			"auf = \"audio/L16;rate=16000\"\n" +
			"# 发音人\n" +
			"vcn = \"xiaoyan\"\n" +
			"# 语速(0~100,默认50)\n" +
			"speed = 50\n" +
			"# 音量(0~100,默认50)\n" +
			"volume = 50\n" +
			"# 音高(0~100,默认50)\n" +
			"pitch = 50\n" +
			"# 英文发音方式\n" +
			"reg = \"0\"\n" +
			"# 数字发音方式\n" +
			"rdn = \"0\"\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("XFConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/SPEECH") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&XFCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
