# GO!Atuber
 
 快速构建属于你的AI主播

## 快速开始

- 修改 `config/CHAT/config.cfg` 中的room_id，填入直播间的房间号。
- 修改 `config/NLP/GPTConfig/gptConfig.cfg` 中的相关信息，完善gpt的配置。
- 把live2d模型文件拖入 `dist/model` 中。具体参考这里的[README.md](https://github.com/lvyonghuan/GoATuber/tree/main/dist/model)。
- 配置代理。请修改`config/proxy/proxyConfig.cfg`。不配置代理，可能无法正常调用OpenAI等服务。
- 运行 `GoTuber.exe`。如遇闪退，请检查配置是否正常。
- 浏览器打开 `http://127.0.0.1:9000/` 。
- Enjoy！
