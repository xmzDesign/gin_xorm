package models

import (
	"time"
)

type LoanRepayUnderlineRecord struct {
	Id               int64     `xorm:"pk autoincr comment('序号') BIGINT(20)"`
	MemberId         int64     `xorm:"comment('用户id') index BIGINT(20)"`
	AccountNo        string    `xorm:"comment('转账账号') VARCHAR(30)"`
	RecieveAccountNo string    `xorm:"comment('到账账号') VARCHAR(30)"`
	AmtPay           string    `xorm:"default 0.00 comment('金额') DECIMAL(18,2)"`
	BillItemId       int64     `xorm:"comment('账单明细Id bill_item_info.id') index BIGINT(20)"`
	PayStatus        int       `xorm:"comment('状态 10：待支付 20：支付成功 30：支付失败 40：支付关闭') INT(2)"`
	RepayStatus      int       `xorm:"comment('还款状态 10：还款中 20：还款成功 30：还款失败') INT(2)"`
	RepaySuccessTime time.Time `xorm:"comment('转账时间') DATETIME"`
	PayChannel       int       `xorm:"comment('支付渠道 70:（alipay支付宝转账）   80 微信转账') INT(2)"`
	Memo             string    `xorm:"comment('备注信息') VARCHAR(100)"`
	GmtCreate        time.Time `xorm:"comment('创建时间') DATETIME"`
	GmtModify        time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateName       string    `xorm:"comment('创建人') VARCHAR(50)"`
	ModifyName       string    `xorm:"comment('修改人') VARCHAR(50)"`
	ApplyName        string    `xorm:"comment('填单人') VARCHAR(50)"`
}
