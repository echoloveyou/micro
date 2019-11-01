package router

import (
	"github.com/astaxie/beego"
	"github.com/echoloveyou/micro/bos_web/controllers"
)

func LoadRouters() {
	ns := beego.NewNamespace("/api/bos",
		beego.NSNamespace("/admin_user",
			beego.NSRouter("/add", &controllers.AdminUserController{}, "post:AdminUserAdd"),
		),
	)
	beego.AddNamespace(ns)
}
