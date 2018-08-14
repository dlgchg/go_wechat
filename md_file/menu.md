# 自定义菜单

## 1、创建

```
1、自定义菜单最多包括3个一级菜单，每个一级菜单最多包含5个二级菜单。
2、一级菜单最多4个汉字，二级菜单最多7个汉字，多出来的部分将会以“...”代替。
3、创建自定义菜单后，菜单的刷新策略是，在用户进入公众号会话页或公众号profile页时，如果发现上一次拉取菜单的请求在5分钟以前，就会拉取一下菜单，如果菜单有更新，就会刷新客户端的菜单。测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。
```

自定义菜单接口可实现多种类型按钮

```
1、click：点击推事件用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互；
2、view：跳转URL用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，可与网页授权获取用户基本信息接口结合，获得用户基本信息。
3、scancode_push：扫码推事件用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可以下发消息。
4、scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息。
5、pic_sysphoto：弹出系统拍照发图用户点击按钮后，微信客户端将调起系统相机，完成拍照操作后，会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息。
6、pic_photo_or_album：弹出拍照或者相册发图用户点击按钮后，微信客户端将弹出选择器供用户选择“拍照”或者“从手机相册选择”。用户选择后即走其他两种流程。
7、pic_weixin：弹出微信相册发图器用户点击按钮后，微信客户端将调起微信相册，完成选择操作后，将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息。
8、location_select：弹出地理位置选择器用户点击按钮后，微信客户端将调起地理位置选择工具，完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息。
9、media_id：下发消息（除文本消息）用户点击media_id类型按钮后，微信服务器会将开发者填写的永久素材id对应的素材下发给用户，永久素材类型可以是图片、音频、视频、图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
10、view_limited：跳转图文消息URL用户点击view_limited类型按钮后，微信客户端将打开开发者在按钮中填写的永久素材id对应的图文消息URL，永久素材类型只支持图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
```


接口调用请求说明

```
http请求方式：POST（请使用https协议） https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
```


参数说明

| 参数	| 是否必须	| 说明|
| --- | --- | --- |
| button	| 是	| 一级菜单数组，个数应为1~3个|
| sub_button	| 否	| 二级菜单数组，个数应为1~5个|
| type	| 是	| 菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型|
| name	| 是	| 菜单标题，不超过16个字节，子菜单不超过60个字节|
| key	| click等点击类型必须	| 菜单KEY值，用于消息接口推送，不超过128字节|
| url	| view、miniprogram类型必须	| 网页 链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。|
| media_id	| media_id类型和view_limited类型必须| 	调用新增永久素材接口返回的合法media_id|
| appid	| miniprogram类型必须| 	小程序的appid（仅认证公众号可配置）|
| pagepath	| miniprogram类型必须	| 小程序的页面路径|



返回结果

正确时的返回JSON数据包如下

```
{"errcode":0,"errmsg":"ok"}
```

错误时的返回JSON数据包如下（示例为无效菜单名长度）

```
{"errcode":40018,"errmsg":"invalid button name size"}
```


代码

menu.go

```
const AddMenuURL = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="

type MenuModel struct {
	ErrCode float64 `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

func CreateWXMenu(cache cache.Cache) {
	infos, err := ioutil.ReadFile("./menu.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	accessToken, err := cache.GetString("access_token")
	if err != nil {
		StartGetAccessTokenTimer(cache)
	} else {
		bytes, e := util.PostJSON(strings.Join([]string{AddMenuURL, accessToken}, ""), infos)

		if e != nil {
			fmt.Println("向微信发送菜单建立请求失败", err)
			return
		} else {
			menuModel := MenuModel{}
			json.Unmarshal(bytes, &menuModel)

			if menuModel.ErrCode == 0 {
				fmt.Println("向微信发送菜单建立成功")
			} else {
				fmt.Println("client向微信发送菜单建立请求失败:", menuModel.Errmsg)
			}
		}
	}
}
```

menu.json

```
{
  "button": [
    {
      "name": "进入商城",
      "type": "view",
      "url": "http://www.baidu.com/"
    },
    {

      "name":"管理中心",
      "sub_button":[
        {
          "name": "用户中心",
          "type": "click",
          "key": "molan_user_center"
        },
        {
          "name": "公告",
          "type": "click",
          "key": "molan_institution"
        }]
    },
    {
      "name": "资料修改",
      "type": "view",
      "url": "http://www.baidu.com/user_view"
    }
  ]
}
```