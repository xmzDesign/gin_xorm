package models

import (
	"time"
)

type MemberAsset struct {
	Id                  int64     `xorm:"pk autoincr comment('id') BIGINT(20)"`
	MemberId            int64     `xorm:"not null comment('用户id') index BIGINT(20)"`
	AvailableFund       string    `xorm:"not null default 0.0000 comment('用户余额') DECIMAL(10,4)"`
	FrozenFund          string    `xorm:"not null default 0.0000 comment('冻结资金') DECIMAL(10,4)"`
	TotalFund           string    `xorm:"not null default 0.0000 comment('总金额') DECIMAL(10,4)"`
	GradeLevel          int       `xorm:"not null default 0 comment('风险等级') INT(1)"`
	Status              int       `xorm:"not null comment('资金状态 10正常') INT(2)"`
	StatusExplain       string    `xorm:"not null default '' comment('状态说明') VARCHAR(200)"`
	TotalWithdrawFund   string    `xorm:"not null default 0.0000 comment('提现总金额') DECIMAL(18,4)"`
	WithdrawCount       int       `xorm:"not null default 0 comment('提现笔数') INT(11)"`
	TotalWithdrawProfit string    `xorm:"not null default 0.0000 comment('提现总手续费') DECIMAL(10,4)"`
	LoanAmt             string    `xorm:"not null default 0.0000 comment('借款总额') DECIMAL(10,4)"`
	LoanCount           int       `xorm:"not null default 0 comment('借款总笔数') INT(11)"`
	LoanFeeAmt          string    `xorm:"not null default 0.0000 comment('贷款总费用') DECIMAL(10,4)"`
	SuccessLoanCount    int       `xorm:"not null default 0 comment('成功借款次数') INT(10)"`
	SuccessLoanAmt      string    `xorm:"default 0.0000 comment('借款成功总额') DECIMAL(18,4)"`
	OverdueCount        int       `xorm:"default 0 comment('逾期次数') INT(10)"`
	OverdueAmt          string    `xorm:"default 0.0000 comment('逾期总金额') DECIMAL(18,4)"`
	Salt                string    `xorm:"not null comment('加密盐值') VARCHAR(50)"`
	BadAmt              string    `xorm:"default 0.0000 comment('坏账金额') DECIMAL(18,4)"`
	BadCount            int       `xorm:"default 0 comment('坏账次数') INT(5)"`
	LatestLoanTime      time.Time `xorm:"comment('最近借款时间') DATETIME"`
	EncryptValue        string    `xorm:"not null comment('加密值') VARCHAR(500)"`
	ModifyName          string    `xorm:"comment('修改人') VARCHAR(50)"`
	GmtCreate           time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify           time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName          string    `xorm:"comment('创建人') VARCHAR(50)"`
}
