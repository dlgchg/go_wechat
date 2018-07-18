package controllers

import (
	"github.com/devfeel/dotweb"
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
	"encoding/xml"
	"time"
)

const (
	Token          = ""
	AppID          = ""
	AppSecret      = ""
	EncodingAESKey = ""
)



type ContentBody struct {
	XMLName        xml.Name `xml:"xml"`
	ToUserName     string
	FromUserName   string
	CreateTime     time.Duration
	MsgType        string
	Content        string
	MsgId          int
}


// 验证微信Signature
func CheckSignature(ctx dotweb.Context) error {
	timestamp := ctx.QueryString("timestamp")
	nonce := ctx.QueryString("nonce")
	signature := ctx.QueryString("signature")

	signatureT := makeSignature(timestamp, nonce)

	if signatureT == signature {
		echostr := ctx.QueryString("echostr")
		ctx.WriteString(echostr)
		return nil
	} else {
		return nil
	}
}


// 获取用户发的消息
func PostContent(ctx dotweb.Context) error {
	body := ctx.Request().PostBody()
	contentBody := &ContentBody{}
	xml.Unmarshal(body, contentBody)

	fmt.Println(contentBody.Content)
	return nil
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{Token, timestamp, nonce}
	sort.Strings(sl)
	hash := sha1.New()
	io.WriteString(hash, strings.Join(sl, ""))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
