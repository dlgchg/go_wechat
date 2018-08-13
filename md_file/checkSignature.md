# 验证服务器配置

##验证消息的确来自微信服务器

| 参数 |	描述 |
| --- | --- |
| signature | 微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。|
| timestamp	| 时间戳|
| nonce	| 随机数|
| echostr	| 随机字符串|