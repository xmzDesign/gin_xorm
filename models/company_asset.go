package models

import (
	"time"
)

type CompanyAsset struct {
	Id                    int       `xorm:"not null pk autoincr INT(11)"`
	TotalAsset            string    `xorm:"not null default 0.0000 comment('总资金') DECIMAL(10,4)"`
	TotalTransferOutAsset string    `xorm:"not null default 0.0000 comment('转出总金额') DECIMAL(10,4)"`
	TotalTransferInAsset  string    `xorm:"not null default 0.0000 comment('转入总金额') DECIMAL(10,4)"`
	TotalWithdrawProfit   string    `xorm:"not null default 0.0000 comment('提现收益总金额') DECIMAL(10,4)"`
	WithdrawCount         int       `xorm:"not null default 0 comment('提现总笔数') INT(11)"`
	TotalLoanProfit       string    `xorm:"not null default 0.0000 comment('贷款收益总金额') DECIMAL(10,4)"`
	LoanCount             int       `xorm:"not null default 0 comment('贷款总笔数') INT(11)"`
	Salt                  string    `xorm:"not null default '' comment('加密盐值') VARCHAR(100)"`
	EncryptValue          string    `xorm:"not null comment('加密值') VARCHAR(500)"`
	GmtCreate             time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify             time.Time `xorm:"comment('更新时间') DATETIME"`
}
