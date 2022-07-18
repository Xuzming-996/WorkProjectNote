2022-7-18
今天看到一篇文章与简化Go中的Json文章相关，因为项目内用到的配表都是json格式。
同时现在的逻辑比较繁琐，需要手动定义结构体，再通过json.unMarshal的方式反序列化到结构体中。
所以参考文章思考有没有新的解决方式。

文章提供了两个包，用于便捷访问Json格式的数据
SJSON: https://github.com/tidwall/sjson
GJSON: https://github.com/tidwall/gjson
在试验了一番后，发现其并没有简化多少，因为GO是强类型语言，想在程序启动时，通过内存访问json格式的数据一定要
解析成相应想要的数据结构。
而且json文件一定要通过系统调用读取json文件，读取到缓存中。这一步是没法省略的。所以尝试失败