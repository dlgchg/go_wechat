package wx_util

import (
	"strings"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"go_wechat/db_util"
	"time"
	"github.com/robfig/cron"
	"go_wechat/static"
	"log"
)

type AccessTokenModel struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type AccessToken struct {
	AccessToken  string
	CreateTime   string
}

func GetAccessToken(url, appid, secret string) (*AccessTokenModel, error) {
	atUrl := strings.Join([]string{url, "?grant_type=client_credential&appid=", appid, "&secret=", secret}, "")
	resp, err := http.Get(atUrl)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil ,nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil ,nil
	}

	if bytes.Contains(body, []byte("access_token")) {
		model := &AccessTokenModel{}
		err := json.Unmarshal(body, model)
		if err != nil {
			return nil, nil
		}
		return model, nil
	} else {
		return nil ,nil
	}
}

func SaveAccessToken(model *AccessTokenModel) error {
	session := db_util.ConnectMongoDB()

	collection := session.DB("wx_wechat").C("accesstoken")

	e := collection.Insert(&AccessToken{model.AccessToken, time.Now().String()})
	if e != nil {
		return e
	}

	return nil
}

func TaskAccessToken()  {
	c := cron.New()
	spec := "0 */1 * * * ?"
	c.AddFunc(spec, func() {
		s, e := GetAccessToken("https://api.weixin.qq.com/cgi-bin/token", static.AppID, static.AppSecret)
		if e == nil {
			log.Println(e)
			SaveAccessToken(s)
		}
	})
	c.Start()
}
