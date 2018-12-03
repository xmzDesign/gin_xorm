package models

import (
	"time"
)

type JiayinMutipleLoan struct {
	Id               int64     `xorm:"BIGINT(20)"`
	UserId           int64     `xorm:"BIGINT(20)"`
	Phone            string    `xorm:"comment('手机号') VARCHAR(13)"`
	LoanType         string    `xorm:"comment('LOAN_REGIST("EMR002", "信贷平台注册详情"), LOAN_APPLY("EMR004", "贷款申请详情"), LOAN_AUDIT("EMR007", "贷款放款详情"), LOAN_REFUSE("EMR009", "贷款驳回详情"), LOAN_OVERDUE("EMR012", "逾期平台详情"), LOAN_DEBTS("EMR013", "平台欠款");') VARCHAR(20)"`
	Cycle            string    `xorm:"comment('时间区间') VARCHAR(64)"`
	PType            int       `xorm:"comment('1:银行类借贷, 2:非银行类借贷') INT(2)"`
	PlatformCode     string    `xorm:"comment('平台代码') VARCHAR(64)"`
	RegisterTime     time.Time `xorm:"comment('该平台上的注册时间') DATETIME"`
	RecordTermDate   time.Time `xorm:"comment('该记录失效时间') DATE"`
	ApplyTime        time.Time `xorm:"comment('借贷申请时间') DATETIME"`
	ApplyAmount      string    `xorm:"comment('申请金额') VARCHAR(64)"`
	ApplyResult      string    `xorm:"comment('申请结果') VARCHAR(64)"`
	LoanLenderTime   time.Time `xorm:"comment('放款时间') DATETIME"`
	LoanLenderAmount string    `xorm:"comment('放款金额') VARCHAR(64)"`
	RefuseTime       time.Time `xorm:"comment('驳回时间') DATETIME"`
	Counts           int       `xorm:"comment('记录数量') INT(11)"`
	OverdueMoney     string    `xorm:"comment('逾期金额区间') VARCHAR(64)"`
	OverdueDate      time.Time `xorm:"comment('最近逾期时间') DATE"`
	Debt             string    `xorm:"comment('欠款金额区间') VARCHAR(64)"`
	GmtCreate        time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify        time.Time `xorm:"comment('修改时间') DATETIME"`
}
