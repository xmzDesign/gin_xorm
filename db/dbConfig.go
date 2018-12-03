package db

import (
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"log"
	"time"
)

var engine *xorm.Engine

func init() {
	var err error
	//读取配置文件
	cfg, err := ini.Load("config/db.ini")
	if err != nil {
		log.Fatal(err)
	}
	username := cfg.Section("mysql").Key("username").Value()
	password := cfg.Section("mysql").Key("password").Value()
	url := cfg.Section("mysql").Key("url").Value()

	source := fmt.Sprintf("%s:%s@%s", username, password, url)
	// 初始化xorm的engine
	engine, err = xorm.NewEngine("mysql", source)
	if err != nil {
		log.Fatalf("db error: %#v\n", err.Error())
	}
	//加载自定义sql文件
	err = engine.RegisterSqlMap(xorm.Xml("./db/xml", ".xml"))
	if err != nil {
		log.Fatalf("db error: %#v\n", err.Error())
	}
	err = engine.RegisterSqlTemplate(xorm.Default("./db/tpl", ".sql"))
	if err != nil {
		log.Fatalf("db error: %#v\n", err.Error())
	}

	err = engine.StartFSWatcher()
	if err != nil {
		log.Printf("sql parse error: %#v\n", err)
	}

	err = engine.Ping()
	if err != nil {
		log.Fatalf("db connect error: %#v\n", err.Error())
	}

	// 30minute ping db to keep connection
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("db connect error: %#v\n", err.Error())
			}
		}
	}(engine)

	//答应日志
	engine.ShowSQL(true)
	//设置连接池大小
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(5)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})
}
