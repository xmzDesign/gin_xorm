package models

import (
	"time"
)

type CashProduct struct {
	Id          int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Name        string    `xorm:"not null comment('产品名称') VARCHAR(32)"`
	Code        string    `xorm:"not null comment('产品编号') VARCHAR(32)"`
	MaxAmount   string    `xorm:"not null comment('最大交易金额') DECIMAL(10,2)"`
	MinAmount   string    `xorm:"not null default 0.00 comment('最小交易金额') DECIMAL(10,2)"`
	MaxRate     string    `xorm:"DECIMAL(4,4)"`
	CostRate    string    `xorm:"not null comment('成本扣率') DECIMAL(4,4)"`
	DefaultRate string    `xorm:"not null comment('默认扣率') DECIMAL(4,4)"`
	Formula     int       `xorm:"not null comment('计算方式') INT(4)"`
	Fee         string    `xorm:"not null comment('单笔手续费') DECIMAL(4,2)"`
	Status      int       `xorm:"comment('状态 10有效 90作废') INT(4)"`
	CreateTime  time.Time `xorm:"not null comment('创建时间') DATETIME"`
	CreateName  string    `xorm:"not null comment('创建人') VARCHAR(32)"`
	ModifyTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	ModifyName  string    `xorm:"not null default '' comment('修改人') VARCHAR(50)"`
	ChannelId   int       `xorm:"not null default 0 comment('渠道号') INT(11)"`
}
