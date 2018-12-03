package models

import (
	"time"
)

type BillItemInfo struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	BillId           int64     `xorm:"default 0 comment('账单ID') index BIGINT(20)"`
	Stages           int       `xorm:"default 0 comment('期数') INT(11)"`
	BillAmount       string    `xorm:"default 0.00 comment('账单金额') index DECIMAL(18,2)"`
	MemberId         int64     `xorm:"default 0 comment('用户ID') index BIGINT(11)"`
	Principal        string    `xorm:"default 0.00 comment('本金') DECIMAL(18,2)"`
	ManageCost       string    `xorm:"default 0.00 comment('管理费') DECIMAL(18,2)"`
	LoanInterestCost string    `xorm:"default 0.00 comment('利息') DECIMAL(18,2)"`
	DueAmount        string    `xorm:"default 0.00 comment('应还金额') DECIMAL(18,2)"`
	IsOverdue        int       `xorm:"default 0 comment('是否逾期, 0:未逾期, 1:已逾期') INT(11)"`
	DueDate          time.Time `xorm:"comment('应还款日') index DATETIME(3)"`
	RepayDate        time.Time `xorm:"comment('实际还款日') index DATETIME(3)"`
	OverdueDays      int       `xorm:"default 0 comment('逾期天数') index INT(11)"`
	OverdueAmt       string    `xorm:"default 0.00 comment('逾期金额') DECIMAL(18,2)"`
	DeductAmt        string    `xorm:"default 0.00 comment('已代扣金额') DECIMAL(18,2)"`
	RepayId          int64     `xorm:"comment('还款记录Id') index BIGINT(18)"`
	AbateAmt         string    `xorm:"default 0.00 comment('减免总金额') DECIMAL(18,2)"`
	Status           int       `xorm:"default 0 comment('还款状态: 0:未还款, 1:已还款') INT(2)"`
	PreStatus        int       `xorm:"default 10 comment('预生成状态 10:预生成  20:正式') INT(2)"`
	GmtCreate        time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify        time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName       string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName       string    `xorm:"comment('修改人') VARCHAR(50)"`
	BadLoanStatus    int       `xorm:"default 10 comment('坏账状态 10,正常 20，坏账') INT(2)"`
	IsAdjourn        int       `xorm:"default 10 comment('是否展期  10，未展期 20，展期') INT(2)"`
}
