package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"fmt"
)

var db *gorm.DB
var once sync.Once

func initDB() {
	var err error
	cnf, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic("数据库配置文件读取失败。请将conf_reference.json复制为conf.json后，修改conf.json中的配置")
	}
	print(cnf)
	results := gjson.GetManyBytes(cnf, "dataSource.host", "dataSource.port", "dataSource.database", "dataSource.username", "dataSource.password")
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", results[3].String(), results[4].String(), results[0].String(), results[1].Int(), results[2].String())
	db, err = gorm.Open("mysql", args)
	if err != nil {
		log.Fatal(err)
	}
}

func DB() *gorm.DB {
	once.Do(initDB)
	return db
}
