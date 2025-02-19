package main

import (
	"es-box/dao"
	_ "es-box/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	dao.Init()
	beego.Run()
}
