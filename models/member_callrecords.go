package models

import (
	"time"
)

type MemberCallrecords struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId       int       `xorm:"comment('用户id') INT(11)"`
	CallTime     time.Time `xorm:"comment('时间') DATETIME"`
	Tenantid     string    `xorm:"comment('商户号') VARCHAR(255)"`
	Peernumber   string    `xorm:"comment('号码') VARCHAR(18)"`
	Location     string    `xorm:"comment('通话地点') VARCHAR(25)"`
	Locationtype string    `xorm:"comment('通话类型, 国内, 国际') VARCHAR(25)"`
	Durationsec  int       `xorm:"comment('通话时长') INT(11)"`
	Dialtype     string    `xorm:"comment('通话类型, 主被叫') VARCHAR(25)"`
	Fee          string    `xorm:"comment('费用') DECIMAL(10,2)"`
	GmtCreate    time.Time `xorm:"DATETIME"`
	GmtModify    time.Time `xorm:"DATETIME"`
}
