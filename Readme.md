# Week 09

## Homework

1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

    1. fix length: 基于固定长度。每个网络包的长度是固定的，如果少了就填充空字节，多了需要用更多的包来满足。
    2. delimiter based: 基于固定分隔符。通过一个固定的分隔符来区分包。比如 HTTP 的 /r/n
    3. length field based frame decoder: 通过传递包长字段，收信方解析出包长后进行裁切。比如 HTTP 的请求头 Content-Length 字段

2. 实现一个从 socket connection 中解码出 goim 协议的解码器

自己实现了基本逻辑后，对照 goim 源码进行参考，优化了一些代码。比如定义协议相关的常量，处理大小端问题等