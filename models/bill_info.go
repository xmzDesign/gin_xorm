package models

import (
	"time"
)

type BillInfo struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	MemberId          int64     `xorm:"comment('用户ID') index BIGINT(20)"`
	Stages            int       `xorm:"default 0 comment('期数') INT(11)"`
	LoanDays          int       `xorm:"default 0 comment('借款天数') INT(3)"`
	ProductType       int       `xorm:"default 0 comment('按天，按月') INT(2)"`
	ProductCode       string    `xorm:"default '0' comment('借款产品编号') index VARCHAR(30)"`
	RateNo            string    `xorm:"default '' comment('费率编号') VARCHAR(50)"`
	LoanId            int       `xorm:"default 0 comment('借款单ID') index INT(11)"`
	LoanTime          time.Time `xorm:"comment('借款时间') DATETIME"`
	LoanAmt           string    `xorm:"default 0.00 comment('借款金额') DECIMAL(18,2)"`
	Amount            string    `xorm:"default 0.00 comment('账单金额') DECIMAL(18,2)"`
	DueAmt            string    `xorm:"default 0.00 comment('应还总金额') DECIMAL(18,2)"`
	CreditAuditCost   string    `xorm:"default 0.00 comment('(贷前扣除)额度审核费') DECIMAL(18,2)"`
	CreditInspectCost string    `xorm:"default 0.00 comment('(贷前扣除)征信查询费') DECIMAL(18,2)"`
	ManageCost        string    `xorm:"default 0.00 comment('(贷后扣除)管理费') DECIMAL(18,2)"`
	LoanInterestCost  string    `xorm:"default 0.00 comment('(贷后扣除)利息') DECIMAL(18,2)"`
	IntroduceCost     string    `xorm:"default 0.00 comment('(贷前扣除)介绍费') DECIMAL(18,2)"`
	LoanProvider      string    `xorm:"comment('贷款方(qlkj:全联科技,cbc:建行小贷)') VARCHAR(20)"`
	IsOverdue         int       `xorm:"default 0 comment('是否逾期, 0:未逾期, 1:已逾期') INT(2)"`
	OverdueAmt        string    `xorm:"default 0.00 comment('逾期总费用') DECIMAL(18,2)"`
	DebtAmt           string    `xorm:"default 0.00 comment('已还金额') DECIMAL(18,2)"`
	OverdueRate       string    `xorm:"comment('预期费率') DECIMAL(18,2)"`
	Status            int       `xorm:"default 0 comment('状态, 0:还款中, 1:已还款') INT(2)"`
	PreStatus         int       `xorm:"default 10 comment('预生成状态 10:预生成  20:正式') INT(2)"`
	DebtOauthStatus   int       `xorm:"default 10 comment('代扣授权状态 10，未授权 20，已授权') INT(2)"`
	GmtCreate         time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify         time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName        string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName        string    `xorm:"comment('修改人') VARCHAR(50)"`
	ActualRepayDate   time.Time `xorm:"comment('还款成功时间') DATETIME"`
}
