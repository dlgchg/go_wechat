package controllers

import (
	"github.com/devfeel/dotweb"
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
)

const token = "xiaoqingxin"

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

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	hash := sha1.New()
	io.WriteString(hash, strings.Join(sl, ""))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
