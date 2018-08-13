package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/cache"
	"go_wechat/routers"
	"go_wechat/wx"
)

func main() {

	app := dotweb.New()
	app.SetDevelopmentMode()
	app.SetCache(cache.NewRuntimeCache())

	wx.StartGetAccessTokenTimer(app.Cache())

	routers.InitRouter(app.HttpServer)

	e := app.StartServer(80)

	fmt.Println("go_wechat dotweb.StartServer e => ", e)

}
