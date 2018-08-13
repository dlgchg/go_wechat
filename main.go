package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"go_wechat/routers"
)

func main() {

	app := dotweb.New()
	app.SetDevelopmentMode()

	routers.InitRouter(app.HttpServer)

	error := app.StartServer(80)

	fmt.Println("go_wechat dotweb.StartServer error => ", error)
}
