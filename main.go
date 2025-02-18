package main

import (
	_ "es-box/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

