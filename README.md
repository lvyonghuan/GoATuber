# GO!Atuber
 
 快速构建属于你的AI主播

## 快速开始

- 修改 `CHAT\config.cfg` 中的room_id，填入直播间的房间号。
- 修改 `NLP\service\gpt\gptConfig.cfg` 中的相关信息，完善gpt的配置。
- 把live2d模型文件拖入 `dist\module` 中，修改 `frontend/backend/get_live2d_module_info/live2d_config.cfg` 中的模型名字信息。将模型文件里motions目录下的model3.json前面的字段填入即可（现只支持此类模型）。
- 配置代理。config文件在proxy文件夹下。
- 运行 `cmd.go`。
- 浏览器打开 `http://127.0.0.1:9000/` 。Enjoy！
