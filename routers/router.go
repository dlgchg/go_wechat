package routers

import (
	"github.com/devfeel/dotweb"
	"go_wechat/controllers"
)

func InitRouter(server *dotweb.HttpServer)  {
	server.Router().GET("/wechat", controllers.CheckSignature)
	server.Router().POST("/wechat", controllers.PostContent)
}
