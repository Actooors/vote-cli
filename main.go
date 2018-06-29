package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"fmt"
	"github.com/Actooors/vote-cli/app"
)

func main() {
	//app.CreateTables()
	//err := app.InitInsert()
	//if err != nil {
	//	panic(err.Error())
	//}
	router := gin.Default()
	g := router.Group("vote-cli")
	{
		g.PATCH("clear", app.ClearHandler)
	}

	conf, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic("配置文件读取失败")
	}
	address := gjson.GetBytes(conf, "address")
	port := gjson.GetBytes(conf, "port")
	router.Run(fmt.Sprintf("%s:%s", address.String(), port.String()))
}
