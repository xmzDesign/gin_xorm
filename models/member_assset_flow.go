package models

import (
	"time"
)

type MemberAsssetFlow struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	GmtCreate  time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify  time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName string    `xorm:"comment('修改人') VARCHAR(50)"`
}
