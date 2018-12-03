package models

import (
	"time"
)

type ApiMethod struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Method     string    `xorm:"comment('通过请求参数中的method来获取url') VARCHAR(255)"`
	Cate       string    `xorm:"comment('属于某个服务') VARCHAR(255)"`
	Url        string    `xorm:"comment('访问路径') VARCHAR(255)"`
	ApiDesc    string    `xorm:"comment('描述') VARCHAR(255)"`
	Status     int       `xorm:"comment('状态: 0:无效, 1:有效') INT(11)"`
	GmtCreate  time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify  time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName string    `xorm:"comment('修改人') VARCHAR(50)"`
}
