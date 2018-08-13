package controllers

import (
	"github.com/devfeel/dotweb"
	"go_wechat/util"
)

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
