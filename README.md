# GO!Atuber
 
 快速构建属于你的AI主播

## 快速开始

- 配置go语言环境。因为目前版本依然为测试版本，所以没有release。~~如何配置请STFW~~。
- 修改 `CHAT\config.cfg` 中的room_id，填入直播间的房间号。
- 修改 `NLP\service\gpt\gptConfig.cfg` 中的相关信息，完善gpt的配置。
- 把live2d模型文件拖入 `dist\model` 中。具体参考这里的[README.md](https://github.com/lvyonghuan/GoATuber/tree/main/dist/model)。
- 配置代理。config文件在proxy文件夹下。不配置代理，可能无法正常调用OpenAI等服务。
- 运行 `cmd.go`。
- 浏览器打开 `http://127.0.0.1:9000/` 。因为程序依靠管道进行信息传输控制，运行程序后不及时（具体来说，在文本生成并向前端传输之前）打开这个页面会造成程序阻塞。如果出现这种情况，请重新运行程序。
- Enjoy！
