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
	"go_wechat/wx_util"
	"go_wechat/static"
)


type MessageBody interface {}

type ContentBody struct {
	XMLName        xml.Name `xml:"xml"`
	ToUserName     string
	FromUserName   string
	CreateTime     time.Duration
	MsgType        string
	Content        string
	MsgId          int
}

type SendContentBody struct {
	XMLName        xml.Name `xml:"xml"`
	ToUserName     string
	FromUserName   string
	CreateTime     time.Duration
	MsgType        string
	Content        string
}

type ImageBody struct {
	XMLName        xml.Name `xml:"xml"`
	ToUserName     string
	FromUserName   string
	CreateTime     time.Duration
	MsgType        string
	PicUrl         string
	MediaId        string
	MsgId          int
}

type SendImageBody struct {
	XMLName        xml.Name `xml:"xml"`
	ToUserName     string
	FromUserName   string
	CreateTime     time.Duration
	MsgType        string
	MediaId        string
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
		wx_util.TaskAccessToken()
		return nil
	} else {
		return nil
	}
}


// 获取用户发的消息
func PostContent(ctx dotweb.Context) error {
	body := ctx.Request().PostBody()

	var messageBody MessageBody

	messageBody = &ContentBody{}

	if xml.Unmarshal(body, messageBody) != nil {
		messageBody = &ImageBody{}
		xml.Unmarshal(body, messageBody)
	}

	sendMessageToUser(messageBody, ctx)


	return nil
}

// 根据收到的消息类型进行回复
func sendMessageToUser(messageBody MessageBody, ctx dotweb.Context)  {
	switch b := messageBody.(type) {
	case *ContentBody:
		sendContentBody := &SendContentBody{}
		sendContentBody.FromUserName = b.ToUserName
		sendContentBody.ToUserName = b.FromUserName
		sendContentBody.MsgType = b.MsgType
		sendContentBody.CreateTime = time.Duration(time.Now().Unix())
		sendContentBody.Content = b.Content
		bytes, _ := xml.MarshalIndent(sendContentBody, " ", " ")

		ctx.WriteString(string(bytes))
	case *ImageBody:// 需要下载收到的文件
		sendImageBody := &SendImageBody{}
		sendImageBody.FromUserName = b.FromUserName
		sendImageBody.ToUserName = b.ToUserName
		sendImageBody.MsgType = b.MsgType
		sendImageBody.MediaId = b.MediaId
		bytes, _ := xml.MarshalIndent(sendImageBody, " ", " ")

		ctx.WriteString(string(bytes))
	}
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{static.Token, timestamp, nonce}
	sort.Strings(sl)
	hash := sha1.New()
	io.WriteString(hash, strings.Join(sl, ""))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
