package models

import (
	"time"
)

type Activity struct {
	Id                 int       `xorm:"not null pk autoincr unique INT(11)"`
	Name               string    `xorm:"not null default '' comment('活动名称') VARCHAR(64)"`
	Brief              string    `xorm:"not null default '' comment('活动简要说明') VARCHAR(255)"`
	EventCode          string    `xorm:"not null default '' comment('活动事件代码') VARCHAR(32)"`
	Status             int       `xorm:"not null default 0 comment('活动状态 10进行中 20临时关闭 90删除') INT(2)"`
	IsTimeLimit        int       `xorm:"not null default 0 comment('活动时间限制 10 无时间限制 20 时间限制') INT(2)"`
	StartTime          time.Time `xorm:"comment('活动开始时间') DATETIME"`
	EndTime            time.Time `xorm:"comment('活动结束时间') DATETIME"`
	TicketUseTyoe      int       `xorm:"not null default 0 comment('卡券有效类型 10 领取时间+有效天数 20时间段有效') INT(2)"`
	TicketUseStartTime time.Time `xorm:"comment('卡券使用开始时间') DATETIME"`
	TicketUseEndTime   time.Time `xorm:"comment('卡券使用结束时间') DATETIME"`
	TicketUseValue     int       `xorm:"not null default 0 comment('卡券有效时间值') INT(2)"`
	Remark             string    `xorm:"not null default '' comment('活动说明') VARCHAR(255)"`
	IsMultiple         int       `xorm:"not null default 0 comment('是否可重复领取 0 表示无限制 1表示限制一次 2表示限制2次') INT(2)"`
	TicketMode         int       `xorm:"not null default 0 comment('10：满减  参数1：满足的金额   参数2：减免的金额    20：立减    参数1：减免的金额     30：折扣：参数1：折扣系数，例如0.85') INT(2)"`
	TicketParam1       string    `xorm:"not null default '' comment('参数1') VARCHAR(32)"`
	TicketParam2       string    `xorm:"not null default '' comment('参数2') VARCHAR(32)"`
	TicketPrice        string    `xorm:"not null default 0.00 comment('卡券价值') DECIMAL(4,2)"`
	GmtCreate          time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify          time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName         string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName         string    `xorm:"comment('修改人') VARCHAR(50)"`
}
