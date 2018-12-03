package models

import (
	"time"
)

type JiayinLoanOverview struct {
	Id                      int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId                  int64     `xorm:"BIGINT(20)"`
	Mobile                  string    `xorm:"comment('手机号') VARCHAR(13)"`
	IdNo                    string    `xorm:"comment('身份证号') VARCHAR(18)"`
	CertName                string    `xorm:"comment('姓名') VARCHAR(25)"`
	CountLoanRegist         int       `xorm:"comment('信贷平台注册数量') INT(11)"`
	CountLoanApply          int       `xorm:"comment('贷款申请数量') INT(11)"`
	CountLoanRefuse         int       `xorm:"comment('申请贷款被拒数量') INT(11)"`
	CountLoanBankoverdue    int       `xorm:"comment('贷款银行逾期数量') INT(11)"`
	CountLoanNonbankoverdue int       `xorm:"comment('贷款非银行逾期数量') INT(11)"`
	CountLoanOverdue        int       `xorm:"comment('总逾期数') INT(11)"`
	CountLoanAudit          int       `xorm:"comment('借款成功数量') INT(11)"`
	CountLoanDebts          int       `xorm:"comment('贷款欠款数量') INT(11)"`
	RecordTermDate          time.Time `xorm:"comment('该记录有效时间') DATETIME"`
	GmtCreate               time.Time `xorm:"DATETIME"`
	GmtModify               time.Time `xorm:"DATETIME"`
	Province                string    `xorm:"comment('归属省') VARCHAR(18)"`
	City                    string    `xorm:"comment('归属市') VARCHAR(18)"`
}
