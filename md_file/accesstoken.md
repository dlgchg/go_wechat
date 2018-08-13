# 获取access_token

access_token是公众号的全局唯一接口调用凭据，公众号调用各接口时都需使用access_token。开发者需要进行妥善保存。access_token的存储至少要保留512个字符空间。access_token的有效期目前为2个小时，需定时刷新，重复获取将导致上次获取的access_token失效。

#### 接口调用请求说明

```
https请求方式: GET
https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
```


#### 参数说明

| 参数	| 是否必须	| 说明 |
| --- | --- | --- |
| grant_type	| 是| 	获取access_token填写client_credential |
| appid	| 是	| 第三方用户唯一凭证|
| secret	| 是	| 第三方用户唯一凭证密钥，即appsecret|

#### 返回说明

正常情况下，微信会返回下述JSON数据包给公众号：

```
{"access_token":"ACCESS_TOKEN","expires_in":7200}
```

#### 参数说明
| 参数	| 说明 |
| --- | --- |
| access_token	| 获取到的凭证|
| expires_in	| 凭证有效时间，单位：秒|


```
{"errcode":40013,"errmsg":"invalid appid"}
```


#### 返回码说明

| 返回码	| 说明|
| --- | --- |
| -1	| 系统繁忙，此时请开发者稍候再试|
| 0	| 请求成功|
| 40001	| AppSecret错误或者AppSecret不属于这个公众号，请开发者确认AppSecret的正确性|
| 40002	| 请确保grant_type字段值为client_credential|
| 40164	| 调用接口的IP地址不在白名单中，请在接口IP白名单中进行设置。（小程序及小游戏调用不要求IP地址在白名单内。）|