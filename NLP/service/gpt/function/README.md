# 函数存储

设想：一个便捷的函数存储子模块——附着在GPT的调用上，使开发者能够方便地添加功能。

要求：统一传入字符串切片，统一传出字符串。

list：用于便捷地添加函数

function：在这里实现函数的具体功能

logic：存储、判断、获取、执行各函数，执行后返回结果

我还没想好该怎么处理传给OpenAI的JSON。