package models

import (
	"time"
)

type AppRelease struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	ProductCode     string    `xorm:"comment('产品编号 10,兴趣贷') VARCHAR(255)"`
	Devicetype      int       `xorm:"comment('设备类型 10,android 20,ios') INT(2)"`
	Releaseversion  string    `xorm:"comment('发布包版本') VARCHAR(255)"`
	Buildversion    int       `xorm:"comment('程序内建版本') INT(5)"`
	Pkgsize         int64     `xorm:"BIGINT(20)"`
	DescriptionText string    `xorm:"VARCHAR(500)"`
	Issuedate       time.Time `xorm:"comment('发行时间') DATETIME"`
	Downloadurl     string    `xorm:"VARCHAR(255)"`
	Forceupdate     int       `xorm:"comment('10,强制 20,非强制') INT(2)"`
	GmtCreate       time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify       time.Time `xorm:"comment('更新时间') DATETIME"`
	Status          int       `xorm:"INT(10)"`
	CreateName      string    `xorm:"comment('创建人') VARCHAR(255)"`
	ModifyName      string    `xorm:"comment('修改人') VARCHAR(255)"`
}
