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

### 关于记忆

本项目在gpt模块中有一个短期记忆的流程，目前支持500个token上下的的连续短期记忆。暂时不可手动更改模式，未来规划加入。（不难，只是我懒了）

本项目的长期记忆目前通过pinecone向量数据库实现，依赖openai的function call进行调用。现目前要实现长期记忆，必须拥有一个OpenAI的api-key以获取embedding（其实还是我懒了，这个月底就更新azure openai的）（什么嘛，这不还是得要key）

要想实现长期记忆，应该进行以下步骤：
- 1.注册一个[pinecone](https://www.pinecone.io)账号。现目前免费账户大概要过几天才能使用。
- 2.在能够使用之后，创建一个index。`Index Name`随便填，但是**Dimensions**要填1536。`Metric`用默认的，其他的我没试过。pod type可以随便选，默认的就行。
- 3.在index创建完毕之后，进入创建的index的详情页，会有一个以io结尾的url。复制它，然后填入[pineconeConfig.cfg](config/MEMORY/pinecone/pineconeConfig.cfg)的url字段的引号里面。
- 4.右侧导航栏。点击`API Keys`，把api-key复制下来，填到刚刚那个配置文件的对应字段上。
- 5.把[MemoryConfig.cfg](config/MEMORY/MemoryConfig.cfg)里唯一的那个变量改成true。
- 6.没有然后了。老实说我不太推荐你使用长期记忆，因为我现在还没摸出来怎么去建构一个合适的长期记忆。但是如果你要帮忙的话，我会很感激的（

## 一些小说明

已实现GPT的函数调用功能。

我花了一晚上搓了个go语言的函数调用的轮子，函数调用可以不用很麻烦地去改动代码了。目前的问题在于函数调用可能导致整个进程的响应时间变慢——因为要多发一次请求。而且我实现的效果也不是特别好。

不保证没有bug。

2023.07.22：现在azure openai模块也已经实现了函数调用功能。

---
## 写给开发者

### GPT函数调用
基于golang的GPT的函数调用已经在`NLP/service/gpt/function`下实现。其中，`logic.go`是封装好的函数处理模块，能够让开发者方便地将自己的函数添加在发送给GPT的请求当中。

使用方法：在该目录下的[`list.go`](NLP/service/gpt/function/list.go)中的`InitFunction`函数中使用`addFunc`函数将自己编写的函数进行添加操作。在[`json.go`](NLP/service/gpt/function/json.go)中编写request所需的json格式的对函数的说明。

完成后即可实现将自己编写的函数纳入GPT函数调用选择的一部分。使用函数调用可以实现诸多功能，例如网页爬取等操作。

### 使用自定义语言模型

基于http协议，本项目现在允许开发者使用自定义的语言模型。

首先更改[`NLOConfig.cfg`](config/NLP/NLPConfig.cfg)的配置，将`use_local_model`设置为true，其他设置为false。

再在[`localConfig.cfg`](config/NLP/localConfig/localConfig.cfg)中填写通信的请求地址。请事先暴露出要使用模型的API地址。

请求发送的body结构如下：
```json
{
  "message": "你好",
  "username": "lvyonghuan"
}
```
message为弹幕或聊天信息，username为该信息发送者的用户名。

请构造如下格式的返回体：

```json
{
	"type": 1,
	"message": "你好，有什么可以帮助你的吗？"
}
```
type为int字段。type为0代表生成过程出现错误，可以将错误信息写入message字段，项目将在cmd窗口中输出错误信息。

type为1时，代表生成成功。message字段为AI生成的回应。

请自行实现记忆等功能。

---

*目前项目维护者都在修炼当中，更新暂停一段时间。但是如果使用中产生了问题要提issue或者想提pull request都是欢迎的。*

关于这个项目最开始的背景可以参看[old_README.md](https://github.com/lvyonghuan/GoATuber/blob/main/old_README.md)。目前项目的主要维护者都是大一在读，笔者开始写代码也就一年前不到的时间，有些地方可能（绝对）会写得很丑陋。

如果有人能够pull requests，为这个项目做优化或者添加功能，我会很感激的。

目前正在进行的项目：
- 一个可视化的配置界面
- 优化记忆模块
- 优化提示词架构

TODO:
- 语音转文字功能——也就是说可以对话了
- 支持VRM模型
