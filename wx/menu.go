package wx

import (
	"encoding/json"
	"fmt"
	"github.com/devfeel/dotweb/cache"
	"go_wechat/util"
	"io/ioutil"
	"strings"
)

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
