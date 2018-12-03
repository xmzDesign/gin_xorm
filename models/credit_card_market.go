package models

import (
	"time"
)

type CreditCardMarket struct {
	Id         int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Name       string    `xorm:"not null default '' comment('名称') VARCHAR(100)"`
	Content    string    `xorm:"not null comment('内容') VARCHAR(200)"`
	ImageUrl   string    `xorm:"not null comment('图片地址') VARCHAR(200)"`
	Url        string    `xorm:"not null comment('地址') VARCHAR(256)"`
	Remarks    string    `xorm:"not null default '' comment('备注') VARCHAR(500)"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME"`
	ModifyTime time.Time `xorm:"not null comment('更新时间') DATETIME"`
	CreateName string    `xorm:"not null default '' comment('创建人') VARCHAR(50)"`
	ModifyName string    `xorm:"not null default '' comment('修改人') VARCHAR(50)"`
	ShowIndex  int       `xorm:"not null default 10 comment('首页显示 10，首页不显示  20，首页显示') INT(2)"`
}
