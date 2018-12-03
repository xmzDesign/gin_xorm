package models

import (
	"time"
)

type CommonsBankCode struct {
	Id         int       `xorm:"not null INT(11)"`
	Paysysbank int       `xorm:"INT(11)"`
	Name       string    `xorm:"VARCHAR(100)"`
	Phonenumbe string    `xorm:"VARCHAR(50)"`
	Bankaddr   string    `xorm:"VARCHAR(100)"`
	GmtCreate  time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify  time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName string    `xorm:"comment('修改人') VARCHAR(50)"`
}
