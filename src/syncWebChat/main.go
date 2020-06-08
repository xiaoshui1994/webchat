package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gcmd"
	_ "syncWebChat/routers"
)

func main() {

	var db gdb.DB

	parser, _ := gcmd.Parse(map[string]bool{
		"name": true,
		"type": true,
	}, false)

	xx := gdb.ConfigNode{
		Host:             "127.0.0.1",
		Port:             "3306",
		User:             "root",
		Pass:             "123456",
		Name:             parser.GetOpt("name", ""),
		Type:             parser.GetOpt("type", "mysql"),
		Role:             "master",
		Charset:          "utf8",
		Weight:           1,
		MaxIdleConnCount: 10,
		MaxOpenConnCount: 10,
		MaxConnLifetime:  600,
	}

	gdb.AddConfigNode("default", xx)

	fmt.Println(xx)

	if r, err := gdb.New(); err != nil {
		fmt.Println(err)
	} else {
		db = r
	}


	schemaTemplate := "CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET UTF8"
	db.Exec(fmt.Sprintf(schemaTemplate, "test_internal"))

	db.SetSchema("test_internal")

	beego.Run()
}

