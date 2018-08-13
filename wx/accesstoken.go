package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/devfeel/dotweb/cache"
	"github.com/robfig/cron"
	"go_wechat/util"
	"time"
)

const GETAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

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
