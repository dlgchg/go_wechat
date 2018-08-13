# 验证服务器配置

## 验证消息的确来自微信服务器

| 参数 |	描述 |
| --- | --- |
| signature | 微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。|
| timestamp	| 时间戳|
| nonce	| 随机数|
| echostr	| 随机字符串|


### 1、将token、timestamp、nonce三个参数进行字典序排序，将三个参数字符串拼接成一个字符串进行sha1加密

```go
func Signature(params ...string) string  {
	sort.Strings(params)
	hash := sha1.New()
	for _, v := range params {
		io.WriteString(hash, v)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
```


### 2、GET请求来自微信服务器,signature对比,返回echostr参数内容，则接入生效，成为开发者成功，否则接入失败


```go
func GetGignature(ctx dotweb.Context) error {

	timestamp := ctx.QueryString("timestamp")
	nonce := ctx.QueryString("nonce")
	signature := ctx.QueryString("signature")

	signaturen := util.Signature(util.Token, timestamp, nonce)

	if signature == signaturen {
		echostr := ctx.QueryString("echostr")
		ctx.WriteString(echostr)
	}

	return nil
}
```

### 3、另请注意，微信公众号接口必须以http://或https://开头，分别支持80端口和443端口


```go
func main() {

	app := dotweb.New()
	app.SetDevelopmentMode()

	routers.InitRouter(app.HttpServer)

	error := app.StartServer(80)

	fmt.Println("go_wechat dotweb.StartServer error => ", error)
}
```

“main 包中的不同的文件的代码不能相互调用，其他包可以”