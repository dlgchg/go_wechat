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


##### accesstoken.go

```go
// AccessToken结构体
type AccessTokenModel struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

// AccessToken错误结构体
type AccessTokenErrModel struct {
	Errcode float64 `json:"errcode"`
	Eeemsg  string  `json:"errmsg"`
}

// GET请求获取AccessToken
func getAccessTokenToWX(cache cache.Cache) error {
	url := fmt.Sprintf(GETAccessTokenURL, util.AppID, util.AppSecret)
	fmt.Printf("拼接的AccessToken URL : %s\n", url)

	body, err := util.HTTPGet(url)

	if err != nil {
		return err
	}

	if bytes.Contains(body, []byte("access_token")) {
		accessTokenModel := AccessTokenModel{}
		err := json.Unmarshal(body, &accessTokenModel)
		if err != nil {
			return err
		}
		cache.Set("access_token", accessTokenModel.AccessToken, 0)
		fmt.Printf("获取到的AccessToken:%s\n", accessTokenModel.AccessToken)
	} else {
		accessTokenErrModel := AccessTokenErrModel{}
		err := json.Unmarshal(body, &accessTokenErrModel)
		if err != nil {
			return err
		}
		fmt.Printf("发送get请求获取 微信返回 的错误信息 %+v\n", accessTokenErrModel)
	}

	return nil
}

// 获取AccessToken任务
func StartGetAccessTokenTimer(cache cache.Cache) {

	timer := cron.New()
	spec := "0 */2 * * * ?"
	timer.AddFunc(spec, func() {
		fmt.Printf("获取AccessToken的spec:%s\n", spec)
		err := getAccessTokenToWX(cache)
		if err != nil {
			fmt.Printf("获取AccessToken失败:%s\n", err)
		}
		spec = "0 */30 * * * ?"
		fmt.Printf("获取AccessToken时间:%s\n", time.Now().Format("2006-01-02 15:04:05"))
	})

	timer.Start()
	fmt.Println("获取AccessToken任务启动")

}
```


##### main.go

```go
func main() {

	app := dotweb.New()
	app.SetDevelopmentMode()
	app.SetCache(cache.NewRuntimeCache())

	wx.StartGetAccessTokenTimer(app.Cache())

	routers.InitRouter(app.HttpServer)

	e := app.StartServer(80)

	fmt.Println("go_wechat dotweb.StartServer e => ", e)

}
```