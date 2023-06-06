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

## 一些小说明

关于这个项目最开始的背景可以参看[old_README.md](https://github.com/lvyonghuan/GoATuber/blob/main/old_README.md)。目前项目的主要维护者都是大一在读，笔者开始写代码也就一年前不到的时间，有些地方可能（绝对）会写得很丑陋。（而且考试周快到了，我们可能抽不出很多时间来一直优化这个项目）

如果有人能够pull requests，为这个项目做优化或者添加功能，我会很感激的。

目前正在进行的项目：
- 一个可视化的配置界面
- 优化记忆模块
- 优化提示词架构

TODO:
- 语音转文字功能——也就是说可以对话了
