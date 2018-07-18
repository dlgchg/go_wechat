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

	port := 80
	server := app.StartServer(port)

	fmt.Println("dotweb.StartServer error => ", server)

}

